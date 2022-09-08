package handler

import (
	"fmt"

	"github.com/abstra-app/PROJECT_NAME/pkgs/config"
	"github.com/abstra-app/PROJECT_NAME/pkgs/services/auth"
	"github.com/abstra-app/PROJECT_NAME/pkgs/utils"
	"github.com/gofiber/fiber/v2"
)

func InternalAuth(c *fiber.Ctx) error {
	internalHeader := c.Get("Internal-Service")
	if internalHeader == config.Config("INTERNAL_SERVICE_TOKEN") {
		return c.Next()
	}
	return c.SendStatus(403)
}

type authScope string

const (
	HookScope authScope = "hookId"
)

func hookCheck(id string, info *auth.AuthInfo) bool {
	for _, ws := range info.Workspaces {
		for _, hook := range ws.Hooks {
			if hook.Id == id {
				return true
			}
		}
	}
	return false
}

func extractAuthHeaders(c *fiber.Ctx) utils.Headers {
	headers := utils.Headers{}

	if apiAuth := c.Get("Api-Authorization"); apiAuth != "" {
		headers["Api-Authorization"] = apiAuth
	}

	if authorAuth := c.Get("Author-Authorization"); authorAuth != "" {
		headers["Author-Authorization"] = authorAuth
	}

	return headers
}

func scopeIsAuthorized(scope authScope, c *fiber.Ctx) bool {
	headers := extractAuthHeaders(c)
	authInfo, err := auth.GetAuthorInfo(headers)
	if err != nil {
		fmt.Println("failed to get author info: " + err.Error())
		return false
	}

	switch scope {
	case HookScope:
		return hookCheck(c.Params("hookId"), authInfo)
	default:
		return false
	}
}

func ScopeAuth(scope authScope) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if scopeIsAuthorized(scope, c) {
			return c.Next()
		}
		return c.SendStatus(403)
	}
}
