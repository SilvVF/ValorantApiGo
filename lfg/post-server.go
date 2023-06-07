package lfg

import (
	"LFGbackend/types"
	"log"
)

type PostServer struct {
	users []types.User
}

func (s *PostServer) Join(r types.JoinPostRequest) {
	s.users = append(s.users, r.User)
}

func (s *PostServer) broadcastPlayerJoined(joined types.User) {
	for _, user := range s.users {
		if user.Conn == joined.Conn {
			continue
		}
		err := writeWsData(
			user.Conn,
			types.PlayerJoined{Player: joined.Player},
			"PlayerJoined",
		)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
