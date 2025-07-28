package mapper

import (
	"go-ddd/internal/domain/entities"
	dbEntities "go-ddd/internal/infrastructure/postgres/entities"
)

// UserMapper はDBエンティティとドメインエンティティの変換を行う
type UserMapper struct{}

// NewUserMapper は新しいUserMapperを作成します
func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

// ToUser はユーザーのDBエンティティをドメインエンティティに変換します
func ToUser(dbUser *dbEntities.User) *entities.User {
	if dbUser == nil {
		return nil
	}

	return &entities.User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		Age:       dbUser.Age,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

// ToUsers はユーザーのDBエンティティのスライスをドメインエンティティのスライスに変換します
func ToUsers(dbUsers []*dbEntities.User) []*entities.User {
	if dbUsers == nil {
		return nil
	}

	entities := make([]*entities.User, len(dbUsers))
	for i, dbUser := range dbUsers {
		entities[i] = ToUser(dbUser)
	}
	return entities
}