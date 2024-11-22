package auth

import (
	"database/sql"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(app *fiber.App, db *sql.DB) {
	app.Use("/", func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/auth") {
			return c.Next()
		}

		authHeader := c.Get("Authorization")
        
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            return c.SendStatus(fiber.StatusUnauthorized)
        }

        token := authHeader[7:]
		if token == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		userId, err := ParseToken(token)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		var isUserValid bool

		checkUserValidityQuery := "SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)"

		err = db.QueryRow(checkUserValidityQuery, userId).Scan(&isUserValid)
		if err != nil {
			log.Printf("Error checking user validity: %v", err)

			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if !isUserValid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Locals("userId", userId)

		return c.Next()
	})
}
