package route

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test-case-roketin/common/constants"
	"test-case-roketin/utils/env"
)

func InitRouter(config env.Config) *gin.Engine {
	if config.GinMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With", "Link", "X-Log-Access"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With, Link, X-Log-Access")
		c.Next()
	})

	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With, Link, X-Log-Access")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.AbortWithStatus(204)
	})

	router.Use(customRecovery())

	// limit file upload
	router.MaxMultipartMemory = 10 << 20

	return router
}

func RunRoute(config env.Config, router *gin.Engine) {
	err := router.Run(fmt.Sprintf(":%s", config.GinPort))
	if err != nil {
		message := fmt.Sprintf("%s failed to start server", constants.Server)
		log.Fatal(message)
	}
}

func customRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": http.StatusInternalServerError,
					"success":     false,
					"message":     "please tell about this error to the administrator or developer",
					"data":        nil,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
