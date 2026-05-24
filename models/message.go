package models

// message types
type MessageType string

const (
	TypeChat         MessageType = "CHAT"
	TypeReaction     MessageType = "REACTION"
	TypeViewerCount  MessageType = "VIEWER_COUNT"
	TypeStreamStatus MessageType = "STREAM_STATUS"
	TypeError        MessageType = "ERROR"
	TypePing         MessageType = "PING"
	TypePong         MessageType = "PONG"
)

// clinet message format
type IncomingMessage struct {
	Type     MessageType `json:"type"`
	Content  string      `json:"content,omitempty"`
	Reaction string      `json:"reaction,omitempty"`
}
