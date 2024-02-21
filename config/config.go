package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	conf       *configImpl
	configOnce sync.Once
)

type Config interface {
	AuthConf() AuthConfig
	FirebaseConf() FirebaseConfig
	DatabaseConf() *Database
	ServerConf() *Server
	ImageConf() ImageConf
}

type configImpl struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Firebase Firebase `mapstructure:"firebase"`
	Jwt      JWT      `mapstructure:"jwt"`
	Image    Image    `mapstructure:"image"`
}

func NewConfig() Config {
	configOnce.Do(func() {
		log.Println("Loading variables is started")
		envMode := os.Getenv("ENV_MODE")
		if envMode == "" {
			envMode = "develop"
		}
		fileConfig := fmt.Sprintf("properties/bwa-startup.%s.yaml", envMode)

		v := viper.New()
		// v.SetConfigType("yaml")
		// v.AddConfigPath(".")
		v.SetConfigFile(fileConfig)

		if err := v.ReadInConfig(); err != nil {
			panic(fmt.Errorf("failed to read config file: %s", err))
		}

		conf = new(configImpl)
		if err := v.Unmarshal(conf); err != nil {
			panic(fmt.Errorf("failed to unmarshal config: %s", err))
		}

		conf.Image.MapImageType = map[string]bool{}
		for _, s := range conf.Image.ImageType {
			conf.Image.MapImageType[s] = true
		}
		log.Println("Loading variables is completed")
	})

	return conf
}

func (c *configImpl) AuthConf() AuthConfig {
	return &c.Jwt
}

func (c *configImpl) FirebaseConf() FirebaseConfig {
	return &c.Firebase
}

func (c *configImpl) DatabaseConf() *Database {
	return &c.Database
}

func (c *configImpl) ServerConf() *Server {
	return &c.Server
}

func (c *configImpl) ImageConf() ImageConf {
	return &c.Image
}
