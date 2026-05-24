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

type Stream struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	TurfID      string    `json:"turfId"`
	Status      string    `json:"status"` // SCHEDULED, LIVE, ENDED
	HLSUrl      string    `json:"hlsUrl"`
	ViewerCount int       `json:"viewerCount"`
	StartedAt   time.Time `json:"startedAt,omitempty"`
}

// Authenticated user info
type AuthUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
