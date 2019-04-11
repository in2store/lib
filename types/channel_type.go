package types

//go:generate libtools gen enum ChannelType
// swagger:enum
type ChannelType uint8

// GIT通道类型
const (
	CHANNEL_TYPE_UNKNOWN ChannelType = iota
	CHANNEL_TYPE__GITHUB             // Github
)
