package brbchat

import "time"

type ConfigLoader interface {
	Load() (*Config, error)
}

type Config struct {
	Channel  string        `yaml:"channel"`
	Cooldown time.Duration `yaml:"cooldown"`
	Obs      Obs           `yaml:"obs"`
	Sources  []Source      `yaml:"sources"`
}

type Obs struct {
	BrbScene string `yaml:"brb_scene"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

type Source struct {
	Name     string        `yaml:"name"`
	Keyword  string        `yaml:"keyword"`
	Cooldown time.Duration `yaml:"cooldown,omitempty"`
}
