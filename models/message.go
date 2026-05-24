package models

import "time"

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

type OutgoingMessage struct {
	Type      MessageType `json:"type"`
	StreamID  string      `json:"streamId,omitempty"`
	UserID    string      `json:"userId,omitempty"`
	UserName  string      `json:"userName,omitempty"`
	Content   string      `json:"content,omitempty"`
	Reaction  string      `json:"reaction,omitempty"`
	Count     int         `json:"count,omitempty"`
	Status    string      `json:"status,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

type Stream struct{}