package middleware

import (
	"fmt"
	"golang-hotel-reservation/db"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := c.GetReqHeaders()["X-Api-Token"]
		if !ok {
			fmt.Println("token not present in theheader")
			return fmt.Errorf("unauthorized")
		}
		claims, err := validateToken(token[0])
		if err != nil {
			return err
		}
		expiresFloat := claims["expires"].(float64)
		expires := int64(expiresFloat)
		// check token expiration
		if time.Now().Unix() > expires {
			return fmt.Errorf("token expired")
		}
		userID := claims["id"].(string) //(primitive.ObjectID)
		user, err := userStore.GetUserByID(c.Context(),userID)
		if err != nil {
			return fmt.Errorf("unauthorized")
		}
		// Set the current authenticated user to the context
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse JWT token:", err)
		return nil, fmt.Errorf("unauthorized")
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}

	return claims, nil
}
