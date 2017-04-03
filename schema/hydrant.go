package schema

type HydrantUnviewedThreads struct {
	Threads       []HydrantThread `json:"threads"`
	HaveOlderMsgs bool            `json:"have_older_msgs"`
}

type HydrantURLTemplates struct {
	UserWeb             string `json:"user_web"`
	AttachmentWeb       string `json:"attachment_web"`
	AttachmentStream    string `json:"attachment_stream"`
	Attachment          string `json:"attachment"`
	FileLargeIcon       string `json:"file_large_icon"`
	PageWeb             string `json:"page_web"`
	PagePreview         string `json:"page_preview"`
	Group               string `json:"group"`
	Realtime            string `json:"realtime"`
	AttachmentDownload  string `json:"attachment_download"`
	GroupMugshot        string `json:"group_mugshot"`
	MessageWeb          string `json:"message_web"`
	Thread              string `json:"thread"`
	Message             string `json:"message"`
	AttachmentPreview   string `json:"attachment_preview"`
	AttachmentScaled    string `json:"attachment_scaled"`
	TopicWeb            string `json:"topic_web"`
	GroupWeb            string `json:"group_web"`
	Topic               string `json:"topic"`
	AttachmentEdit      string `json:"attachment_edit"`
	UserMugshot         string `json:"user_mugshot"`
	Page                string `json:"page"`
	ThreadWeb           string `json:"thread_web"`
	User                string `json:"user"`
	AttachmentThumbnail string `json:"attachment_thumbnail"`
}

type HydrantFeed struct {
	ID                string                  `json:"id"`
	ChannelID         string                  `json:"channel_id"`
	NetworkID         int                     `json:"network_id"`
	HasOlderThreads   bool                    `json:"has_older_threads"`
	Threads           []*HydrantThread        `json:"threads"`
	UnseenThreadCount int                     `json:"unseen_thread_count"`
	LastSeenMessageID int                     `json:"last_seen_message_id"`
	UnviewedThreads   *HydrantUnviewedThreads `json:"unviewed_threads"`
	ReferenceData     *HydrantReferenceData   `json:"reference_data"`
	URLTemplates      *HydrantURLTemplates    `json:"url_templates"`
}

type HydrantThreadFeed struct {
	ID            []int                 `json:"id"`
	Threads       []*HydrantThread      `json:"threads"`
	ReferenceData *HydrantReferenceData `json:"reference_data"`
	URLTemplates  *HydrantURLTemplates  `json:"url_templates"`
}

type HydrantThread struct {
	ID                    int                 `json:"id"`
	NetworkID             int                 `json:"network_id"`
	InPrivateConversation bool                `json:"in_private_conversation"`
	HasAttachments        bool                `json:"has_attachments"`
	GroupIds              []int               `json:"group_ids"`
	ThreadStarterID       int                 `json:"thread_starter_id"`
	MessagesCount         int                 `json:"messages_count"`
	State                 *HydrantThreadState `json:"state"`
	Messages              []*HydrantMessage   `json:"messages"`
	ReadOnly              bool                `json:"read_only"`
	FirstReplyID          int                 `json:"first_reply_id,omitempty"`
	Participants          []int               `json:"participants,omitempty"`
}

type HydrantThreadState struct {
	LastReadMessage    int `json:"last_read_message"`
	UnseenMessageCount int `json:"unseen_message_count"`
}

type HydrantMessage struct {
	ID                    int                `json:"id"`
	ThreadID              int                `json:"thread_id"`
	NetworkID             int                `json:"network_id"`
	SenderID              int                `json:"sender_id"`
	MessageType           string             `json:"message_type"`
	GroupID               int                `json:"group_id"`
	PageIds               []int              `json:"page_ids"`
	CreatedAt             int64              `json:"created_at"`
	Body                  string             `json:"body"`
	Language              string             `json:"language"`
	SenderType            string             `json:"sender_type"`
	ClientTypeID          int                `json:"client_type_id"`
	LikeIDs               []int              `json:"like_ids"`
	References            *HydrantReferences `json:"references"`
	InPrivateConversation bool               `json:"in_private_conversation"`
	RepliedToID           int                `json:"replied_to_id"`
	AdditionalData        string             `json:"additional_data"`
}

type HydrantReferences struct {
	User *HydrantUserReferences `json:"user"`
}

type HydrantUserReferences struct {
	CC       []int `json:"cc"`
	RE       []int `json:"re"`
	InThread []int `json:"in_thread"`
}

type HydrantReferenceData struct {
	Users              map[string]*HydrantUser              `json:"users"`
	Networks           map[string]*HydrantNetwork           `json:"networks"`
	Groups             map[string]*HydrantGroup             `json:"groups"`
	ClientApplications map[string]*HydrantClientApplication `json:"client_applications"`
}

type HydrantClientApplication struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type HydrantGroup struct {
	ID        int    `json:"id"`
	NetworkID int    `json:"network_id"`
	Name      string `json:"name"`
	Permalink string `json:"permalink"`
	Private   bool   `json:"private"`
	Moderated bool   `json:"moderated"`
	MugshotID string `json:"mugshot_id"`
	Color     string `json:"color"`
}

type HydrantNetwork struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type HydrantUser struct {
	ID          int    `json:"id"`
	NetworkID   int    `json:"network_id"`
	FullName    string `json:"full_name"`
	Permalink   string `json:"permalink"`
	Email       string `json:"email"`
	JobTitle    string `json:"job_title"`
	MugshotID   string `json:"mugshot_id"`
	ActivatedAt int64  `json:"activated_at"`
	State       string `json:"state"`
	IsFollowing bool   `json:"is_following"`
}
