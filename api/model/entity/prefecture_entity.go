package entity

import "github.com/google/uuid"

type Prefecture struct {
	BaseModel
	ID    uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name  string
	Elias string
}
