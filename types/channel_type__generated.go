package types

import (
	"bytes"
	"encoding"
	"errors"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
)

var InvalidChannelType = errors.New("invalid ChannelType")

func init() {
	github_com_johnnyeven_libtools_courier_enumeration.RegisterEnums("ChannelType", map[string]string{
		"GITHUB": "Github",
	})
}

func ParseChannelTypeFromString(s string) (ChannelType, error) {
	switch s {
	case "":
		return CHANNEL_TYPE_UNKNOWN, nil
	case "GITHUB":
		return CHANNEL_TYPE__GITHUB, nil
	}
	return CHANNEL_TYPE_UNKNOWN, InvalidChannelType
}

func ParseChannelTypeFromLabelString(s string) (ChannelType, error) {
	switch s {
	case "":
		return CHANNEL_TYPE_UNKNOWN, nil
	case "Github":
		return CHANNEL_TYPE__GITHUB, nil
	}
	return CHANNEL_TYPE_UNKNOWN, InvalidChannelType
}

func (ChannelType) EnumType() string {
	return "ChannelType"
}

func (ChannelType) Enums() map[int][]string {
	return map[int][]string{
		int(CHANNEL_TYPE__GITHUB): {"GITHUB", "Github"},
	}
}
func (v ChannelType) String() string {
	switch v {
	case CHANNEL_TYPE_UNKNOWN:
		return ""
	case CHANNEL_TYPE__GITHUB:
		return "GITHUB"
	}
	return "UNKNOWN"
}

func (v ChannelType) Label() string {
	switch v {
	case CHANNEL_TYPE_UNKNOWN:
		return ""
	case CHANNEL_TYPE__GITHUB:
		return "Github"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*ChannelType)(nil)

func (v ChannelType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidChannelType
	}
	return []byte(str), nil
}

func (v *ChannelType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseChannelTypeFromString(string(bytes.ToUpper(data)))
	return
}
