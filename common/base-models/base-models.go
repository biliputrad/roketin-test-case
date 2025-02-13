package base_models

import (
	"time"
)

type Base struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"<-create" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
