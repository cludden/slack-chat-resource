package main

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

type message struct {
	ThreadTS *string      `json:"thread_ts,omitempty"`
	Blocks   slack.Blocks `json:"blocks"`
}

type slackBlockType struct {
	Type slack.MessageBlockType `json:"type"`
}

func main() {
	input := `
	{
		"thread_ts": null,
		"blocks": [
			{
				"type": "context",
				"block_id": "123.34",
				"elements": [
					{
						"type": "mrkdwn",
						"text": "*Hello* **there**"
					}
				]
			}
		]
	}
	`

	var msg message
	if err := json.Unmarshal([]byte(input), &msg); err != nil {
		fmt.Printf("unmarshal error: %v", err)
		os.Exit(1)
	}
}
