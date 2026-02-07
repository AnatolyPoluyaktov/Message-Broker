package config

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Queues map[string]QueueConfig `mapstructure:"queues"`
	Server ServerConfig           `mapstructure:"server"`
}

type QueueConfig struct {
	Name                     string `mapstructure:"name"`
	MaxItems                 int64  `mapstructure:"max_items"`
	MaxSubscribers           int64  `mapstructure:"max_subscribers"`
	MaxMessagesPerSubscriber int64  `mapstructure:"max_messages_per_subscriber"`
}

type ServerConfig struct {
	Address string        `mapstructure:"address"`
	Timeout time.Duration `mapstructure:"timeout"`
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "read config")
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "Unmarshal config")
	}
	return &cfg, nil
}
