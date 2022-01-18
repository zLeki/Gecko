package structs

import (
	"time"
)

type ChannelData []struct {
	ID              string        `json:"id"`
	Type            int           `json:"type"`
	Content         string        `json:"content"`
	ChannelID       string        `json:"channel_id"`
	Attachments     []interface{} `json:"attachments"`
	Embeds          []interface{} `json:"embeds"`
	Mentions        []interface{} `json:"mentions"`
	MentionRoles    []interface{} `json:"mention_roles"`
	Pinned          bool          `json:"pinned"`
	MentionEveryone bool          `json:"mention_everyone"`
	Tts             bool          `json:"tts"`
	Timestamp       time.Time     `json:"timestamp"`
	EditedTimestamp interface{}   `json:"edited_timestamp"`
	Flags           int           `json:"flags"`
	Components      []interface{} `json:"components"`
	Author          struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		Avatar        string `json:"avatar"`
		Discriminator string `json:"discriminator"`
		PublicFlags   int    `json:"public_flags"`
		Bot           bool   `json:"bot"`
	} `json:"author,omitempty"`

	MessageReference struct {
		ChannelID string `json:"channel_id"`
		GuildID   string `json:"guild_id"`
		MessageID string `json:"message_id"`
	} `json:"message_reference,omitempty"`
	ReferencedMessage struct {
		ID        string `json:"id"`
		Type      int    `json:"type"`
		Content   string `json:"content"`
		ChannelID string `json:"channel_id"`
		Author    struct {
			ID            string `json:"id"`
			Username      string `json:"username"`
			Avatar        string `json:"avatar"`
			Discriminator string `json:"discriminator"`
			PublicFlags   int    `json:"public_flags"`
		} `json:"author"`
		Attachments     []interface{} `json:"attachments"`
		Embeds          []interface{} `json:"embeds"`
		Mentions        []interface{} `json:"mentions"`
		MentionRoles    []interface{} `json:"mention_roles"`
		Pinned          bool          `json:"pinned"`
		MentionEveryone bool          `json:"mention_everyone"`
		Tts             bool          `json:"tts"`
		Timestamp       time.Time     `json:"timestamp"`
		EditedTimestamp time.Time     `json:"edited_timestamp"`
		Flags           int           `json:"flags"`
		Components      []interface{} `json:"components"`
	} `json:"referenced_message,omitempty"`
}
type GuidldChannelData []struct {
	ID                   string      `json:"id"`
	LastMessageID        string      `json:"last_message_id,omitempty"`
	LastPinTimestamp     time.Time   `json:"last_pin_timestamp,omitempty"`
	Type                 int         `json:"type"`
	Name                 string      `json:"name"`
	Position             int         `json:"position"`
	ParentID             string      `json:"parent_id"`
	Topic                interface{} `json:"topic,omitempty"`
	GuildID              string      `json:"guild_id"`
	PermissionOverwrites []struct {
		ID       string `json:"id"`
		Type     string `json:"type"`
		Allow    int    `json:"allow"`
		Deny     int    `json:"deny"`
		AllowNew string `json:"allow_new"`
		DenyNew  string `json:"deny_new"`
	} `json:"permission_overwrites"`
	Nsfw             bool        `json:"nsfw"`
	RateLimitPerUser int         `json:"rate_limit_per_user,omitempty"`
	Banner           interface{} `json:"banner,omitempty"`
	Bitrate          int         `json:"bitrate,omitempty"`
	UserLimit        int         `json:"user_limit,omitempty"`
	RtcRegion        interface{} `json:"rtc_region,omitempty"`
}
