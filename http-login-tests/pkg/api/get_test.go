package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockClient struct {
	GetResponceOutput *http.Response
	PostResponeOutput *http.Response
}

func (m MockClient) Get(url string) (resp *http.Response, err error) {
	return m.GetResponceOutput, nil
}

func (m MockClient) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return m.PostResponeOutput, nil
}

func TestDoGetRequest(t *testing.T) {
	words := WordsPage{
		Page: Page{"words"},
		Words: Words{
			Input: "abc",
			Words: []string{"a", "b"},
		},
	}

	wordsBytes, err := json.Marshal(words)
	if err != nil {
		t.Errorf("Marshal error: %s", err)
	}

	apiInstance := api{
		Options: Options{},
		Client: MockClient{
			GetResponceOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(wordsBytes)),
			},
		},
	}
	responce, err := apiInstance.DoGetRequest("http://localhost/words")
	if err != nil {
		t.Error("DoGetRequest error: $s", err)
	}
	if responce == nil {
		t.Fatalf("Respomce is empty")
	}
	if responce.GetResponse() != strings.Join([]string{"a", "b"}, ", ") {
		t.Errorf("Unxepected responce: %s", responce.GetResponse())
	}
}
