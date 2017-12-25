package kkbox

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"strconv"

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

func (b *Box) fetchData(url string, res interface{}, params ...Param) error {
	req := httplib.Get(url).Debug(b.Debug)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	offset := 0
	limit := 10
	territory := "TW"
	q := ""
	contentType := ""
	if len(params) > 0 {
		param := params[0]
		if param.PerPage > 0 {
			limit = param.PerPage
		}
		if param.Page > 1 {
			offset = (param.Page - 1) * limit
		}
		if param.Territory != "" {
			territory = param.Territory
		}
		if param.Q != "" {
			q = param.Q
		}
		if param.Type != "" {
			contentType = param.Type
		}
	}

	// Add authorization header
	authorization := b.Auth.TokenType + " " + b.Auth.AccessToken
	req.Header("Authorization", authorization)

	req.Param("offset", strconv.Itoa(offset))
	req.Param("limit", strconv.Itoa(limit))
	req.Param("territory", territory)

	if q != "" {
		req.Param("q", q)
	}
	if contentType != "" {
		req.Param("type", contentType)
	}
	return req.ToJSON(res)
}
