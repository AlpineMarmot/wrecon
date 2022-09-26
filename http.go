package wrecon

import (
	"crypto/tls"
	"net/http"
)

var (
	client = http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{},
		},
	}
)

func Request(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Accept":     {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
		"User-Agent": {"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Mobile Safari/537.36"},
	}

	return client.Do(req)
}
