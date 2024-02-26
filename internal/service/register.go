package service

import (
	"fido2server/internal/repository"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofiber/fiber/v2"
)

type RegisterService struct {
	UserRepository repository.UserRepository
}

func (r *RegisterService) Register(c *fiber.Ctx) error {
	user := new(webauthn.User)
	r.UserRepository.SaveUser(*user)
	return c.SendStatus(fiber.StatusOK)
}
