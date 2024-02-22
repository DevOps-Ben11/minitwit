package api

import (
	"fmt"
)

// configuration

func UrlFor(path string, arg string) string {
	switch path {
	case "add_message":
		return "/add_message"
	case "user_timeline":
		return fmt.Sprintf("/%s", arg)
	case "unfollow_user":
		return fmt.Sprintf("/%s/unfollow", arg)
	case "follow_user":
		return fmt.Sprintf("/%s/follow", arg)
	case "timeline":
		return "/"
	case "public_timeline":
		return "/public"
	case "register":
		return "/register"
	case "login":
		return "/login"
	case "logout":
		return "/logout"
	default:
		return "/"
	}
}
