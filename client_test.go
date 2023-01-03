package acg

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_doRequest(t *testing.T) {
	type fields struct {
		BaseUrl    string
		HTTPClient *http.Client
		Auth       HeaderStruct
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				BaseUrl:    tt.fields.BaseUrl,
				HTTPClient: tt.fields.HTTPClient,
				Auth:       tt.fields.Auth,
			}
			got, err := c.doRequest(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_newRequest(t *testing.T) {
	type fields struct {
		BaseUrl    string
		HTTPClient *http.Client
		Auth       HeaderStruct
	}
	type args struct {
		requestMethod string
		requestPath   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				BaseUrl:    tt.fields.BaseUrl,
				HTTPClient: tt.fields.HTTPClient,
				Auth:       tt.fields.Auth,
			}
			got, err := c.newRequest(tt.args.requestMethod, tt.args.requestPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("newRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		apiKey     *string
		consumerId *string
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.apiKey, tt.args.consumerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
