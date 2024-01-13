package entity

import "github.com/google/uuid"

type City struct {
	BaseModel
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name         string
	NameKana     string
	PrefectureID uuid.UUID
	Prefecture   Prefecture
}
