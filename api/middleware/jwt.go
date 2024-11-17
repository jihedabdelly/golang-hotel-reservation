package middleware

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error  {
	fmt.Println("--- JWT authing")

	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok {
		return fmt.Errorf("unauthorized")
	}

	if err := parseToken(token[0]); err != nil {
		return err
	}

	fmt.Println("token:", token)
	return nil
}

func parseToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		fmt.Println("Never print secret", secret)
		return []byte(secret), nil
	})
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("unauthorized")
	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims) 
	}

	return fmt.Errorf("unauthorized")
}