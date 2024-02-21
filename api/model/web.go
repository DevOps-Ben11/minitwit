package model

type FlashMessage struct {
	Message string
}

type RenderMessage struct {
	Message
	User
}

type RenderRequest struct {
	Endpoint string
}

type Template struct {
	User *User
	// When on another users page
	Profile  *User
	Messages []RenderMessage
	Request  RenderRequest
	Followed bool
	Error    *string
	Flashes  []FlashMessage
}
