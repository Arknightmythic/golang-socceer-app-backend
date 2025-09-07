package repositories

import (
	// Mengubah import ke package user yang spesifik
	repositories "user-service/repositories/user" 
	"gorm.io/gorm"
)

type Register struct{
	db *gorm.DB
}


type IRepositoryRegistry interface{
	// Menggunakan interface dari package user
	GetUser() repositories.IUserRepository 
}


func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Register{db: db}
}

// Mengembalikan implementasi dari package user
func (r *Register) GetUser() repositories.IUserRepository { 
	return repositories.NewUserRepository(r.db)
}