package kkbox

import (
	"encoding/base64"
	"errors"
)

// Box struct
type Box struct {
	ClientID     string
	ClientSecret string
	Debug        bool
}

// AuthRes for access token
type AuthRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// New MyAllocator object
func New(id, secret string) (*Box, error) {
	if id == "" || secret == "" {
		return nil, errors.New("missing client id or secret")
	}

	box := &Box{
		ClientID:     id,
		ClientSecret: secret,
	}

	return box, nil
}

// GetToken get access token
func (b *Box) GetToken() string {
	data := []byte(b.ClientID + ":" + b.ClientSecret)
	certificate := base64.StdEncoding.EncodeToString(data)

	return certificate
}
