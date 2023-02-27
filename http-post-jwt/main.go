package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type RequestsError struct {
	HTTPCode int
	Body     string
	Err      string
}

func (r RequestsError) Error() string {
	return r.Err
}

type response interface {
	Getresponse() string
}

type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (w Words) Getresponse() string {
	return fmt.Sprintf("%s", strings.Join(w.Words, ", "))
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func (o Occurrence) Getresponse() string {
	out := []string{}
	for word, occurrence := range o.Words {
		out = append(out, fmt.Sprintf("%s (%d)", word, occurrence))
	}
	return fmt.Sprintf("%s", strings.Join(out, ", "))
}

type LoginRequest struct {
	Password string `json:password`
}

type LoginResponse struct {
	Token string `json:token`
}

type MyJWTTransport struct {
	transport http.RoundTripper
	token     string
}

func (m MyJWTTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.token != "" {
		req.Header.Add("Authorization", "Bearer "+m.token)
	}
	return m.transport.RoundTrip(req)
}

func doLoginRequest(client http.Client, requestURL string, password string) (string, error) {
	loginRequest := LoginRequest{
		Password: password,
	}

	body, err := json.Marshal(loginRequest)
	if err != nil {
		return "", fmt.Errorf("Marshal error: %s", err)
	}

	response, err := client.Post(requestURL, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return "", fmt.Errorf("httpPOST error: %s", err)
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)

	if err != nil {
		return "", fmt.Errorf("ReadAll error: %s", err)
	}

	if response.StatusCode != 200 {
		return "", fmt.Errorf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, string(resBody))
	}

	if !json.Valid(resBody) {
		return "", RequestsError{
			HTTPCode: response.StatusCode,
			Body:     string(resBody),
			Err:      fmt.Sprintf("No valid JSON returned"),
		}
	}

	var loginResponse LoginResponse

	err = json.Unmarshal(resBody, &loginResponse)
	if err != nil {
		return "", RequestsError{
			HTTPCode: response.StatusCode,
			Body:     string(resBody),
			Err:      fmt.Sprintf("Page unmarshal error: %s", err),
		}
	}

	return loginResponse.Token, nil
}

func main() {
	var (
		requestURL string
		password   string
		parsedURL  *url.URL
		err        error
	)

	flag.StringVar(&requestURL, "url", "", "url to access")
	flag.StringVar(&password, "password", "", "password to use api")

	flag.Parse()

	if parsedURL, err = url.ParseRequestURI(requestURL); err != nil {
		fmt.Printf("Validation error: URL is not valid %s\n", err)
		flag.Usage()
		os.Exit(1)
	}

	client := http.Client{}

	if password != "" {
		token, err := doLoginRequest(client, parsedURL.Scheme+"://"+parsedURL.Host+"/login", password)
		if err != nil {
			if requestError, ok := err.(RequestsError); ok {
				fmt.Printf("Error: %s, (HTTP Code: %d, Body: %s)\n", requestError.Err, requestError.HTTPCode, requestError.Body)
				os.Exit(1)
			}
		}
		client.Transport = MyJWTTransport{
			transport: http.DefaultTransport,
			token:     token,
		}
	}

	res, err := doRequest(client, parsedURL.String())
	if err != nil {
		if requestError, ok := err.(RequestsError); ok {
			fmt.Printf("Error: %s, (HTTP Code: %d, Body: %s)\n", requestError.Err, requestError.HTTPCode, requestError.Body)
			os.Exit(1)
		}
	}

	if res == nil {
		fmt.Printf("No response")
		os.Exit(1)
	}

	fmt.Printf("response: %s", res.Getresponse())

}

func doRequest(client http.Client, requestURL string) (response, error) {

	response, err := client.Get(requestURL)

	if err != nil {
		return nil, fmt.Errorf("httpGet error: %s", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("ReadAll error: %s", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, string(body))
	}

	if !json.Valid(body) {
		return nil, RequestsError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Sprintf("No valid JSON returned"),
		}
	}

	var page Page

	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, RequestsError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Sprintf("Page unmarshal error: %s", err),
		}
	}

	switch page.Name {
	case "words":
		var words Words

		err = json.Unmarshal(body, &words)
		if err != nil {
			return nil, RequestsError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Words unmarshal error: %s", err),
			}
		}
		return words, nil
	case "occurrence":
		var occurrence Occurrence

		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			return nil, RequestsError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Occurrence unmarshal error: %s", err),
			}
		}
		return occurrence, nil
	}

	return nil, nil
}
