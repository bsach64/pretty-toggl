package togglapi

import (
	"net/http"
)

func (c *Client) AuthUsingToken(token string) bool {
	endpoint := "/me"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return false
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(token, "api_token")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true
	}

	return false
}
