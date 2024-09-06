package togglapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) CreateProject(projectName, color string) error {
	me, err := c.MeReq()
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/workspaces/%v/projects", me.DefaultWorkspaceID)
	fullURL := baseURL + endpoint

	details := struct {
		Active      bool   `json:"active"`
		ProjectName string `json:"name"`
		Color       string `json:"color"`
	}{true, projectName, color}

	data, err := json.Marshal(details)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fullURL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	err = AddHeadersAuth(req)
	if err != nil {
		return err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("Could not create project")
	}
	return nil
}
