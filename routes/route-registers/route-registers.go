package route_registers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	challengeOne "test-case-roketin/routes/challenge-one"
	challengeTwo "test-case-roketin/routes/challenge-two"
)

func RouteRegister(db *gorm.DB, routerGroup *gin.RouterGroup) {
	challengeOne.ChallengeOneRoute(routerGroup)
	challengeTwo.ChallengeTwoRoute(db, routerGroup)
}
