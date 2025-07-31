package test

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/k1LoW/runn"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"go-ddd/bootstrap"
)

const (
	envFile     = "../.env.test"
	runnBookDir = "./scenarios"

	// 終了コード定数
	ExitError   = 1 // エラー終了
)

// TestMain はDBコンテナのセットアップとE2Eテストの実行を行います。
func TestMain(m *testing.M) {

	// .envファイルを読み込み
	if err := godotenv.Load(envFile); err != nil {
		fmt.Printf("Warning: .env file not found or could not be loaded: %v\n", err)
		os.Exit(ExitError)
	}

	// 環境変数からアプリケーションのポートを取得
	appPort := os.Getenv("SERVER_PORT")
	if appPort == "" {
		fmt.Println("SERVER_PORT environment variable is not set")
		os.Exit(ExitError)
	}

	// DBコンテナの起動
	ctx := context.Background()
	postgresC, err := startPostgresContainer(ctx)
	if err != nil {
		fmt.Printf("Failed to start postgres container: %v\n", err)
		os.Exit(ExitError)
	}
	defer func() {
		if err := postgresC.Terminate(ctx); err != nil {
			fmt.Printf("Failed to terminate postgres container: %v\n", err)
		}
	}()

	fmt.Println("Test database initialized with Dockerfile and init.sql")

	// アプリケーション起動
	go func() {
		e, err := bootstrap.Initialize(envFile)
		if err != nil {
			fmt.Printf("Failed to initialize application: %v\n", err)
			os.Exit(ExitError)
		}
		if err := e.Start(":" + appPort); err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
			os.Exit(ExitError)
		}
	}()
	time.Sleep(3 * time.Second) // 起動待機

	// テストの実行
	code := m.Run()
	os.Exit(code)
}

// startPostgresContainer はPostgreSQLコンテナを起動し、接続確認を行います。
func startPostgresContainer(ctx context.Context) (testcontainers.Container, error) {

	// コンテナの設定
	hostPort := os.Getenv("DB_PORT")
	if hostPort == "" {
		return nil, fmt.Errorf("DB_PORT environment variable is not set")
	}

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "./docker",
			Dockerfile: "Dockerfile.postgres",
		},
		ExposedPorts: []string{hostPort + ":5432"},
		WaitingFor:   wait.ForListeningPort("5432/tcp").WithStartupTimeout(100 * time.Second),
	}
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres container: %w", err)
	}

	// データベース接続の確認
	host, err := postgresC.Host(ctx)
	if err != nil {
		return nil, err
	}
	port, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		return nil, err
	}

	// DB接続
	dsn := fmt.Sprintf("postgres://user:password@%s:%s/testdb?sslmode=disable", host, port.Port())
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	// database/sql + pqドライバを使う場合
	var db *sql.DB
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return postgresC, db.Ping()
}

// Test_APIIntegrationはAPI結合テストを実行する関数です。
func Test_APIIntegration(t *testing.T) {
	ctx := context.Background()

	// runnbookファイルを取得
	runnBookFiles, err := getRunnBookFiles()
	if err != nil {
		t.Fatalf("Failed to get runnbook files: %v", err)
	}

	// 環境変数からアプリケーションのホストとポートを取得
	appHost := os.Getenv("SERVER_HOST")
	appPort := os.Getenv("SERVER_PORT")
	if appHost == "" || appPort == "" {
		t.Fatalf("SERVER_HOST or SERVER_PORT environment variable is not set")
	}
	url := fmt.Sprintf("http://%s:%s", appHost, appPort)

	// NOTE: ランブックごとにAPIテストを実行
	for _, runnBookFile := range runnBookFiles {
		t.Run(runnBookFile, func(t *testing.T) {
			runnBookFilePath := filepath.Join(runnBookDir, runnBookFile)

			// NOTE: runn を実行する
			opts := []runn.Option{
				runn.T(t),
				runn.Book(runnBookFilePath),
				runn.Runner("req", url),
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
