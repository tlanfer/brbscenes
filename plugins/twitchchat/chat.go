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

	lastScene string
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

			keywordFound := false
			sources := map[string]bool{}

			for _, scene := range c.config.Sources {

				containsKeyword := strings.Contains(message.Message, scene.Keyword)
				wasLastScene := scene.Name == c.lastScene

				if !keywordFound && containsKeyword && !wasLastScene {
					sources[scene.Name] = true

					minimum := c.config.Cooldown
					if scene.Cooldown > 0 {
						minimum = scene.Cooldown
					}
					c.lastScene = scene.Name
					c.timeout = time.Now().Add(minimum)
					keywordFound = true
				} else {
					sources[scene.Name] = false
				}
			}

			if keywordFound {
				err := c.broadcaster.SetupSources(sources)
				if err != nil {
					log.Printf("failed to change scenes. Is obs running? %v", err.Error())
					return
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
