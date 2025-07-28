package entities

import (
	"time"
	"fmt"
)

// User represents a user entity in the domain
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser creates a new User entity
func NewUser(id int, name, email string, age int) *User {
	now := time.Now()
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Age:       age,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Show はユーザーの情報を表示するメソッドです。
func (u *User) Show() string {
	return u.showName() + fmt.Sprintf("(info: %s)", u.toString())
}

// showName はユーザーの名前を返すメソッドです。
func (u *User) showName() string {
	return u.Name + "さん"
}

// toString はユーザーの文字列表現を返すメソッドです。
func (u *User) toString() string {
	return fmt.Sprintf("ID: %d, Name: %s, Email: %s, Age: %d, CreatedAt: %s, UpdatedAt: %s",
		u.ID, u.Name, u.Email, u.Age, u.CreatedAt, u.UpdatedAt)
}
