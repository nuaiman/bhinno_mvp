package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GoogleOAuthTokenInfo holds Google access token info
type GoogleOAuthTokenInfo struct {
	Aud           string `json:"aud"`
	UserID        string `json:"sub"`
	ExpiresIn     string `json:"expires_in"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
}

// VerifyGoogleOAuthAccessToken verifies a Google access token via Google tokeninfo endpoint
func VerifyGoogleOAuthAccessToken(accessToken string) (*GoogleOAuthTokenInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=%s", accessToken))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid access token, status: %d", resp.StatusCode)
	}

	var info GoogleOAuthTokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}

	return &info, nil
}
