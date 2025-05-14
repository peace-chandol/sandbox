package models

import "github.com/google/uuid"

type Role string

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string    `gorm:"unique;not null"`
	Email    string    `gorm:"not null"`
	Password string    `gorm:"not null"`
	Role     string    `gorm:"not null"`
}

type Todo struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	Text   string    `gorm:"not null"`
	Done   bool      `gorm:"not null"`
	UserID uuid.UUID `gorm:"not null"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
