package twitchchat

import (
	"brbchat"
	"github.com/gempir/go-twitch-irc/v3"
	"log"
	"strings"
	"time"
)

func New(config *brbchat.Config, broadcaster brbchat.Broadcaster, status brbchat.Status) brbchat.StreamChat {
	status.ChatConnected(false)
	return &chat{
		broadcaster: broadcaster,
		config:      config,
		status:      status,
		enabled:     true,
	}
}

type chat struct {
	broadcaster brbchat.Broadcaster
	config      *brbchat.Config

	enabled bool
	timeout time.Time
	status  brbchat.Status
}

func (c *chat) Listen() {

	go func() {
		sceneSwitch := c.broadcaster.OnSceneSwitch()
		for scene := range sceneSwitch {
			c.enabled = scene == c.config.Obs.BrbScene
		}
	}()

	tc := twitch.NewAnonymousClient()
	tc.Join(c.config.Channel)

	tc.OnPrivateMessage(func(message twitch.PrivateMessage) {

		if c.enabled && time.Now().After(c.timeout) {

			for _, scene := range c.config.Scenes {
				if strings.Contains(message.Message, scene.Keyword) {

					err := c.broadcaster.EnableSource(scene)
					if err != nil {
						log.Printf("failed to enable scene %v. Is obs running?", scene.Name)
						return
					}
					minimum := c.config.Minimum
					if scene.Minimum > 0 {
						minimum = scene.Minimum
					}
					c.timeout = time.Now().Add(minimum)
				}
			}

		}
	})

	c.status.ChatConnected(true)
	err := tc.Connect()
	if err != nil {
		c.status.ChatConnected(false)
	}
}
