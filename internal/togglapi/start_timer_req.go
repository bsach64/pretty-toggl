package togglapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) StartTimeEntry(timeEntry NewTimeEntry) (bool, error) {
	endpoint := fmt.Sprintf("/workspaces/%v/time_entries", timeEntry.WorkspaceID)
	fullURL := baseURL + endpoint

	data, err := json.Marshal(timeEntry)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest(http.MethodPost, fullURL, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", "application/json")
	err = AddHeadersAuth(req)
	if err != nil {
		return false, err
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	if resp.StatusCode == 200 {
		fmt.Println(resp.StatusCode)
		return true, nil
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	return false, errors.New(string(dat))
}
