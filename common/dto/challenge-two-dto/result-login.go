package challenge_two_dto

import "time"

type ResultLogin struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	TokenType string    `json:"token_type"`
}
