package service

import "github.com/gofiber/fiber/v2"

type RegisterService struct {
}

func (r *RegisterService) Register(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
