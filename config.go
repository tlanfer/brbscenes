package brbchat

import "time"

type ConfigLoader interface {
	Load() (*Config, error)
}

type Config struct {
	Channel string        `yaml:"channel"`
	Minimum time.Duration `yaml:"min_duration"`
	Obs     Obs           `yaml:"obs"`
	Scenes  []Scene       `yaml:"scenes"`
}

type Obs struct {
	BrbScene string `yaml:"brb_scene"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

type Scene struct {
	Name    string        `yaml:"name"`
	Keyword string        `yaml:"keyword"`
	Minimum time.Duration `yaml:"min_duration,omitempty"`
}
