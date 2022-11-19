package github

import (
	"fmt"
	"net/http"
)

type Release struct {
	Name string `json:"name"`
}

func (c *Client) GetLatestRelease(owner, repo string) (*Release, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/repos/%s/%s/releases/latest", owner, repo), nil)
	if err != nil {
		return nil, err
	}

	var r Release
	if err := c.do(req, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
