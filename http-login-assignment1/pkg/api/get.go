package api

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Response interface {
	GetResponse() string
}

type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

type WordsPage struct {
	Page
	Words
}

func (w Words) GetResponse() string {
	return fmt.Sprintf("%s", strings.Join(w.Words, ", "))
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func (o Occurrence) GetResponse() string {
	out := []string{}
	for word, occurrence := range o.Words {
		out = append(out, fmt.Sprintf("%s (%d)", word, occurrence))
	}
	return fmt.Sprintf("%s", strings.Join(out, ", "))
}

type Assignment struct {
	Percentages map[string]float64 `json:"percentages"`
	Words       []string           `json:"words"`
	Special     []string           `json:"special"`
	Extra       []string           `json:"extraSpecial"`
}

func (a Assignment) GetResponse() string {
	out := []string{}
	for k, v := range a.Percentages {
		out = append(out, fmt.Sprintf("%s (%d)", k, v))
	}
	out = append(out, fmt.Sprintf("%s", strings.Join(a.Words, ", ")))
	out = append(out, fmt.Sprintf("%s", strings.Join(a.Special, ", ")))
	out = append(out, fmt.Sprintf("%s", strings.Join(a.Extra, ", ")))
	return fmt.Sprintf("%s", strings.Join(out, ", "))
}

func (a api) DoGetRequest(requestURL string) (Response, error) {

	response, err := a.Client.Get(requestURL)

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
	case "assignment1":
		var assignment Assignment

		err = json.Unmarshal(body, &assignment)
		if err != nil {
			return nil, RequestsError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Assignment unmarshal error: %s", err),
			}
		}
		return assignment, nil
	}
	return nil, nil
}
