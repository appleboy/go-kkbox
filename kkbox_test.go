package kkbox

import (
	"reflect"
	"testing"
)

func TestBox_getBase64Encode(t *testing.T) {
	type fields struct {
		ClientID     string
		ClientSecret string
		Debug        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "base 64 func",
			fields: fields{
				ClientID:     "client_id",
				ClientSecret: "client_secret",
			},
			want: "Y2xpZW50X2lkOmNsaWVudF9zZWNyZXQ=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				ClientID:     tt.fields.ClientID,
				ClientSecret: tt.fields.ClientSecret,
				Debug:        tt.fields.Debug,
			}
			if got := b.getCredential(); got != tt.want {
				t.Errorf("Box.getBase64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetToken(t *testing.T) {
	type fields struct {
		ClientID     string
		ClientSecret string
		Auth         *AuthData
		Debug        bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    *AuthData
		wantErr bool
	}{
		{
			name: "base 64 func",
			fields: fields{
				ClientID:     "client_id",
				ClientSecret: "client_secret",
			},
			want: &AuthData{
				AccessToken: "",
				ExpiresIn:   0,
				TokenType:   "",
				Error:       "invalid_client",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				ClientID:     tt.fields.ClientID,
				ClientSecret: tt.fields.ClientSecret,
				Auth:         tt.fields.Auth,
				Debug:        tt.fields.Debug,
			}
			got, err := b.GetToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("Box.GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Box.GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
