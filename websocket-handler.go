package main

import (
	"LFGbackend/lfg"
	"LFGbackend/types"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketHandler(ctx context.Context, client *redis.Client, server *lfg.LfgServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		postId, pidOk := vars["id"]
		c, err := r.Cookie("middleware")

		if !pidOk || err != nil {
			log.Println(err)
			return
		}

		bytes, err := client.Get(ctx, c.Value).Bytes()
		session := &types.LfgSession{}

		err = json.Unmarshal(bytes, session)
		if err != nil {
			return
		}

		go func() {

			err = server.JoinPost(types.JoinPostRequest{
				User: types.User{
					Conn:   conn,
					Player: session.Data,
				},
				ClientId: c.Value,
				PostId:   postId,
			})

			if err != nil {
				return
			}

			for {
				messageType, p, err := conn.ReadMessage()
				if err != nil {
					// handle closure
					break
				}

				switch messageType {
				case websocket.TextMessage:
					unmarshallWSData(p)
				}
			}
		}()
	}
}

func unmarshallWSData(
	payload []byte,
) {
	wsData := &types.WsData{}
	err := json.Unmarshal(payload, &wsData)
	if err != nil {
		return
	}
	switch wsData.Type {
	case "":

	}
}
