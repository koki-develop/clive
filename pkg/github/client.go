package github

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"path"
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{new(http.Client)}
}

func (c *Client) newRequest(method, pathname string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, "https://api.github.com", body)
	if err != nil {
		return nil, err
	}

	req.URL.Path = path.Join(req.URL.Path, pathname)
	return req, nil
}

func (c *Client) do(req *http.Request, dst interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || 300 <= resp.StatusCode {
		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(buf))
	}

	if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
		return err
	}

	return nil
}
