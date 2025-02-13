package challenge_two

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	challengeTwoDto "test-case-roketin/common/dto/challenge-two-dto"
	responseMessage "test-case-roketin/common/response-message"
	challengeTwo "test-case-roketin/services/challenge-two"
	"test-case-roketin/utils/pagination"
)

type challengeTwoController struct {
	challengeTwoService challengeTwo.ChallengeTwoService
	paginate            pagination.Pagination
}

func NewChallengeTwoController(challengeTwoService challengeTwo.ChallengeTwoService, paginate pagination.Pagination) *challengeTwoController {
	return &challengeTwoController{challengeTwoService, paginate}
}

func (c *challengeTwoController) Register(ctx *gin.Context) {
	var dto challengeTwoDto.Register
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		errorMessages := responseMessage.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := responseMessage.GetResponse(http.StatusBadRequest, false, errorMessage, false)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.challengeTwoService.Register(dto)

	ctx.JSON(result.StatusCode, result)
}

func (c *challengeTwoController) Login(ctx *gin.Context) {
	var dto challengeTwoDto.Login
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		errorMessages := responseMessage.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := responseMessage.GetResponse(http.StatusBadRequest, false, errorMessage, false)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.challengeTwoService.Login(dto)

	ctx.JSON(result.StatusCode, result)
}

func (c *challengeTwoController) CreateMovie(ctx *gin.Context) {
	var dto challengeTwoDto.CreateMovie
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		errorMessages := responseMessage.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := responseMessage.GetResponse(http.StatusBadRequest, false, errorMessage, false)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.challengeTwoService.CreateMovie(dto)

	ctx.JSON(result.StatusCode, result)
}

func (c *challengeTwoController) UpdateMovie(ctx *gin.Context) {
	var dto challengeTwoDto.UpdateMovie
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		errorMessages := responseMessage.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := responseMessage.GetResponse(http.StatusBadRequest, false, errorMessage, false)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	tempId := ctx.Param("id")
	dto.Id, err = strconv.ParseInt(tempId, 10, 64)
	if err != nil {
		res := responseMessage.GetResponse(http.StatusBadRequest, false, "Id must be integer", false)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.challengeTwoService.UpdateMovie(dto)

	ctx.JSON(result.StatusCode, result)
}

func (c *challengeTwoController) GetAllMovie(ctx *gin.Context) {
	paginationResult, search := c.paginate.GetPagination(ctx)
	result := c.challengeTwoService.GetAllMovie(paginationResult, search)

	ctx.JSON(result.StatusCode, result)
}
