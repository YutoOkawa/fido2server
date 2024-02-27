package service

import (
	"fido2server/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type RegisterOptionsService struct {
	UserRepository repository.UserRepository
}

func (r *RegisterOptionsService) RegisterOptions(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
