package togglapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/bsach64/pretty-toggl/internal/util"
)

func (c *Client) MeReq() (Me, error) {
	endpoint := "/me?with_related_data=true"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return Me{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	token, valid := util.ReadAuthFromEnv()
	if !valid {
		return Me{}, errors.New("Could not get auth token")
	}

	req.SetBasicAuth(token, "api_token")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return Me{}, err
	}

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
