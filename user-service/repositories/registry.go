package repositories

import (
	repositories "user-service/repositories"
	"gorm.io/gorm"
)

type Register struct{
	db *gorm.DB
}


type IRepositoryRegistry interface{
	GetUser() repositories.IUserRepository
}


func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Register{db: db}
}

func (r *Register) GetUser() repositories.IUserRepository {
	return repositories.NewUserRepository(r.db)
}