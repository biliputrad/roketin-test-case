package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"test-case-roketin/common/constants"
	responseMessage "test-case-roketin/common/response-message"
	"test-case-roketin/utils/env"
	jwtToken "test-case-roketin/utils/jwt-token"
	"time"
)

func Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		bearerToken := context.GetHeader(constants.Authorization)
		if !strings.Contains(bearerToken, constants.Bearer) {
			res := responseMessage.GetResponse(http.StatusUnauthorized, false, constants.ResponseInvalidToken, nil)
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}

		tokenString := ""
		arrayToken := strings.Split(bearerToken, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		err := ValidateUserToken(tokenString)
		if err != nil {
			res := responseMessage.GetResponse(http.StatusUnauthorized, false, err.Error(), nil)
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}

		context.Next()
	}
}

func ValidateUserToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtToken.JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GlobalEnv.SecretKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*jwtToken.JwtClaim)
	if !ok {
		err = errors.New(constants.MsgParseErr)
		return err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New(constants.MsgTokenExpired)
		return err
	}

	return nil
}
