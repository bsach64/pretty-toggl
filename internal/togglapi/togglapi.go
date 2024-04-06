package togglapi

import (
	"net/http"
	"time"
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
