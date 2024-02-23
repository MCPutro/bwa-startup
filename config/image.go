package config

type ImageConf interface {
	IsSupport(t string) bool
	MaxAvatarSize() int64
	SupportType() map[string]bool
}

type Image struct {
	ImageType    []string `mapstructure:"support-type"`
	MaxAvatar    int64    `mapstructure:"max-avatar-size-in-mb"`
	MapImageType map[string]bool
}

func (c *Image) MaxAvatarSize() int64 {
	return c.MaxAvatar << 20
}

func (c *Image) IsSupport(t string) bool {
	return c.MapImageType[t]
}

func (c *Image) SupportType() map[string]bool {
	return c.MapImageType
}
