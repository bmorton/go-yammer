package schema

type PageCollection struct {
	Pages      []PageReference `json:"pages"`
	References []struct {
		Type                      string `json:"type"`
		ID                        int    `json:"id"`
		Email                     string `json:"email"`
		Name                      string `json:"name"`
		Community                 bool   `json:"community,omitempty"`
		Permalink                 string `json:"permalink,omitempty"`
		WebURL                    string `json:"web_url"`
		ShowUpgradeBanner         bool   `json:"show_upgrade_banner,omitempty"`
		HeaderBackgroundColor     string `json:"header_background_color,omitempty"`
		HeaderTextColor           string `json:"header_text_color,omitempty"`
		NavigationBackgroundColor string `json:"navigation_background_color,omitempty"`
		NavigationTextColor       string `json:"navigation_text_color,omitempty"`
		Paid                      bool   `json:"paid,omitempty"`
		Moderated                 bool   `json:"moderated,omitempty"`
		IsOrgChartEnabled         bool   `json:"is_org_chart_enabled,omitempty"`
		IsGroupEnabled            bool   `json:"is_group_enabled,omitempty"`
		IsChatEnabled             bool   `json:"is_chat_enabled,omitempty"`
		IsTranslationEnabled      bool   `json:"is_translation_enabled,omitempty"`
		CreatedAt                 string `json:"created_at,omitempty"`
		ProfileFieldsConfig       struct {
			EnableWorkHistory bool `json:"enable_work_history"`
			EnableEducation   bool `json:"enable_education"`
			EnableJobTitle    bool `json:"enable_job_title"`
			EnableWorkPhone   bool `json:"enable_work_phone"`
			EnableMobilePhone bool `json:"enable_mobile_phone"`
			EnableBirthday    bool `json:"enable_birthday"`
			EnableSummary     bool `json:"enable_summary"`
			EnableInterests   bool `json:"enable_interests"`
			EnableExpertise   bool `json:"enable_expertise"`
			EnableLocation    bool `json:"enable_location"`
			EnableIm          bool `json:"enable_im"`
			EnableSkype       bool `json:"enable_skype"`
			EnableWebsites    bool `json:"enable_websites"`
		} `json:"profile_fields_config,omitempty"`
		BrowserDeprecationURL  interface{} `json:"browser_deprecation_url,omitempty"`
		ExternalMessagingState string      `json:"external_messaging_state,omitempty"`
		State                  string      `json:"state,omitempty"`
		FullName               string      `json:"full_name,omitempty"`
		NetworkID              int         `json:"network_id,omitempty"`
		Description            string      `json:"description,omitempty"`
		Privacy                string      `json:"privacy,omitempty"`
		URL                    string      `json:"url,omitempty"`
		MugshotURL             string      `json:"mugshot_url,omitempty"`
		MugshotURLTemplate     string      `json:"mugshot_url_template,omitempty"`
		MugshotID              string      `json:"mugshot_id,omitempty"`
		ShowInDirectory        string      `json:"show_in_directory,omitempty"`
		Members                int         `json:"members,omitempty"`
		Color                  string      `json:"color,omitempty"`
		External               bool        `json:"external,omitempty"`
		JobTitle               string      `json:"job_title,omitempty"`
		ActivatedAt            string      `json:"activated_at,omitempty"`
		AutoActivated          bool        `json:"auto_activated,omitempty"`
		Stats                  struct {
			Following int `json:"following"`
			Followers int `json:"followers"`
			Updates   int `json:"updates"`
		} `json:"stats,omitempty"`
	} `json:"references"`
	Meta struct {
		Pagination struct {
			IsLastPage bool `json:"is_last_page"`
		} `json:"pagination"`
		FollowedPageIds []int `json:"followed_page_ids"`
	} `json:"meta"`
	ExternalReferences []interface{} `json:"external_references"`
}

type PageReference struct {
	ID                        int           `json:"id"`
	NetworkID                 int           `json:"network_id"`
	URL                       string        `json:"url"`
	WebURL                    string        `json:"web_url"`
	PreviewURL                string        `json:"preview_url"`
	APIPreviewURL             string        `json:"api_preview_url"`
	Type                      string        `json:"type"`
	Name                      string        `json:"name"`
	FullName                  string        `json:"full_name"`
	Title                     string        `json:"title"`
	GroupID                   int           `json:"group_id"`
	PaddieHost                string        `json:"paddie_host"`
	PaddieID                  string        `json:"paddie_id"`
	OwnerID                   int           `json:"owner_id"`
	OwnerType                 string        `json:"owner_type"`
	Active                    bool          `json:"active"`
	Official                  bool          `json:"official"`
	UpdatedAt                 string        `json:"updated_at"`
	LastPublishedBy           int           `json:"last_published_by"`
	LastPublishedAt           string        `json:"last_published_at"`
	LatestRevision            int           `json:"latest_revision"`
	UnpublishedContributorIds []interface{} `json:"unpublished_contributor_ids"`
	Topics                    []interface{} `json:"topics"`
	Privacy                   string        `json:"privacy"`
	Description               string        `json:"description"`
	LatestVersion             struct {
		ID             int         `json:"id"`
		NetworkID      int         `json:"network_id"`
		Type           string      `json:"type"`
		Title          string      `json:"title"`
		CreatedAt      string      `json:"created_at"`
		WebURL         string      `json:"web_url"`
		DownloadURL    string      `json:"download_url"`
		RevertURL      interface{} `json:"revert_url"`
		PublishedBy    int         `json:"published_by"`
		RevisionNumber int         `json:"revision_number"`
		RevertedID     interface{} `json:"reverted_id"`
		Contributors   []int       `json:"contributors"`
	} `json:"latest_version"`
	Stats struct {
		Followers int `json:"followers"`
	} `json:"stats"`
}

