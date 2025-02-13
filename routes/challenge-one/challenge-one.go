package challenge_one

import (
	"github.com/gin-gonic/gin"
	controllers "test-case-roketin/controllers/challenge-one"
	services "test-case-roketin/services/challenge-one"
)

func ChallengeOneRoute(routerGroup *gin.RouterGroup) {
	// Services
	challengeOneService := services.NewChallengeOneService()

	// Controllers
	challengeOneController := controllers.NewChallengeOneController(challengeOneService)

	// Endpoints
	routerGroup.POST("/challenge-one/", challengeOneController.ConvertEarthTimeToRoketinTime)
}
