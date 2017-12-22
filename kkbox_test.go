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

func TestNew(t *testing.T) {
	type args struct {
		id     string
		secret string
	}
	tests := []struct {
		name     string
		args     args
		want     *Box
		wantErr  bool
		errorMsg error
	}{
		{
			name: "missing id",
			args: args{
				id:     "",
				secret: "1234",
			},
			want:     nil,
			wantErr:  true,
			errorMsg: ErrorMissingIDorSecret,
		},
		{
			name: "missing secret",
			args: args{
				id:     "1234",
				secret: "",
			},
			want:     nil,
			wantErr:  true,
			errorMsg: ErrorMissingIDorSecret,
		},
		{
			name: "wrong id and secret",
			args: args{
				id:     "1234",
				secret: "1234",
			},
			want: &Box{
				ClientID:     "1234",
				ClientSecret: "1234",
				Auth: &AuthData{
					AccessToken: "",
					ExpiresIn:   0,
					TokenType:   "",
					Error:       "",
				},
				Debug: false,
			},
			wantErr:  true,
			errorMsg: ErrorInvalidClient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.id, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err != tt.errorMsg {
				t.Errorf("New() error = %v, wantErr %v", err, tt.errorMsg)
				return
			}

			if got != nil {
				if got.ClientID != tt.want.ClientID {
					t.Errorf("ClientID = %v, want %v", got.ClientID, tt.want.ClientID)
					return
				}

				if got.ClientSecret != tt.want.ClientSecret {
					t.Errorf("ClientSecret = %v, want %v", got.ClientSecret, tt.want.ClientSecret)
					return
				}
			}
		})
	}
}
