package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection はデータベース接続の設定を表します。
type Connection struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// ToString はデータベース接続設定文字列を返します。
func (c *Connection) toString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.DBName)
}

// CreateDbConnection はデータベース接続を作成します。
func (c *Connection) CreateDbConnection() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(c.toString()), &gorm.Config{})
}
