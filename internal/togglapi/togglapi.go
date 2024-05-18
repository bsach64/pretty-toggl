package togglapi

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/bsach64/pretty-toggl/internal/util"
)

const baseURL = "https://api.track.toggl.com/api/v9"

type Client struct {
	HttpClient http.Client
}

func NewClient(interval time.Duration) Client {
	return Client{
		HttpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func AddHeadersAuth(req *http.Request) error {
	req.Header.Add("Content-Type", "application/json")

	token, err := util.ReadAuthToken()
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid API Token: %v", err.Error()))
	}

	req.SetBasicAuth(token, "api_token")
	return nil
}
