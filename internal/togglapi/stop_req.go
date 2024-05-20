package togglapi

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) StopReq(workspaceID, timeEntryID int) (bool, error) {
	endpoint := fmt.Sprintf("/workspaces/%v/time_entries/%v/stop",
		workspaceID, timeEntryID)
	fullURL := baseURL + endpoint

	req, err := http.NewRequest(http.MethodPatch, fullURL, nil)
	if err != nil {
		return false, err
	}
	err = AddHeadersAuth(req)
	if err != nil {
		return false, err
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true, nil
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	return false, nil 
}
