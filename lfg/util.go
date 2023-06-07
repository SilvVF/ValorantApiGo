package lfg

import (
	"LFGbackend/types"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func writeWsData[T any](conn *websocket.Conn, data T, name string) error {
	message, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = conn.WriteJSON(types.WsData{
		Type: name,
		Raw:  message,
	})
	return err
}
