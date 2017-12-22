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
			if got := b.getBase64Encode(); got != tt.want {
				t.Errorf("Box.getBase64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
