package hash_password

import (
	"golang.org/x/crypto/bcrypt"
	"test-case-roketin/utils/env"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), env.GlobalEnv.SaltPassword)
	if err != nil {
		return string(hashed), err
	}
	return string(hashed), err
}
