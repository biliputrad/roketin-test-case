package models

import baseModels "test-case-roketin/common/base-models"

type Admin struct {
	baseModels.Base
	Username string `gorm:"type:varchar(255);unique;not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
}
