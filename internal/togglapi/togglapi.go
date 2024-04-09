package togglapi

import (
	"github.com/bsach64/pretty-toggl/internal/util"
	"net/http"
	"time"
)

const baseURL = "https://api.track.toggl.com/api/v9"

type InvalidAPIToken struct{}

func (e *InvalidAPIToken) Error() string {
	return "Invalid API Token"
}

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

func NewGetRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	err = AddHeadersAuth(req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func AddHeadersAuth(req *http.Request) error {
	req.Header.Add("Content-Type", "application/json")

	token, valid := util.ReadAuthFromEnv()
	if !valid {
		return &InvalidAPIToken{}
	}

	req.SetBasicAuth(token, "api_token")
	return nil
}
