package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Prefecture struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;primary_key;"`
	Code  string
	Name  string
	Elias string
}
