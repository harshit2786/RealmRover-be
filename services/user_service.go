package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"realmrovers/config"
	"realmrovers/model"

	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
	Cfg	*config.Config
}

func (u *UserService) GetUsers(id uint) (*model.User, error) {
	var user model.User
	err := u.Db.Take(&user, id).Error
	if(err != nil ){
		return nil,err
	}
	return &user,nil
}

func (u *UserService) SignUser(code string) (string,error) {
	githubTokenURL := "https://github.com/login/oauth/access_token"

	// Prepare the request body
	payload := map[string]string{
		"client_id":     u.Cfg.GITCLI,
		"client_secret": u.Cfg.GITPASS,
		"code":          code,
		"redirect_uri":  u.Cfg.REDURI,
	}
	body,err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}
	req, err := http.NewRequest("POST", githubTokenURL, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 response from GitHub: %s", resp.Status)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
	if err := json.Unmarshal(respBody, &tokenResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if tokenResponse.AccessToken == "" {
		return "", errors.New("no access token in response")
	}

	return tokenResponse.AccessToken, nil
}