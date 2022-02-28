package obs

import (
	"brbchat"
	"fmt"
	obsws "github.com/FlowingSPDG/go-obs-websocket"
	"io"
	"log"
	"time"
)

func New(config *brbchat.Config, status brbchat.Status) (brbchat.Broadcaster, error) {
	status.BroadcasterConnected(false)

	obsws.Logger = log.New(io.Discard, "", 0)
	ws := obsws.Client{Host: "localhost", Port: 4444, Password: config.Obs.Password}
	o := obs{
		config:  config,
		ws:      ws,
		channel: make(chan string),
		status:  status,
	}

	for ws.Connect() != nil {
		time.Sleep(5 * time.Second)
	}

	ws.MustAddEventHandler("SwitchScenes", func(event obsws.Event) {
		scenesEvent := event.(obsws.SwitchScenesEvent)
		o.channel <- scenesEvent.SceneName
	})
	status.BroadcasterConnected(true)

	go o.watchdog()
	
	return &o, nil
}

type obs struct {
	config  *brbchat.Config
	ws      obsws.Client
	channel chan string
	status  brbchat.Status
}

func (o *obs) watchdog() {
	for {
		time.Sleep(10 * time.Second)
		if !o.ws.Connected() {
			o.status.BroadcasterConnected(false)
			err := o.ws.Connect()
			if err == nil {
				o.status.BroadcasterConnected(true)
			}
		}
	}
}

func (o *obs) OnSceneSwitch() <-chan string {
	return o.channel
}

func (o *obs) EnableSource(source brbchat.Scene) error {

	if !o.ws.Connected() {
		o.status.BroadcasterConnected(false)
		err := o.ws.Connect()
		if err != nil {
			return fmt.Errorf("failed to connect to obs: %w", err)
		}
		o.status.BroadcasterConnected(true)
	}

	req := obsws.NewGetSceneItemListRequest(o.config.Obs.BrbScene)
	if err := req.Send(o.ws); err != nil {
		return fmt.Errorf("failed to request current source from OBS: %w", err)
	}

	sceneItemList, err := req.Receive()
	if err != nil {
		panic(err)
	}

	for _, item := range sceneItemList.SceneItems {
		sourceName := item["sourceName"].(string)
		sourceId := int(item["itemId"].(float64))

		en := sourceName == source.Name

		r := obsws.NewSetSceneItemRenderRequest(o.config.Obs.BrbScene, sourceName, sourceId, en)
		if err := r.Send(o.ws); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
