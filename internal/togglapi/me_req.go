package togglapi

import (
	"encoding/json"
	"io"
)

func (c *Client) MeReq() (Me, error) {
	endpoint := "/me?with_related_data=true"
	fullURL := baseURL + endpoint

	req, err := NewGetRequest(fullURL)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return Me{}, err
	}

	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Me{}, err
	}

	me := Me{}
	err = json.Unmarshal(dat, &me)
	if err != nil {
		return Me{}, err
	}
	return me, nil
}
