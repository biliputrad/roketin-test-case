package challenge_one

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	challengeOneDto "test-case-roketin/common/dto/challenge-one-dto"
	responseMessage "test-case-roketin/common/response-message"
	challengeOne "test-case-roketin/services/challenge-one"
)

type challengeOneController struct {
	challengeOneService challengeOne.ChallengeOneService
}

func NewChallengeOneController(challengeOneService challengeOne.ChallengeOneService) *challengeOneController {
	return &challengeOneController{challengeOneService}
}

func (c *challengeOneController) ConvertEarthTimeToRoketinTime(ctx *gin.Context) {
	var dto challengeOneDto.ChallengeOne
	// Bind request to dto
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		errorMessages := responseMessage.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := responseMessage.GetResponse(http.StatusBadRequest, false, errorMessage, false)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Call service
	result := c.challengeOneService.ConvertEarthTimeToRoketinTime(dto)

	// Return response message
	ctx.JSON(result.StatusCode, result)
}
