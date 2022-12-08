package handlers

import (
	"github.com/daniilmikhaylov2005/town-shop-rest-api/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(user models.User) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", err
	}
	claims := &models.UserClaims{
		Name:     user.Name,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stringKey := os.Getenv("SIGN_KEY")
	byteKey := []byte(stringKey)

	stringToken, err := token.SignedString(byteKey)
	return stringToken, err
}