type Page struct {
	ID                        int           `json:"id"`
	NetworkID                 int           `json:"network_id"`
	URL                       string        `json:"url"`
	WebURL                    string        `json:"web_url"`
	PreviewURL                string        `json:"preview_url"`
	APIPreviewURL             string        `json:"api_preview_url"`
	Type                      string        `json:"type"`
	Name                      string        `json:"name"`
	FullName                  string        `json:"full_name"`
	Title                     string        `json:"title"`
	GroupID                   int           `json:"group_id"`
	PaddieHost                string        `json:"paddie_host"`
	PaddieID                  string        `json:"paddie_id"`
	OwnerID                   int           `json:"owner_id"`
	OwnerType                 string        `json:"owner_type"`
	Active                    bool          `json:"active"`
	Official                  bool          `json:"official"`
	UpdatedAt                 string        `json:"updated_at"`
	LastPublishedBy           int           `json:"last_published_by"`
	LastPublishedAt           string        `json:"last_published_at"`
	LatestRevision            int           `json:"latest_revision"`
	UnpublishedContributorIds []interface{} `json:"unpublished_contributor_ids"`
	Topics                    []interface{} `json:"topics"`
	Privacy                   string        `json:"privacy"`
	Description               string        `json:"description"`
	LatestVersion             struct {
		ID             int         `json:"id"`
		NetworkID      int         `json:"network_id"`
		Type           string      `json:"type"`
		Title          string      `json:"title"`
		CreatedAt      string      `json:"created_at"`
		WebURL         string      `json:"web_url"`
		DownloadURL    string      `json:"download_url"`
		RevertURL      interface{} `json:"revert_url"`
		PublishedBy    int         `json:"published_by"`
		RevisionNumber int         `json:"revision_number"`
		RevertedID     interface{} `json:"reverted_id"`
		Contributors   []int       `json:"contributors"`
	} `json:"latest_version"`
	Stats struct {
		Followers int `json:"followers"`
	} `json:"stats"`
	References []struct {
		Type                      string `json:"type"`
		ID                        int    `json:"id"`
		Email                     string `json:"email"`
		Name                      string `json:"name"`
		Community                 bool   `json:"community,omitempty"`
		Permalink                 string `json:"permalink,omitempty"`
		WebURL                    string `json:"web_url"`
		ShowUpgradeBanner         bool   `json:"show_upgrade_banner,omitempty"`
		HeaderBackgroundColor     string `json:"header_background_color,omitempty"`
		HeaderTextColor           string `json:"header_text_color,omitempty"`
		NavigationBackgroundColor string `json:"navigation_background_color,omitempty"`
		NavigationTextColor       string `json:"navigation_text_color,omitempty"`
		Paid                      bool   `json:"paid,omitempty"`
		Moderated                 bool   `json:"moderated,omitempty"`
		IsOrgChartEnabled         bool   `json:"is_org_chart_enabled,omitempty"`
		IsGroupEnabled            bool   `json:"is_group_enabled,omitempty"`
		IsChatEnabled             bool   `json:"is_chat_enabled,omitempty"`
		IsTranslationEnabled      bool   `json:"is_translation_enabled,omitempty"`
		CreatedAt                 string `json:"created_at,omitempty"`
		ProfileFieldsConfig       struct {
			EnableWorkHistory bool `json:"enable_work_history"`
			EnableEducation   bool `json:"enable_education"`
			EnableJobTitle    bool `json:"enable_job_title"`
			EnableWorkPhone   bool `json:"enable_work_phone"`
			EnableMobilePhone bool `json:"enable_mobile_phone"`
			EnableBirthday    bool `json:"enable_birthday"`
			EnableSummary     bool `json:"enable_summary"`
			EnableInterests   bool `json:"enable_interests"`
			EnableExpertise   bool `json:"enable_expertise"`
			EnableLocation    bool `json:"enable_location"`
			EnableIm          bool `json:"enable_im"`
			EnableSkype       bool `json:"enable_skype"`
			EnableWebsites    bool `json:"enable_websites"`
		} `json:"profile_fields_config,omitempty"`
		BrowserDeprecationURL  interface{} `json:"browser_deprecation_url,omitempty"`
		ExternalMessagingState string      `json:"external_messaging_state,omitempty"`
		State                  string      `json:"state,omitempty"`
		FullName               string      `json:"full_name,omitempty"`
		NetworkID              int         `json:"network_id,omitempty"`
		Description            string      `json:"description,omitempty"`
		Privacy                string      `json:"privacy,omitempty"`
		URL                    string      `json:"url,omitempty"`
		MugshotURL             string      `json:"mugshot_url,omitempty"`
		MugshotURLTemplate     string      `json:"mugshot_url_template,omitempty"`
		MugshotID              string      `json:"mugshot_id,omitempty"`
		ShowInDirectory        string      `json:"show_in_directory,omitempty"`
		Members                int         `json:"members,omitempty"`
		Color                  string      `json:"color,omitempty"`
		External               bool        `json:"external,omitempty"`
		JobTitle               string      `json:"job_title,omitempty"`
		ActivatedAt            string      `json:"activated_at,omitempty"`
		AutoActivated          bool        `json:"auto_activated,omitempty"`
		Stats                  struct {
			Following int `json:"following"`
			Followers int `json:"followers"`
			Updates   int `json:"updates"`
		} `json:"stats,omitempty"`
	} `json:"references"`
	Body struct {
		Plain  string `json:"plain"`
		Parsed string `json:"parsed"`
		Rich   string `json:"rich"`
	} `json:"body"`
	Following          bool          `json:"following"`
	ExternalReferences []interface{} `json:"external_references"`
}
