package twitch

import "net/url"

const (
	// APIV5Header - default v5 api header
	APIV5Header = "application/vnd.twitchtv.v5+json"
)

var (
	// DefaultV5URL - defaul
	DefaultV5URL = &url.URL{
		Scheme: "https",
		Host: "api.twitch.tv",
		Path: "helix",
	}
)
