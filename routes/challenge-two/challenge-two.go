package challenge_two

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	controllers "test-case-roketin/controllers/challenge-two"
	"test-case-roketin/repositories/admin"
	"test-case-roketin/repositories/movie"
	services "test-case-roketin/services/challenge-two"
	"test-case-roketin/utils/middleware"
	"test-case-roketin/utils/pagination"
)

func ChallengeTwoRoute(db *gorm.DB, routerGroup *gin.RouterGroup) {
	//Repositories
	adminRepository := admin.NewAdminRepository(db)
	movieRepository := movie.NewMovieRepository(db)

	// Services
	challengeTwoService := services.NewChallengeTwoService(adminRepository, movieRepository)

	//pagination
	paginate := pagination.NewPagination()

	// Controllers
	challengeTwoController := controllers.NewChallengeTwoController(challengeTwoService, *paginate)

	// Endpoints
	routerGroup.POST("/challenge-two/register-admin", challengeTwoController.Register)
	routerGroup.POST("/challenge-two/login-admin", challengeTwoController.Login)
	routerGroup.POST("/challenge-two/create-movie", middleware.Middleware(), challengeTwoController.CreateMovie)
	routerGroup.PATCH("/challenge-two/update-movie/:id", middleware.Middleware(), challengeTwoController.UpdateMovie)
	routerGroup.GET("/challenge-two/get-all-movie", middleware.Middleware(), challengeTwoController.GetAllMovie)

}
