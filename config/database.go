package config

import "fmt"

type Database struct {
	Hostname                      string `mapstructure:"hostname"`
	Port                          string `mapstructure:"port"`
	DatabaseName                  string `mapstructure:"dbname"`
	Username                      string `mapstructure:"user"`
	Password                      string `mapstructure:"password"`
	SslMode                       string `mapstructure:"ssl-mode"`
	TimeZone                      string `mapstructure:"timeZone"`
	MaxIdleConnectionsInSecond    int    `mapstructure:"max-idle-connections-in-second"`
	MaxOpenConnectionsInSecond    int    `mapstructure:"max-open-connections-in-second"`
	ConnectionMaxLifetimeInSecond int64  `mapstructure:"connection-max-life-time-in-second"`
}

func (d *Database) DNS() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		d.Hostname,
		d.Username,
		d.Password,
		d.DatabaseName,
		d.Port,
		d.SslMode,
		d.TimeZone,
	)
}
