package movie

import (
	"gorm.io/gorm"
	"test-case-roketin/common/constants"
	"test-case-roketin/models"
	"test-case-roketin/utils/pagination"
)

type MovieRepository interface {
	Create(movie models.Movie) (models.Movie, error)
	Update(movie models.Movie) (models.Movie, error)
	FindAll(paginate pagination.Pagination, search string) ([]models.Movie, pagination.Pagination, error)
	FindById(id int64) (models.Movie, error)
}

type movieRepo struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *movieRepo {
	return &movieRepo{db}
}

func (r *movieRepo) Create(movie models.Movie) (models.Movie, error) {
	err := r.db.Create(&movie).Error

	return movie, err
}

func (r *movieRepo) Update(movie models.Movie) (models.Movie, error) {
	err := r.db.Save(&movie).Error

	return movie, err
}

func (r *movieRepo) FindAll(paginate pagination.Pagination, search string) ([]models.Movie, pagination.Pagination, error) {
	var countries []models.Movie

	err := r.db.Scopes(pagination.Paginate(&countries, &paginate, r.db)).Where(search).Find(&countries).Error

	return countries, paginate, err
}

func (r *movieRepo) FindById(id int64) (models.Movie, error) {
	var movie models.Movie

	err := r.db.Where(constants.ById, id).First(&movie).Error

	return movie, err
}
