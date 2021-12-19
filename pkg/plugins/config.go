package plugins

import "github.com/hashicorp/go-plugin"

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "onebot-plus",
	MagicCookieValue: "onebot-plus",
}
