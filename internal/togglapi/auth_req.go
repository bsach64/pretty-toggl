package togglapi

import (
	"errors"
	"io"
	"net/http"
)

func (c *Client) AuthUsingToken(token string) (bool, error) {
	endpoint := "/me"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(token, "api_token")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true, nil
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	return false, errors.New(string(dat))
}
