package togglapi

import "time"

type CurrentTimeEntry struct {
	ID              int64      `json:"id"`
	WorkspaceID     int        `json:"workspace_id"`
	ProjectID       *int       `json:"project_id"`
	TaskID          *int       `json:"task_id"`
	Billable        bool       `json:"billable"`
	Start           time.Time  `json:"start"`
	Stop            *time.Time `json:"stop"`
	Duration        int        `json:"duration"`
	Description     string     `json:"description"`
	Tags            []string   `json:"tags"`
	TagIds          []int      `json:"tag_ids"`
	Duronly         bool       `json:"duronly"`
	At              time.Time  `json:"at"`
	ServerDeletedAt *time.Time `json:"server_deleted_at"`
	UserID          int        `json:"user_id"`
	UID             int        `json:"uid"`
	Wid             int        `json:"wid"`
	Permissions     []string   `json:"permissions"`
}
