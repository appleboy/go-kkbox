package kkbox

import (
	"crypto/tls"
	"encoding/base64"
	"errors"

	"github.com/astaxie/beego/httplib"
)

// Box struct
type Box struct {
	ClientID     string
	ClientSecret string
	Auth         *AuthData
	Debug        bool
}

// AuthData for access token
type AuthData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Error       string `json:"error"`
}

var (
	// ErrorMissingIDorSecret for missing client id or secret
	ErrorMissingIDorSecret = errors.New("missing client id or secret")
	// ErrorInvalidClient for wrong token
	ErrorInvalidClient = errors.New("invalid_client")
)

// New MyAllocator object
func New(id, secret string) (*Box, error) {
	if id == "" || secret == "" {
		return nil, ErrorMissingIDorSecret
	}

	box := &Box{
		ClientID:     id,
		ClientSecret: secret,
	}

	// get token
	auth, err := box.GetToken()
	box.Auth = auth

	if auth.Error != "" {
		return nil, ErrorInvalidClient
	}

	return box, err
}

func (b *Box) getCredential() string {
	data := []byte(b.ClientID + ":" + b.ClientSecret)
	return base64.StdEncoding.EncodeToString(data)
}

// GetToken get access token
func (b *Box) GetToken() (*AuthData, error) {
	req := httplib.Post(OAuthTokenURL).Debug(b.Debug)
	req.Param("grant_type", "client_credentials")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("Authorization", "Basic "+b.getCredential())
	res := new(AuthData)
	err := req.ToJSON(res)
	return res, err
}
