package challenge_two

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"net/http"
	baseModels "test-case-roketin/common/base-models"
	"test-case-roketin/common/constants"
	challengeTwoDto "test-case-roketin/common/dto/challenge-two-dto"
	hashPassword "test-case-roketin/common/hash-password"
	responseMessage "test-case-roketin/common/response-message"
	"test-case-roketin/models"
	"test-case-roketin/repositories/admin"
	"test-case-roketin/repositories/movie"
	jwtToken "test-case-roketin/utils/jwt-token"
	"test-case-roketin/utils/pagination"
	"time"
)

type ChallengeTwoService interface {
	Register(dto challengeTwoDto.Register) responseMessage.Response
	Login(dto challengeTwoDto.Login) responseMessage.Response
	CreateMovie(dto challengeTwoDto.CreateMovie) responseMessage.Response
	UpdateMovie(dto challengeTwoDto.UpdateMovie) responseMessage.Response
	GetAllMovie(pagination pagination.Pagination, search string) responseMessage.ResponsePaginate
}

type challengeTwoService struct {
	adminRepository admin.AdminRepository
	movieRepository movie.MovieRepository
}

func NewChallengeTwoService(
	adminRepository admin.AdminRepository,
	movieRepository movie.MovieRepository,
) *challengeTwoService {
	return &challengeTwoService{
		adminRepository,
		movieRepository,
	}
}

func (s *challengeTwoService) Register(dto challengeTwoDto.Register) responseMessage.Response {
	adminData, err := s.adminRepository.FindByUsername(dto.Username)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return responseMessage.Response{
				StatusCode: http.StatusInternalServerError,
				Success:    false,
				Message:    constants.ResponseInternalServerError,
				Data:       nil,
			}
		}
		err = nil
	}

	if adminData.ID != 0 {
		return responseMessage.Response{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    "Username already exists",
			Data:       nil,
		}
	}

	newPassword, err := hashPassword.HashPassword(dto.Password)
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    constants.ResponseInternalServerError,
			Data:       nil,
		}
	}

	_, err = s.adminRepository.Create(models.Admin{
		Base: baseModels.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: dto.Username,
		Password: newPassword,
	})
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    constants.ResponseInternalServerError,
			Data:       nil,
		}
	}

	return responseMessage.Response{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    constants.ResponseCreated,
		Data:       nil,
	}
}

func (s *challengeTwoService) Login(dto challengeTwoDto.Login) responseMessage.Response {
	adminData, err := s.adminRepository.FindByUsername(dto.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responseMessage.Response{
				StatusCode: http.StatusUnauthorized,
				Success:    false,
				Message:    "Username or password is incorrect",
				Data:       nil,
			}
		}
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    constants.ResponseInternalServerError,
			Data:       nil,
		}
	}

	if !hashPassword.ComparePassword(adminData.Password, dto.Password) {
		return responseMessage.Response{
			StatusCode: http.StatusUnauthorized,
			Success:    false,
			Message:    "Username or password is incorrect",
			Data:       nil,
		}
	}

	generateToken, expiredAt, err := jwtToken.GenerateToken(jwtToken.JwtClaim{
		ID:             adminData.ID,
		Username:       adminData.Username,
		StandardClaims: jwt.StandardClaims{},
	})
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    constants.ResponseInternalServerError,
			Data:       nil,
		}
	}

	return responseMessage.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data: challengeTwoDto.ResultLogin{
			Token:     generateToken,
			ExpiredAt: expiredAt,
			TokenType: constants.Bearer,
		},
	}
}

func (s *challengeTwoService) CreateMovie(dto challengeTwoDto.CreateMovie) responseMessage.Response {
	result, err := s.movieRepository.Create(models.Movie{
		Base: baseModels.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:       dto.Title,
		Description: dto.Description,
		Duration:    dto.Duration,
		Artists:     dto.Artists,
		Genres:      dto.Genres,
	})
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    constants.ResponseInternalServerError,
			Data:       nil,
		}
	}

	return responseMessage.Response{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    constants.ResponseCreated,
		Data:       result,
	}
}

func (s *challengeTwoService) UpdateMovie(dto challengeTwoDto.UpdateMovie) responseMessage.Response {
	movieData, err := s.movieRepository.FindById(dto.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responseMessage.Response{
				StatusCode: http.StatusNotFound,
				Success:    false,
				Message:    fmt.Errorf("movie with id %d not found", dto.Id).Error(),
				Data:       nil,
			}
		}
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    constants.ResponseInternalServerError,
			Data:       nil,
		}
	}

	model := models.Movie{
		Base: baseModels.Base{
			ID:        dto.Id,
			CreatedAt: movieData.CreatedAt,
			UpdatedAt: time.Now(),
		},
		Title:       dto.Title,
		Description: dto.Description,
		Duration:    dto.Duration,
		Artists:     dto.Artists,
		Genres:      dto.Genres,
	}

	result, err := s.movieRepository.Update(model)
	if err != nil {
		return responseMessage.Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    constants.ResponseInternalServerError,
			Data:       nil,
		}
	}

	return responseMessage.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
	}
}

func (s *challengeTwoService) GetAllMovie(pagination pagination.Pagination, search string) responseMessage.ResponsePaginate {
	if search != "" {
		search = "title ILIKE '%" + search + "%' OR description ILIKE '%" + search + "%' OR artists ILIKE '%" + search + "%' OR genres ILIKE '%" + search + "%'"
	}

	result, paginateResult, err := s.movieRepository.FindAll(pagination, search)
	if err != nil {
		return responseMessage.ResponsePaginate{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
			Data:       nil,
			Pagination: nil,
		}
	}

	return responseMessage.ResponsePaginate{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    constants.ResponseOK,
		Data:       result,
		Pagination: paginateResult,
	}
}
