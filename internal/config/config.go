package config

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	StageLocal Stage = "local"
	StageDEV   Stage = "dev"
	StageSIT   Stage = "sit"
	StageNFT   Stage = "nft"
	StageUAT   Stage = "uat"
	StageProd  Stage = "prod"
)

type (
	Stage = string

	Config struct {
		vn         *viper.Viper
		mux        *sync.Mutex
		configPath string
		stage      Stage

		Merchant Merchant `mapstructure:"merchant"`
		Database Database `mapstructure:"database"`
	}

	Merchant struct {
		Enable bool `mapstructure:"enable"`
	}

	Database struct {
		Type    string  `mapstructure:"type" default:"mongodb"`
		MongoDB MongoDB `mapstructure:"mongo_db"`
	}

	MongoDB struct {
		Addresses []string      `mapstructure:"addresses"`
		Username  string        `mapstructure:"username"`
		Password  string        `mapstructure:"password"`
		Database  string        `mapstructure:"database"`
		Timeout   time.Duration `mapstructure:"timeout"`
	}
)

func (c *Config) Init(stage, cfgPath string) error {
	c.mux = &sync.Mutex{}
	c.stage = stage
	c.configPath = cfgPath
	name := fmt.Sprintf("config.%s", c.stage)

	vn := viper.New()
	vn.AddConfigPath(c.configPath)
	vn.SetConfigName(name)

	if err := vn.ReadInConfig(); err != nil {
		return err
	}
	c.vn = vn

	if err := c.binding(); err != nil {
		return err
	}

	vn.WatchConfig()
	vn.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed:", e.Name)
		if err := c.binding(); err != nil {
			return
		}
	})

	return nil
}

func (c *Config) binding() error {
	c.mux.Lock()
	defer c.mux.Unlock()

	if err := c.vn.Unmarshal(&c); err != nil {
		return err
	}
	return nil
}

func ParseStage(s string) Stage {
	switch s {
	case "local", "localhost", "l":
		return StageLocal
	case "dev", "develop", "development", "d":
		return StageDEV
	case "sit", "staging", "s":
		return StageSIT
	case "nft":
		return StageNFT
	case "uat":
		return StageUAT
	case "prod", "production", "p":
		return StageProd
	}
	return s
}
