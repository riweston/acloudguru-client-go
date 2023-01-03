package acg

import (
	"net/http"
	"net/http/httptest"
	_ "net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestClient_GetSubscription(t *testing.T) {
	type fields struct {
		BaseUrl    string
		HTTPClient *http.Client
		Auth       HeaderStruct
	}
	tests := []struct {
		name    string
		server  *httptest.Server
		fields  fields
		want    *Subscription
		wantErr bool
	}{
		{
			name: "Return a good response",
			server: httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusOK)
				writer.Write([]byte(`{
				  "organisationId": "123a1b2c3-a1s2d3a3sd-1a543123-121212312",
				  "name": "Hello Cloud Gurus",
				  "startDate": "2017-10-06T03:04:06.000000722Z",
				  "endDate": "2018-10-06T03:04:06.000000722Z",
				  "totalSeats": 1337,
				  "seatsInUse": 101
				}`))
			})),
			fields: fields{
				HTTPClient: &http.Client{},
				Auth: HeaderStruct{
					apiKey:     "test",
					consumerId: "test",
				},
			},
			want: &Subscription{
				OrganisationId: "123a1b2c3-a1s2d3a3sd-1a543123-121212312",
				Name:           "Hello Cloud Gurus",
				StartDate:      time.Date(2017, 10, 06, 03, 04, 06, 722, time.UTC),
				EndDate:        time.Date(2018, 10, 06, 03, 04, 06, 722, time.UTC),
				TotalSeats:     1337,
				SeatsInUse:     101,
			},
			wantErr: false,
		},
		{
			name: "Bad auth tokens",
			server: httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusForbidden)
				writer.Write([]byte(`{"message":"Forbidden"}`))
			})),
			fields: fields{
				HTTPClient: &http.Client{},
				Auth: HeaderStruct{
					apiKey:     "test",
					consumerId: "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.server.Close()
			c := &Client{
				BaseUrl:    tt.server.URL,
				HTTPClient: tt.fields.HTTPClient,
				Auth:       tt.fields.Auth,
			}
			got, err := c.GetSubscription()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubscription() got = %v, want %v", got, tt.want)
			}
		})
	}
}
