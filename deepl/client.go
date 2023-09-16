package deepl

import (
	"fmt"
	"io"
	"net/http"
)

const deepLBaseURL = "https://api-free.deepl.com"

type deepLClient struct {
	apiKey string
}

func DeepLClient(apiKey string) deepLClient {
	return deepLClient{apiKey}
}

func requestDeepL(
	method string,
	url deepLEndpoint,
	apiKey string,
	body io.Reader,
) (*http.Response, error) {
	req, err := http.NewRequest(method, string(url), body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("DeepL-Auth-Key %s", apiKey))
	req.Header.Add("Content-Type", "application/json")

	return http.DefaultClient.Do(req)
}

type deepLEndpoint string

func deepLURL(path string) deepLEndpoint {
	return deepLEndpoint(fmt.Sprintf("%s/%s", deepLBaseURL, path))
}
