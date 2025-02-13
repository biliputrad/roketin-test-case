package jwt_token

import (
	"github.com/dgrijalva/jwt-go"
	"test-case-roketin/utils/env"
	"time"
)

type JwtClaim struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

// GenerateToken is used to create new token and will return token and time expired token
func GenerateToken(claim JwtClaim) (string, time.Time, error) {
	// Load Local Time Zone
	timeZone, err := time.LoadLocation(env.GlobalEnv.DbTz)
	if err != nil {
		return "", time.Time{}, err
	}

	jwtKey := []byte(env.GlobalEnv.SecretKey)
	expirationTime := time.Now().Add(time.Duration(env.GlobalEnv.ExpiredTime) * time.Hour).In(timeZone)

	claims := claim
	claims.StandardClaims = jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, expirationTime, err
}
