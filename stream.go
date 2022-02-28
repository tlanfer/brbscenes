package brbchat

type Broadcaster interface {
	SetupSources(map[string]bool) error
	OnSceneSwitch() <-chan string
}

type StreamChat interface {
	Listen()
}

type Status interface {
	BroadcasterConnected(bool)
	ChatConnected(bool)
}
