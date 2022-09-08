package auth

import (
	"encoding/json"
	"errors"

	"github.com/felipereyel/PROJECT_NAME/pkgs/config"
	"github.com/felipereyel/PROJECT_NAME/pkgs/utils"
)

func authClient(method string, path string, headers utils.Headers) ([]byte, error) {
	baseUrl := config.Config("AUTH_URL")
	client := utils.GetHTTPClient(baseUrl, headers)
	return client(method, path, nil)
}

func GetAuthorInfo(headers utils.Headers) (*AuthInfo, error) {
	res, err := authClient("GET", "PROJECT_NAME", headers)
	if err != nil {
		return nil, err
	}

	var info AuthInfo
	if err := json.Unmarshal(res, &info); err != nil {
		return nil, errors.New("failed to parse auth info: " + err.Error())
	}

	return &info, nil
}
