package systray

import (
	"brbchat"
	_ "embed"
	st "github.com/getlantern/systray"
	"os"
	"time"
)

//go:embed alasuwu_off_off.ico
var alasuwuOffOff []byte

//go:embed alasuwu_off_on.ico
var alasuwuOffOn []byte

//go:embed alasuwu_on_off.ico
var alasuwuOnOff []byte

//go:embed alasuwu_on_on.ico
var alasuwuOnOn []byte

func New() brbchat.Status {
	s := &status{false, false}
	go st.Run(s.onReady, s.onExit)
	time.Sleep(100 * time.Millisecond)
	return s
}

type status struct {
	broadcasterConnected bool
	chatConnected        bool
}

func (s *status) BroadcasterConnected(b bool) {
	s.broadcasterConnected = b
	s.icon()
}

func (s *status) ChatConnected(b bool) {
	s.chatConnected = b
	s.icon()
}

func (s *status) icon() {
	if s.broadcasterConnected && s.chatConnected {
		st.SetTemplateIcon(alasuwuOnOn, alasuwuOnOn)
	}
	if !s.broadcasterConnected && s.chatConnected {
		st.SetTemplateIcon(alasuwuOnOff, alasuwuOnOff)
	}
	if s.broadcasterConnected && !s.chatConnected {
		st.SetTemplateIcon(alasuwuOffOn, alasuwuOffOn)
	}
	if !s.broadcasterConnected && !s.chatConnected {
		st.SetTemplateIcon(alasuwuOffOff, alasuwuOffOff)
	}
}

func (s *status) onReady() {
	st.SetTemplateIcon(alasuwuOffOff, alasuwuOffOff)
	st.SetTitle("Brb Scenes")
	st.SetTooltip("Brb Scenes")

	item := st.AddMenuItem("Quit", "Quit")

	go func() {
		for {
			select {
			case <-item.ClickedCh:
				st.Quit()
			}
		}
	}()
}

func (s *status) onExit() {
	os.Exit(0)
}
