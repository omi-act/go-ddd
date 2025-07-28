package value_objects

import (
	"errors"
	"strconv"
)

// UserID はユーザーIDを表現する値オブジェクトです。
type UserID struct {
	IDNumber int
}

// validate はUserIDの値を検証します。
func (u *UserID) validate() error {
	if u.IDNumber <= 0 {
		return errors.New("invalid ID: ID must be greater than 0")
	}
	return nil
}

// NewUserID は新しいUserIDを作成します。
func NewUserID(value int) (*UserID, error) {
	idEntity := &UserID{IDNumber: value}
	if err := idEntity.validate(); err != nil {
		return nil, err
	}
	return idEntity, nil
}

// NewUserIDFromString は文字列から新しいUserIDを作成します。
func NewUserIDFromString(idStr string) (*UserID, error) {

	// 必須チェック
	if idStr == "" {
		return nil, errors.New("ID is required")
	}

	// 数値変換
	value, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, errors.New("invalid ID format: ID must be a number")
	}

	// UserIDの作成時にビジネスルール検証も実行される
	return NewUserID(value)
}
