package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type MockRoundTripper struct {
	RoundTripperOutput *http.Response
}

func (m MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripperOutput, nil
}

func TestRoundTrip(t *testing.T) {
	loginRespone := LoginResponse{
		Token: "abc",
	}
	loginResponeBytes, err := json.Marshal(loginRespone)
	if err != nil {
		t.Errorf("Marshal error: %s", err)
	}
	myJWTTransport := MyJWTTransport{
		transport: MockRoundTripper{
			RoundTripperOutput: &http.Response{
				StatusCode: 200,
			},
		},
		HTTPClient: MockClient{
			PostResponeOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(loginResponeBytes)),
			},
		},
	}
	req := &http.Request{}
	res, err := myJWTTransport.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip error: %s", err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Status codeis not 200, got %d", res.StatusCode)
	}

}
