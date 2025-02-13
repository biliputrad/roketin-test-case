package admin

import (
	"gorm.io/gorm"
	"test-case-roketin/common/constants"
	"test-case-roketin/models"
)

type AdminRepository interface {
	Create(admin models.Admin) (models.Admin, error)
	FindByUsername(username string) (models.Admin, error)
}

type adminRepo struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepo {
	return &adminRepo{db}
}

func (r *adminRepo) Create(admin models.Admin) (models.Admin, error) {
	err := r.db.Create(&admin).Error

	return admin, err
}

func (r *adminRepo) FindByUsername(username string) (models.Admin, error) {
	var result models.Admin

	err := r.db.Where(constants.ByUsername, username).First(&result).Error

	return result, err
}
