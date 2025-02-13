package models

import baseModels "test-case-roketin/common/base-models"

type Movie struct {
	baseModels.Base
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Description string `gorm:"type:text;not null" json:"description"`
	Duration    string `gorm:"type:varchar(255);not null" json:"duration"`
	Artists     string `gorm:"type:text;not null" json:"artists"`
	Genres      string `gorm:"type:text;not null" json:"genres"`
}
