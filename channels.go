package twitch

// ChannelType - describes a channel on twitch
type ChannelType []struct {
	BroadcasterID string `json:"broadcaster_id"`
	BroadcasterName string `json:"broadcaster_name"`
	BroadcasterLanguage string `json:"broadcaster_language"`
	GameID string `json:"game_id"`
	GameName string `json:"game_name"`
	Title string `json:"title"`
}

// Data - the root of the json
type Data struct {
	Data ChannelType `json:"data"`
}

//
// Implementation and their respective request/response types
//

// GetChannelInputType - input type for the GetChannel function
type GetChannelInputType struct {
	BroadcasterID string `url:"broadcaster_id"`
}

// GetChannelOutputType - returned type containing the channel
type GetChannelOutputType Data

// GetChannel - returns the specified channel
func (session *Session) GetChannel(getChannelInputType *GetChannelInputType) (*GetChannelOutputType, error) {

	var out GetChannelOutputType
	err := session.request("GET", "/channels", getChannelInputType, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetChannelTeamsInputType - input type for the GetChannelTeams function
type GetChannelTeamsInputType struct {
	Channel string
}

// GetChannelTeamsOutputType - returned type container an array of teams
type GetChannelTeamsOutputType struct {
	Teams []TeamType        `json:"teams"`
	Links map[string]string `json:"_links"`
}

// GetChannelTeams - - returns an array of the teams the specified channel belongs to
func (session *Session) GetChannelTeams(getChannelTeamsInputType *GetChannelTeamsInputType) (*GetChannelTeamsOutputType, error) {
	var out GetChannelTeamsOutputType
	err := session.request("GET", "/channels/"+getChannelTeamsInputType.Channel+"/teams", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
