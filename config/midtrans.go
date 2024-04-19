package config

type MidtransConf interface {
	GetServerKey() string
	GetClientKey() string
}

type Midtrans struct {
	Serverkey string `mapstructure:"server-key"`
	ClientKey string `mapstructure:"client-key"`
}

func (m *Midtrans) GetServerKey() string {
	return m.Serverkey
}

func (m *Midtrans) GetClientKey() string {
	return m.ClientKey
}
