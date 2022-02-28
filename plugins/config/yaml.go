package yamlconfig

import (
	"brbchat"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

func New(filename string) brbchat.ConfigLoader {
	return &loader{
		filename: filename,
	}
}

type loader struct {
	filename string
}

func (l *loader) Load() (*brbchat.Config, error) {
	dto := &brbchat.Config{}

	file, err := os.OpenFile(l.filename, os.O_RDONLY, os.ModePerm)
	if err != nil {

		if os.IsNotExist(err) {

			file, err := os.OpenFile(l.filename, os.O_CREATE, os.ModePerm)
			if err != nil {
				return nil, fmt.Errorf("failed to create example config: %w", err)
			}
			defer file.Close()

			dto.Channel = "alasdair"
			dto.Cooldown = 30 * time.Second
			dto.Obs = brbchat.Obs{
				BrbScene: "brb",
				Password: "roflmao",
				Port:     4444,
			}

			dto.Sources = []brbchat.Source{
				{Name: "tetris", Keyword: "!tetris"},
				{Name: "trover", Keyword: "!trover"},
				{Name: "worms", Keyword: "!worms"},
				{Name: "twiggie", Keyword: "!twiggie"},
			}

			err = yaml.NewEncoder(file).Encode(dto)
			if err != nil {
				return nil, fmt.Errorf("failed to create example config: %w", err)
			}

			return nil, fmt.Errorf("created example config %q. please fill in the config and start the app again", l.filename)
		}

		return nil, err
	}

	defer file.Close()

	err = yaml.NewDecoder(file).Decode(dto)
	if err != nil {
		return nil, err
	}

	if dto.Cooldown == 0 {
		return nil, fmt.Errorf("missing minimum duration")
	}
	if dto.Channel == "" {
		return nil, fmt.Errorf("missing channel")
	}
	if dto.Obs.BrbScene == "" {
		return nil, fmt.Errorf("missing obs.brb_scene")
	}
	if dto.Obs.Port == 0 {
		return nil, fmt.Errorf("missing obs.port")
	}
	if len(dto.Sources) == 0 {
		return nil, fmt.Errorf("missing scenes")
	}

	return dto, nil
}
