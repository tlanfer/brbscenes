package brbchat

type Broadcaster interface {
	EnableSource(scene Scene) error
	OnSceneSwitch() <-chan string
}

type StreamChat interface {
	Listen()
}

type Status interface {
	BroadcasterConnected(bool)
	ChatConnected(bool)
}
