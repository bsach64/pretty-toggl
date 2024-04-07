package togglapi

import (
	"encoding/json"
	"io"
)

func (c *Client) CurrentTimeEntryReq() (CurrentTimeEntry, error) {
	endpoint := "/me/time_entries/current"
	fullURL := baseURL + endpoint

	req, err := NewGetRequest(fullURL)
	if err != nil {
		return CurrentTimeEntry{}, err
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return CurrentTimeEntry{}, err
	}

	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return CurrentTimeEntry{}, err
	}

	ct := CurrentTimeEntry{}
	err = json.Unmarshal(dat, &ct)
	if err != nil {
		return CurrentTimeEntry{}, err
	}
	return ct, nil
}
