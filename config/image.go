package config

type ImageConf interface {
	SupportType(t string) bool
	MaxAvatarSize() int64
}

type Image struct {
	ImageType    []string `mapstructure:"support-type"`
	MaxAvatar    int64    `mapstructure:"max-avatar-size-in-mb"`
	MapImageType map[string]bool
}

func (c *Image) MaxAvatarSize() int64 {
	return c.MaxAvatar << 20
}

func (c *Image) SupportType(t string) bool {
	return c.MapImageType[t]
}
