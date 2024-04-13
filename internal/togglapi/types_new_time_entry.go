package togglapi

type NewTimeEntry struct {
	Billable bool `json:"billable"`
	CreatedWith string `json:"created_with"`
	Description *string `json:"description"`
	Duration int `json:"duration"`
	ProjectID *int `json:"project_id"`
	SharedWithUserIDs []int `json:"shared_with_user_id"`
	Tags []string `json:"tags"`
	Start string `json:"start"`
	WorkspaceID int `json:"workspace_id"`
}

func CreateNewTimeEntry() NewTimeEntry {
	return NewTimeEntry{
		CreatedWith: "pretty-toggl",
		Duration: -1,
	}
}
