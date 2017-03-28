package schema

type Group struct {
	Type               string `json:"type"`
	ID                 int    `json:"id"`
	Email              string `json:"email"`
	FullName           string `json:"full_name"`
	NetworkID          int    `json:"network_id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Privacy            string `json:"privacy"`
	URL                string `json:"url"`
	WebURL             string `json:"web_url"`
	MugshotURL         string `json:"mugshot_url"`
	MugshotURLTemplate string `json:"mugshot_url_template"`
	MugshotID          string `json:"mugshot_id"`
	ShowInDirectory    string `json:"show_in_directory"`
	CreatedAt          string `json:"created_at"`
	Color              string `json:"color"`
	External           bool   `json:"external"`
	Moderated          bool   `json:"moderated"`
	CreatorType        string `json:"creator_type"`
	CreatorID          int    `json:"creator_id"`
	State              string `json:"state"`
	Stats              struct {
		Members       int    `json:"members"`
		Updates       int    `json:"updates"`
		LastMessageID int    `json:"last_message_id"`
		LastMessageAt string `json:"last_message_at"`
	} `json:"stats"`
	Member        bool   `json:"member"`
	Pending       bool   `json:"pending"`
	Admin         bool   `json:"admin"`
	HasAdmin      bool   `json:"has_admin"`
	CanAddMembers bool   `json:"can_add_members"`
	CanInvite     bool   `json:"can_invite"`
	NetworkName   string `json:"network_name"`
	FeaturedUsers []struct {
		Type               string `json:"type"`
		ID                 int    `json:"id"`
		Name               string `json:"name"`
		State              string `json:"state"`
		FullName           string `json:"full_name"`
		JobTitle           string `json:"job_title"`
		NetworkID          int    `json:"network_id"`
		MugshotURL         string `json:"mugshot_url"`
		MugshotURLTemplate string `json:"mugshot_url_template"`
		URL                string `json:"url"`
		WebURL             string `json:"web_url"`
		ActivatedAt        string `json:"activated_at"`
		AutoActivated      bool   `json:"auto_activated"`
		Stats              struct {
			Following int `json:"following"`
			Followers int `json:"followers"`
			Updates   int `json:"updates"`
		} `json:"stats"`
		Email string `json:"email"`
		Admin bool   `json:"admin"`
	} `json:"featured_users"`
	RelatedGroups []struct {
		Type               string `json:"type"`
		ID                 int    `json:"id"`
		Email              string `json:"email"`
		FullName           string `json:"full_name"`
		NetworkID          int    `json:"network_id"`
		Name               string `json:"name"`
		Description        string `json:"description"`
		Privacy            string `json:"privacy"`
		URL                string `json:"url"`
		WebURL             string `json:"web_url"`
		MugshotURL         string `json:"mugshot_url"`
		MugshotURLTemplate string `json:"mugshot_url_template"`
		MugshotID          string `json:"mugshot_id"`
		ShowInDirectory    string `json:"show_in_directory"`
		CreatedAt          string `json:"created_at"`
		Members            int    `json:"members"`
		Color              string `json:"color"`
		External           bool   `json:"external"`
		Moderated          bool   `json:"moderated"`
		RelationshipID     int    `json:"relationship_id"`
	} `json:"related_groups"`
	RelatedContents []RelatedContent `json:"related_contents"`
	FeedKey         string           `json:"feed_key"`
	Subscriptions   struct {
		Email string `json:"email"`
	} `json:"subscriptions"`
	GroupInfo              string `json:"group_info"`
	PendingRequestsCount   int    `json:"pending_requests_count"`
	InvitedRequestsCount   int    `json:"invited_requests_count"`
	SuggestedRequestsCount int    `json:"suggested_requests_count"`
}

type RelatedContent struct {
	Type                 string      `json:"type"`
	ID                   int64       `json:"id"`
	NetworkID            int         `json:"network_id"`
	URL                  string      `json:"url"`
	WebURL               string      `json:"web_url"`
	InternalURL          string      `json:"internal_url"`
	HostURL              string      `json:"host_url"`
	FullName             string      `json:"full_name"`
	OfficialName         string      `json:"official_name"`
	Name                 string      `json:"name"`
	ObjectType           string      `json:"object_type"`
	Image                string      `json:"image"`
	VideoURL             string      `json:"video_url"`
	Description          string      `json:"description"`
	SiteName             string      `json:"site_name"`
	ExtendedInfo         interface{} `json:"extended_info"`
	CreatorID            int         `json:"creator_id"`
	CreatorType          string      `json:"creator_type"`
	CreatedAt            string      `json:"created_at"`
	MediaURL             string      `json:"media_url"`
	MediaType            string      `json:"media_type"`
	MediaWidth           interface{} `json:"media_width"`
	MediaHeight          interface{} `json:"media_height"`
	Private              bool        `json:"private"`
	IsFirstParty         bool        `json:"is_first_party"`
	TitleAndDescEditable bool        `json:"title_and_desc_editable"`
	RelationshipID       int         `json:"relationship_id"`
}
