package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name         string
	PrefectureID uuid.UUID
	Prefecture   Prefecture
}
