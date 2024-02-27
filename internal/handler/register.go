package handler

import (
	"fido2server/internal/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(rergister service.RegisterService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := rergister.Register(c); err != nil {
			return err
		}
		return nil
	}
}

func RegisterOptionsHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func RegisterResultHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
