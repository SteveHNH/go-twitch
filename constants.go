package twitch

import "net/url"


var (
	// DefaultURL - default
	DefaultURL = &url.URL{
		Scheme: "https",
		Host: "api.twitch.tv",
		Path: "helix",
	}
)
