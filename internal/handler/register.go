package handler

import (
	"fido2server/internal/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterOptionsHandler(service service.RegisterOptionsService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return service.RegisterOptions(c)
	}
}

func RegisterResultHandler(service service.RegisterResultService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return service.RegisterResult(c)
	}
}
