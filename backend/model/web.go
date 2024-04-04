package model

type RenderMessage struct {
	Message
	User
}

type Template struct {
	User *User

	// When on another users page
	Profile  *User
	Messages []RenderMessage
	Followed bool
}
