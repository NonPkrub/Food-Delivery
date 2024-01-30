package middleware

import (
	"Food-delivery/domain"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if accessToken == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		fmt.Println(token)
		fmt.Println(err)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized access token",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			c.Locals("email", claims["email"])
			return c.Next()
		}

		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"status":      fiber.ErrUnauthorized.Message,
			"status_code": fiber.ErrUnauthorized.Code,
			"message":     "error, unauthorized",
			"result":      nil,
		})
	}
}

// var users = map[string]string{
// 	"<EMAIL>": "<PASSWORD>",
// }

func BasicAuth(c *fiber.Ctx) error {
	// Get the "Authorization" header
	auth := c.Get("Authorization")
	fmt.Println(auth)
	// Check if the header is empty
	if auth == "" {
		// Unauthorized if no credentials are provided
		c.Status(fiber.StatusUnauthorized)
		return c.SendString("Unauthorized")
	}

	// Parse the header value
	// The format is usually: "Basic BASE64(username:password)"
	// Extract the BASE64 part and decode it
	creds := strings.SplitN(auth, " ", 2)
	if len(creds) != 2 || creds[0] != "Basic" {
		// Invalid Authorization header format
		c.Status(fiber.StatusBadRequest)
		return c.SendString("Bad Request")
	}

	// Decode the BASE64
	payload, err := base64.StdEncoding.DecodeString(creds[1])
	if err != nil {
		// Error decoding BASE64
		c.Status(fiber.StatusBadRequest)
		return c.SendString("Bad Request")
	}

	// Extract username and password
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		// Invalid username:password format
		c.Status(fiber.StatusBadRequest)
		return c.SendString("Bad Request")
	}

	loginForm := domain.UserLoginForm{
		Email:    pair[0],
		Password: pair[1],
	}

	// Check if the provided credentials are valid
	if isValidUser(loginForm) {
		// Invalid email or password
		c.Status(fiber.StatusUnauthorized)
		return c.SendString("Unauthorized Invalid email or password")
	}

	// Continue processing if authentication is successful
	return c.Next()
}

func isValidUser(loginForm domain.UserLoginForm) bool {
	return loginForm.Email == "<EMAIL>" && loginForm.Password == "<PASSWORD>"
}
