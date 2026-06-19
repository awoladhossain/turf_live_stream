// Package models defines shared data structures for WebSocket messages,
// stream metadata, and authenticated user information.
package models

import "time"

// MessageType represents the type of WebSocket message exchanged between client and server.
type MessageType string

// Supported WebSocket message types for the live streaming chat and control protocol.
const (
	TypeChat         MessageType = "CHAT"
	TypeReaction     MessageType = "REACTION"
	TypeViewerCount  MessageType = "VIEWER_COUNT"
	TypeStreamStatus MessageType = "STREAM_STATUS"
	TypeError        MessageType = "ERROR"
	TypePing         MessageType = "PING"
	TypePong         MessageType = "PONG"
)

// IncomingMessage represents a message received from a WebSocket client.
type IncomingMessage struct {
	Type     MessageType `json:"type"`
	Content  string      `json:"content,omitempty"`
	Reaction string      `json:"reaction,omitempty"`
}

// OutgoingMessage represents a message sent from the server to WebSocket clients.
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

// Stream represents a live stream session with its metadata and current state.
type Stream struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	TurfID      string    `json:"turfId"`
	Status      string    `json:"status"` // SCHEDULED, LIVE, ENDED
	HLSUrl      string    `json:"hlsUrl"`
	ViewerCount int       `json:"viewerCount"`
	StartedAt   time.Time `json:"startedAt,omitempty"`
}

// AuthUser represents an authenticated user extracted from a verified JWT token.
type AuthUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
