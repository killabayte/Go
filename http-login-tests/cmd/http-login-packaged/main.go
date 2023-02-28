package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/killabayte/Go/http-complete-package/pkg/api"
)

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

	apiInstance := api.New(api.Options{
		Password: password,
		LoginURL: parsedURL.Scheme + "://" + parsedURL.Host + "/login",
	})

	res, err := apiInstance.DoGetRequest(parsedURL.String())
	if err != nil {
		if requestError, ok := err.(api.RequestsError); ok {
			fmt.Printf("Error: %s, (HTTP Code: %d, Body: %s)\n", requestError.Err, requestError.HTTPCode, requestError.Body)
			os.Exit(1)
		}
	}

	if res == nil {
		fmt.Printf("No response")
		os.Exit(1)
	}

	fmt.Printf("response: %s", res.GetResponse())

}
