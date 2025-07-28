package testcases

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/k1LoW/runn"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"go-ddd/internal/app"
)

var (
	testAppPort   = "9100"
	testServerURL = "http://localhost:9100"
	runnBookDir   = "./testcases"
)

// TestMain はDBコンテナのセットアップとE2Eテストの実行を行います。
func TestMain(m *testing.M) {

	// ポート番号の定義
	const hostPort = "5430"

	// コンテキストの作成
	ctx := context.Background()

	// TODO: 設定整理
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "./docker",
			Dockerfile: "Dockerfile.postgres",
		},
		ExposedPorts: []string{hostPort + ":5432"},
		// Env: map[string]string{
		// 	"POSTGRES_PASSWORD": "password",
		// 	"POSTGRES_USER":     "user",
		// 	"POSTGRES_DB":       "testdb",
		// },
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(100 * time.Second),
	}

	// コンテナの起動
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	defer postgresC.Terminate(ctx) // コンテナの終了処理

	// DB疎通確認
	if err := ping(ctx, postgresC); err != nil {
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}
	fmt.Println("Test database initialized with Dockerfile and init.sql")

	// テスト環境フラグを設定
	os.Setenv("TEST_ENV", "true")

	// サーバ起動
	go func() {
		e := app.Initialize()
		if err := e.Start(":" + testAppPort); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// テストの実行
	code := m.Run()
	os.Exit(code)
}

// ping はPostgreSQLコンテナへの接続確認を行います。
// エラーが発生した場合はその内容を返します。
func ping(ctx context.Context, postgresC testcontainers.Container) error {

	host, err := postgresC.Host(ctx)
	if err != nil {
		return err
	}
	port, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		return err
	}

	// DB接続
	dsn := fmt.Sprintf("postgres://user:password@%s:%s/testdb?sslmode=disable", host, port.Port())
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	// database/sql + pqドライバを使う場合
	var db *sql.DB
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	// データベース接続の確認（init.sqlは自動実行されるため手動実行不要）
	return db.Ping()
}

// Test_e2eはE2Eテストを実行する関数です。
func Test_e2e(t *testing.T) {
	ctx := context.Background()

	// runnbookファイルを取得
	runnBookFiles, err := getRunnBookFiles()
	if err != nil {
		t.Fatalf("Failed to get runnbook files: %v", err)
	}

	// NOTE: ランブックごとにAPIテストを実行
	for _, runnBookFile := range runnBookFiles {
		t.Run(runnBookFile, func(t *testing.T) {
			runnBookFilePath := filepath.Join(runnBookDir, runnBookFile)

			// NOTE: runn を実行する
			opts := []runn.Option{
				runn.T(t),
				runn.Book(runnBookFilePath),
				runn.Runner("req", testServerURL),
				runn.Scopes("read:parent"),
			}
			o, err := runn.New(opts...)
			if err != nil {
				t.Errorf("Failed to create runn instance for %s: %v", runnBookFile, err)
				return
			}
			if err := o.Run(ctx); err != nil {
				t.Errorf("Failed to run runn for %s: %v", runnBookFile, err)
				return
			}
		})
	}
}

// getRunnBookFiles は実行対象のrunnbookファイル一覧を取得します。
func getRunnBookFiles() ([]string, error) {
	var runnBookFiles []string

	err := filepath.WalkDir(runnBookDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// .ymlまたは.yamlファイルを対象とする
		if !d.IsDir() && (strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml")) {
			// 相対パスを取得
			relativePath, err := filepath.Rel(runnBookDir, path)
			if err != nil {
				return err
			}
			runnBookFiles = append(runnBookFiles, relativePath)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %w", err)
	}
	if len(runnBookFiles) == 0 {
		return nil, fmt.Errorf("no runnbook files found")
	}

	return runnBookFiles, nil
}
