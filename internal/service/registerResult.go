package service

import (
	"fido2server/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type RegisterResultService struct {
	UserRepository repository.UserRepository
}

func (r *RegisterResultService) RegisterResult(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)

}
