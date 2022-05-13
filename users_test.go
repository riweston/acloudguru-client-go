package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestClient_GetUserFromEmail(t *testing.T) {
	type fields struct {
		BaseUrl    string
		HTTPClient *http.Client
		Auth       HeaderStruct
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		server  *httptest.Server
		fields  fields
		args    args
		want    *[]User
		wantErr bool
	}{
		{
			name: "Return a good response",
			server: httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusOK)
				writer.Write([]byte(`[
				  {
					"userId": "5e1dd7dc-2371-4ee1-a188-d14fd00ee275",
					"name": "Fred Flinstone",
					"email": "fred@flinstones.net",
					"admin": false,
					"lastSeenTimestamp": "2019-03-19T00:00:00.000Z",
					"status": "Active"
				  }
				]`))
			})),
			fields: fields{
				HTTPClient: &http.Client{},
				Auth: HeaderStruct{
					apiKey:     "test",
					consumerId: "test",
				},
			},
			args: args{
				email: "test@test.com",
			},
			want: &[]User{
				{
					UserId:            "5e1dd7dc-2371-4ee1-a188-d14fd00ee275",
					Name:              "Fred Flinstone",
					Email:             "fred@flinstones.net",
					Admin:             false,
					LastSeenTimestamp: time.Date(2019, 03, 19, 00, 00, 00, 000, time.UTC),
					Status:            "Active",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				BaseUrl:    tt.server.URL,
				HTTPClient: tt.fields.HTTPClient,
				Auth:       tt.fields.Auth,
			}
			got, err := c.GetUserFromEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserFromEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserFromEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetUsers(t *testing.T) {
	type fields struct {
		BaseUrl    string
		HTTPClient *http.Client
		Auth       HeaderStruct
	}
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		server  *httptest.Server
		fields  fields
		args    args
		want    *[]User
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
			got, err := c.GetUsers(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetUsersAll(t *testing.T) {
	type fields struct {
		BaseUrl    string
		HTTPClient *http.Client
		Auth       HeaderStruct
	}
	tests := []struct {
		name    string
		server  *httptest.Server
		fields  fields
		want    *[]User
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
			got, err := c.GetUsersAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsersAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsersAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetUserActivated(t *testing.T) {
	type fields struct {
		BaseUrl    string
		HTTPClient *http.Client
		Auth       HeaderStruct
	}
	type args struct {
		user     *User
		activate bool
	}
	tests := []struct {
		name    string
		server  *httptest.Server
		fields  fields
		args    args
		want    *Response
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
			got, err := c.SetUserActivated(tt.args.user, tt.args.activate)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetUserActivated() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUserActivated() got = %v, want %v", got, tt.want)
			}
		})
	}
}
