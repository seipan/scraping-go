package scrape

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func DoAPI(method, path string, values url.Values, body io.Reader) ([]byte, error) {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = values.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
