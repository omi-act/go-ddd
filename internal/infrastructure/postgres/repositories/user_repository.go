package repositories

import (
	"fmt"

	"go-ddd/internal/domain/entities"
	"go-ddd/internal/domain/repositories"
	"go-ddd/internal/domain/value_objects"
	dbEntities "go-ddd/internal/infrastructure/postgres/entities"
	"go-ddd/internal/infrastructure/postgres/mapper"

	"gorm.io/gorm"
)

// GormUserRepository は GORM を使用してユーザーデータを操作するリポジトリの実装です。
type GormUserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository は新しい GormUserRepository を作成します。
func NewGormUserRepository(db *gorm.DB) repositories.UserRepository {
	return &GormUserRepository{
		db: db,
	}
}

// FindByID は ID に基づいてユーザーを取得します。
func (r *GormUserRepository) FindByID(userID *value_objects.UserID) (*entities.User, error) {
	var dbUser dbEntities.User
	if err := r.db.First(&dbUser, userID.IDNumber).Error; err != nil {
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}
	return mapper.ToUser(&dbUser), nil
}

// FindAll はすべてのユーザーを取得します。
func (r *GormUserRepository) FindAll(limit, offset int) ([]*entities.User, error) {
	var dbUsers []*dbEntities.User
	if err := r.db.Limit(limit).Offset(offset).Find(&dbUsers).Error; err != nil {
		return nil, fmt.Errorf("failed to find all users: %w", err)
	}
	return mapper.ToUsers(dbUsers), nil
}
