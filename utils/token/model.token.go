package token

import "github.com/golang-jwt/jwt/v5"

type MyCustomClaims struct {
	UserID int `json:"userID"`
	jwt.RegisteredClaims
}
