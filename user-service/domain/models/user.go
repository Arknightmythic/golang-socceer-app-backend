package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uint      `gorm:"primaryKey;autoincrement"`
	UUID        uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"type:varchar(50);not null"`
	Username    string    `gorm:"type:varchar(50);not null;"`
	Password    string    `gorm:"type:varchar(255);not null"`
	PhoneNumber string    `gorm:"type:varchar(15);not null"`
	Email       string    `gorm:"type:varchar(100);not null"`
	RoleID      uint      `gorm:"type:uint;not null"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	Role        Role `gorm:"foreignKey:role_id; references:id;constraint:OnUpdate:CASCADE,Ondelete:CASCADE"`
}
