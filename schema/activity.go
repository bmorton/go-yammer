package schema

var ActivityTypeCreate string = "create"

type Activity struct {
	Actor   *ActivityUser    `json:"actor"`
	Action  string           `json:"action"`
	Object  *OpenGraphObject `json:"object"`
	Message string           `json:"message"`
	Users   []*ActivityUser  `json:"users"`
}

type ActivityUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type OpenGraphObject struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}
