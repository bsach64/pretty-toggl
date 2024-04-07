package togglapi

import (
	"time"
)


type Me struct {
	ID                 int       `json:"id"`
	APIToken           string    `json:"api_token"`
	Email              string    `json:"email"`
	Fullname           string    `json:"fullname"`
	Timezone           string    `json:"timezone"`
	TogglAccountsID    string    `json:"toggl_accounts_id"`
	DefaultWorkspaceID int       `json:"default_workspace_id"`
	BeginningOfWeek    int       `json:"beginning_of_week"`
	ImageURL           string    `json:"image_url"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	OpenidEmail        any       `json:"openid_email"`
	OpenidEnabled      bool      `json:"openid_enabled"`
	CountryID          int       `json:"country_id"`
	HasPassword        bool      `json:"has_password"`
	At                 time.Time `json:"at"`
	IntercomHash       string    `json:"intercom_hash"`
	OauthProviders     []string  `json:"oauth_providers"`
	Tags               []struct {
		ID          int       `json:"id"`
		WorkspaceID int       `json:"workspace_id"`
		Name        string    `json:"name"`
		At          time.Time `json:"at"`
		CreatorID   int       `json:"creator_id"`
		Permissions any       `json:"permissions"`
	} `json:"tags"`
	Projects []struct {
		ID                  int       `json:"id"`
		WorkspaceID         int       `json:"workspace_id"`
		ClientID            any       `json:"client_id"`
		Name                string    `json:"name"`
		IsPrivate           bool      `json:"is_private"`
		Active              bool      `json:"active"`
		At                  time.Time `json:"at"`
		CreatedAt           time.Time `json:"created_at"`
		ServerDeletedAt     any       `json:"server_deleted_at"`
		Color               string    `json:"color"`
		Billable            any       `json:"billable"`
		Template            any       `json:"template"`
		AutoEstimates       any       `json:"auto_estimates"`
		EstimatedHours      any       `json:"estimated_hours"`
		EstimatedSeconds    any       `json:"estimated_seconds"`
		Rate                any       `json:"rate"`
		RateLastUpdated     any       `json:"rate_last_updated"`
		Currency            any       `json:"currency"`
		Recurring           bool      `json:"recurring"`
		TemplateID          any       `json:"template_id"`
		RecurringParameters any       `json:"recurring_parameters"`
		FixedFee            any       `json:"fixed_fee"`
		ActualHours         int       `json:"actual_hours"`
		ActualSeconds       int       `json:"actual_seconds"`
		StartDate           string    `json:"start_date"`
		Status              string    `json:"status"`
		Wid                 int       `json:"wid"`
		Cid                 any       `json:"cid"`
		Permissions         any       `json:"permissions"`
	} `json:"projects"`
	Workspaces []struct {
		ID                          int       `json:"id"`
		OrganizationID              int       `json:"organization_id"`
		Name                        string    `json:"name"`
		Profile                     int       `json:"profile"`
		Premium                     bool      `json:"premium"`
		BusinessWs                  bool      `json:"business_ws"`
		Admin                       bool      `json:"admin"`
		Role                        string    `json:"role"`
		SuspendedAt                 any       `json:"suspended_at"`
		ServerDeletedAt             any       `json:"server_deleted_at"`
		DefaultHourlyRate           any       `json:"default_hourly_rate"`
		RateLastUpdated             any       `json:"rate_last_updated"`
		DefaultCurrency             string    `json:"default_currency"`
		OnlyAdminsMayCreateProjects bool      `json:"only_admins_may_create_projects"`
		OnlyAdminsMayCreateTags     bool      `json:"only_admins_may_create_tags"`
		OnlyAdminsSeeBillableRates  bool      `json:"only_admins_see_billable_rates"`
		OnlyAdminsSeeTeamDashboard  bool      `json:"only_admins_see_team_dashboard"`
		ProjectsBillableByDefault   bool      `json:"projects_billable_by_default"`
		ProjectsPrivateByDefault    bool      `json:"projects_private_by_default"`
		LastModified                time.Time `json:"last_modified"`
		ReportsCollapse             bool      `json:"reports_collapse"`
		Rounding                    int       `json:"rounding"`
		RoundingMinutes             int       `json:"rounding_minutes"`
		APIToken                    string    `json:"api_token"`
		At                          time.Time `json:"at"`
		LogoURL                     string    `json:"logo_url"`
		IcalURL                     string    `json:"ical_url"`
		IcalEnabled                 bool      `json:"ical_enabled"`
		CsvUpload                   any       `json:"csv_upload"`
		Subscription                any       `json:"subscription"`
		HideStartEndTimes           bool      `json:"hide_start_end_times"`
		WorkingHoursInMinutes       any       `json:"working_hours_in_minutes"`
		Permissions                 any       `json:"permissions"`
	} `json:"workspaces"`
	AuthorizationUpdatedAt time.Time `json:"authorization_updated_at"`
}
