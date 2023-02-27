package main

import (
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

type Responce interface {
	GetResponce() string
}

type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (w Words) GetResponce() string {
	return fmt.Sprintf("%s", strings.Join(w.Words, ", "))
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func (o Occurrence) GetResponce() string {
	out := []string{}
	for word, occurrence := range o.Words {
		out = append(out, fmt.Sprintf("%s (%d)", word, occurrence))
	}
	return fmt.Sprintf("%s", strings.Join(out, ", "))
}

func doLoginRequest(loginUrl string, password string) (string, error) {
	return "", nil
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

	if password != "" {
		token, err := doLoginRequest(parsedURL.Scheme+"://"+parsedURL.Host+"/login", password)
		if err != nil {
			if requestError, ok := err.(RequestsError); ok {
				fmt.Printf("Error: %s, (HTTP Code: %d, Body: %s)\n", requestError.Err, requestError.HTTPCode, requestError.Body)
				os.Exit(1)
			}
		}
		fmt.Printf("Token: %s", token)
		os.Exit(1)
	}

	res, err := doRequest(parsedURL.String())
	if err != nil {
		if requestError, ok := err.(RequestsError); ok {
			fmt.Printf("Error: %s, (HTTP Code: %d, Body: %s)\n", requestError.Err, requestError.HTTPCode, requestError.Body)
			os.Exit(1)
		}
	}

	if res == nil {
		fmt.Printf("No responce")
		os.Exit(1)
	}

	fmt.Printf("Responce: %s", res.GetResponce())

}

func doRequest(requestURL string) (Responce, error) {

	if _, err := url.ParseRequestURI(requestURL); err != nil {
		return nil, fmt.Errorf("Validation error: URL is not valid %s", err)
	}

	response, err := http.Get(requestURL)

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
