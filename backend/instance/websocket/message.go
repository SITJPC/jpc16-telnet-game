package websocket

type InboundEvent string
type OutboundEvent string

const (
	ConnectionSwitchEvent OutboundEvent = "general/switch"
	EchoEvent             OutboundEvent = "general/echo"
	ChatMessageEvent      OutboundEvent = "chat/message"
	ChatHistoryEvent      OutboundEvent = "chat/history"
)

type OutboundMessage struct {
	Event   OutboundEvent `json:"event"`
	Payload any           `json:"payload"`
}
