package mng

import (
	"github.com/kamva/mgm/v3"

	"jpc16-telnet-game/type/collection"
)

var ChatMessageCollection *mgm.Collection

func Collection() {
	ChatMessageCollection = mgm.Coll(new(collection.ChatMessage))
}
