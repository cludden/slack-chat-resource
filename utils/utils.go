package utils

import (
	"encoding/json"
	"regexp"

	"github.com/slack-go/slack"
)

// Regexp type definition
type Regexp struct{ regexp.Regexp }

// UnmarshalJSON custom unmarshaller
func (r *Regexp) UnmarshalJSON(payload []byte) error {
	var pattern string
	err := json.Unmarshal(payload, &pattern)
	if err != nil {
		return err
	}

	regexp, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	*r = Regexp{*regexp}
	return nil
}

// MessageFilter type definition
type MessageFilter struct {
	AuthorID    string  `json:"author"`
	TextPattern *Regexp `json:"text_pattern"`
}

// Source type definition
type Source struct {
	Token           string         `json:"token"`
	ChannelID       string         `json:"channel_id"`
	CheckMostRecent int            `json:"check_most_recent"`
	Filter          *MessageFilter `json:"matching"`
	ReplyFilter     *MessageFilter `json:"not_replied_by"`
}

// Version type definition
type Version map[string]string

// Metadata type definition
type Metadata []MetadataField

// MetadataField type definition
type MetadataField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// InRequest type definition
type InRequest struct {
	Source  Source   `json:"source"`
	Version Version  `json:"version"`
	Params  InParams `json:"params"`
}

// InResponse type definition
type InResponse struct {
	Version  Version  `json:"version"`
	Metadata Metadata `json:"metadata"`
}

// InParams type definition
type InParams struct {
	TextPattern *Regexp `json:"text_pattern"`
}

// OutRequest type definition
type OutRequest struct {
	Source Source    `json:"source"`
	Params OutParams `json:"params"`
}

// OutParams type definition
type OutParams struct {
	Message     *OutMessage `json:"message"`
	MessageFile string      `json:"message_file"`
}

// OutMessage type definition
type OutMessage struct {
	Attachments []slack.Attachment `json:"attachments"`
	Blocks      slack.Blocks       `json:"blocks"`
	Text        string             `json:"text"`
	slack.PostMessageParameters
}

// OutResponse type definition
type OutResponse struct {
	Version  Version  `json:"version"`
	Metadata Metadata `json:"metadata"`
}

// CheckRequest type definition
type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

// CheckResponse type definition
type CheckResponse []Version

// SlackRequest type definition
type SlackRequest struct {
	Contents string
}
