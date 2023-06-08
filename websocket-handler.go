package main

import (
	"LFGbackend/graph/model"
	"LFGbackend/lfg"
	"LFGbackend/middleware"
	"LFGbackend/types"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketCreatePostHandler(manager *lfg.SessionManager, server *lfg.LfgServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		vars := mux.Vars(r)

		client := middleware.ForContext(r.Context())

		player, ok := manager.Get(client.Id)

		needed, nOk := vars["needed"]
		gameMode, gmOk := vars["gamemode"]
		minRank, mrOk := vars["minrank"]

		if !nOk || !gmOk || !mrOk || !ok {
			return
		}

		need, err := strconv.Atoi(needed)
		if err != nil {
			log.Println(err)
		}

		postId := server.CreatePost(need, gameMode, minRank)

		go handleWS(conn, client.Id, player, postId, server)
	}
}

func websocketJoinPostHandler(manager *lfg.SessionManager, server *lfg.LfgServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)

		vars := mux.Vars(r)

		if err != nil {
			log.Println(err)
			return
		}
		postId, pidOk := vars["id"]

		client := middleware.ForContext(r.Context())

		if !pidOk || client == nil {
			log.Println(err)
			return
		}
		playerData, ok := manager.Get(client.Id)
		if !ok {
			return
		}

		go handleWS(conn, client.Id, playerData, postId, server)
	}
}

func handleWS(conn *websocket.Conn, clientID string, player model.Player, postId string, server *lfg.LfgServer) {

	err := server.JoinPost(types.JoinPostRequest{
		User: types.User{
			Conn:   conn,
			Player: player,
		},
		ClientId: clientID,
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
			err = conn.WriteJSON(p)
			if err != nil {
				log.Println(err)
			}
		}
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
