package middleware

import (
	"os"
	"time"

	"github.com/AdluAghnia/go-artic/models"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var jwtSecretKey []byte

func GenerateJWTKey(user *models.User) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	jwtSecretKey = []byte(os.Getenv("SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(jwtSecretKey)
}


func JWTMiddleware() fiber.Handler {
	return func (c *fiber.Ctx) error {
		tokenString := c.Cookies("jwt")


		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).Redirect("/login")
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).Redirect("/login")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).Redirect("/login")
		}

		userID := claims["id"]
		c.Locals("userID", userID)

		return c.Next()
	}
}
