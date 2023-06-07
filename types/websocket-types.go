package types

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

type WsData struct {
	Type string `json:"type"`
	Raw  json.RawMessage
}

type PlayerJoined struct {
	Player PlayerData
}

type User struct {
	Conn   *websocket.Conn
	Player PlayerData
}
