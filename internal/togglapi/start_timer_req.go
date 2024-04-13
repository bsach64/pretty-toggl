package togglapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) StartTimeEntry(timeEntry NewTimeEntry) bool {
	endpoint := fmt.Sprintf("/workspaces/%v/time_entries", timeEntry.WorkspaceID)
	fullURL := baseURL + endpoint
	fmt.Println(fullURL)
	data, err := json.Marshal(timeEntry)
	if err != nil {
		return false
	}
	req, err := http.NewRequest(http.MethodPost, fullURL, bytes.NewBuffer(data))
	if err != nil {
		return false
	}
	req.Header.Add("Content-Type", "application/json")
	err = AddHeadersAuth(req)
	if err != nil {
		return false
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if resp.StatusCode != 200 {
		log.Println(resp.StatusCode)
		return false
	}
	return true
}
