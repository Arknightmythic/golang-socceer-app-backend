package repositories

import (
	// Mengubah import ke package user yang spesifik
	repositories "user-service/repositories/user" 
	"gorm.io/gorm"
)

type Registry struct{
	db *gorm.DB
}


type IRepositoryRegistry interface{
	// Menggunakan interface dari package user
	GetUser() repositories.IUserRepository 
}


func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Registry{db: db}
}

// Mengembalikan implementasi dari package user
func (r *Registry) GetUser() repositories.IUserRepository { 
	return repositories.NewUserRepository(r.db)
}