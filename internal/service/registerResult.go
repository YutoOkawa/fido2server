package service

import (
	"encoding/json"
	"fido2server/internal/repository"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type RegisterResultService struct {
	UserRepository    repository.UserRepository
	SessionRepository repository.SessionDataRepository
	WebAuthn          *webauthn.WebAuthn
}

type TmpRequestParams struct {
	UserName string `json:"userName"`
}

func (r *RegisterResultService) RegisterResult(c *fiber.Ctx) error {
	body := c.Body()

	var params TmpRequestParams
	if err := json.Unmarshal(body, &params); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user, err := r.UserRepository.GetUser(params.UserName)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	sessionData, err := r.SessionRepository.GetSessionData(params.UserName)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	request, err := adaptor.ConvertRequest(c, false)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	credential, err := r.WebAuthn.FinishRegistration(user, *sessionData, request)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	user.AddCredential(*credential)
	if err := r.UserRepository.SaveUser(user); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// TODO: 登録処理
	return c.SendStatus(fiber.StatusOK)
}
