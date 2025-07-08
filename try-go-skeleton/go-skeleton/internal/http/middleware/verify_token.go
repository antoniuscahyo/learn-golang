package middleware

import (
	apperr "github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/error"
	"github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/internal/http/auth"
	"github.com/gofiber/fiber/v2"
)

func VerifyJWTToken(c *fiber.Ctx) error {
	if err := auth.VerifyToken(c); err != nil {
		return c.Status(apperr.ErrInvalidToken().HTTPCode).JSON(apperr.ErrInvalidToken())
	}

	return c.Next()
}
