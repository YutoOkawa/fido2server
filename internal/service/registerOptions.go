package service

import (
	"encoding/json"
	"errors"
	"fido2server/internal/repository"
	webauthnlib "fido2server/pkg/webauthn"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofiber/fiber/v2"
)

type RequestParam struct {
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
	Icon        string `json:"icon"`
}

type RegisterOptionsService struct {
	UserRepository        repository.UserRepository
	SessionDataRepository repository.SessionDataRepository
	WebAuthn              *webauthn.WebAuthn
}

func (r *RegisterOptionsService) RegisterOptions(c *fiber.Ctx) error {
	body := c.Body()

	var params RequestParam
	if err := json.Unmarshal(body, &params); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := validateRequestParam(params); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var user *webauthnlib.RegisteredUser
	var err error
	user, err = r.UserRepository.GetUser(params.UserName)
	if !errors.Is(err, webauthnlib.ErrRegisterUserNotFound) && err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if errors.Is(err, webauthnlib.ErrRegisterUserNotFound) {
		user, err = webauthnlib.NewRegisteredUser(params.UserName, params.DisplayName, params.Icon)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if err := r.UserRepository.SaveUser(user); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	options, sessionData, err := r.WebAuthn.BeginRegistration(user)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := r.SessionDataRepository.SaveSessionData(sessionData, user.UserName); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	optionsBytes, err := json.Marshal(options)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Send(optionsBytes)
}

// TODO: sanitize
func validateRequestParam(params RequestParam) error {
	if params.UserName == "" {
		return errors.New("userName is required")
	}
	if params.DisplayName == "" {
		return errors.New("displayName is required")
	}
	return nil
}
