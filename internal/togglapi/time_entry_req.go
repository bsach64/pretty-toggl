package togglapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) CurrentTimeEntryReq() (CurrentTimeEntry, error) {
	endpoint := "/me/time_entries/current"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return CurrentTimeEntry{}, err
	}
	err = AddHeadersAuth(req)
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
	if string(dat) == "null" {
		return CurrentTimeEntry{}, errors.New("No Running Time Entry!")
	}
	ct := CurrentTimeEntry{}
	err = json.Unmarshal(dat, &ct)
	if err != nil {
		return CurrentTimeEntry{}, err
	}
	return ct, nil
}
