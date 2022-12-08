package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserClaims struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
