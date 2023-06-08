package types

import (
	"LFGbackend/graph/model"
	"encoding/json"
	"github.com/gorilla/websocket"
)

type WsData struct {
	Type string `json:"type"`
	Raw  json.RawMessage
}

type PlayerJoined struct {
	Player model.Player
}

type User struct {
	Conn   *websocket.Conn
	Player model.Player
}
