package testcases

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
	"time"
	"os"

	"github.com/k1LoW/runn"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	testServerURL           = "http://localhost:9000"
	runnBookDir             = "./testcases"
)

// TestMain はDBコンテナのセットアップとE2Eテストの実行を行います。
func TestMain(m *testing.M) {

	// ポート番号の定義
	const hostPort = "5430"

	// コンテキストの作成
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "./docker",
			Dockerfile: "Dockerfile.postgres",
		},
		ExposedPorts: []string{hostPort + ":5432"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_USER":     "user",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(100 * time.Second),
	}

	// コンテナの起動
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	// コンテナの終了処理
	if err != nil {
		panic(err)
	}
	defer postgresC.Terminate(ctx)

	// ホストとポートの取得
	host, err := postgresC.Host(ctx)
	if err != nil {
		panic(err)
	}
	port, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		panic(err)
	}

	// DB接続
	dsn := fmt.Sprintf("postgres://user:password@%s:%s/testdb?sslmode=disable", host, port.Port())
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	// database/sql + pqドライバを使う場合
	var db *sql.DB
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// データベース接続の確認（init.sqlは自動実行されるため手動実行不要）
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	fmt.Println("Test database initialized with Dockerfile and init.sql")

	// テストの実行
	code := m.Run()
	os.Exit(code)
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

		// .ymlまたは.yamlファイルで、test.ymlで終わるファイルを対象とする
		if !d.IsDir() && (strings.HasSuffix(path, "_test.yml") || strings.HasSuffix(path, "_test.yaml")) {
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

	return runnBookFiles, nil
}

