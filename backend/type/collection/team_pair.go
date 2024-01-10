package collection

import (
	mh "jpc16-telnet-game/common/mng/helper"
)

type ChatMessage struct {
	mh.ModelBase `bson:"-,inline"`
	SourceIp     *string `bson:"sourceIp,omitempty"`
}
