package config

import (
	"errors"
	"flag"
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
	MidtransConf() MidtransConf
}

type configImpl struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Firebase Firebase `mapstructure:"firebase"`
	Jwt      JWT      `mapstructure:"jwt"`
	Image    Image    `mapstructure:"image"`
	Midtrans Midtrans `mapstructure:"midtrans"`
}

func NewConfig() (Config, error) {

	var configPath string

	// Set up a CLI flag called "-config" to allow users
	flag.StringVar(&configPath, "config", "", "path to config file")

	// Actually parse the flags
	flag.Parse()

	if err := validateConfigPath(configPath); err != nil {
		return nil, err
	} else {
		configOnce.Do(func() {
			log.Println("Loading variables is started")

			v := viper.New()
			// v.SetConfigType("yaml")
			// v.AddConfigPath(".")
			v.SetConfigFile(configPath)

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
	}

	return conf, nil
}

func validateConfigPath(path string) error {
	if path == "" {
		return errors.New("config file was not found")
	}
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
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

func (c *configImpl) MidtransConf() MidtransConf {
	return &c.Midtrans
}
