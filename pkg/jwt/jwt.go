package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWTToken(userId int) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":          userId,
		"login_time_stamp": time.Now(),
	})

	tokenString, err = token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return
}
