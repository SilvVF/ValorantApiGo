package lfg

import (
	"LFGbackend/graph/model"
	"LFGbackend/types"
	"log"
)

type PostServer struct {
	ID       string
	Needed   int
	MinRank  model.Rank
	GameMode model.GameMode
	Users    []types.User
}

func NewPostServer(id string, needed int, minRank string, gameMode string) *PostServer {
	return &PostServer{
		ID:       id,
		Needed:   needed,
		MinRank:  model.Rank(minRank),
		GameMode: model.GameMode(gameMode),
		Users:    make([]types.User, 0),
	}
}

func (s *PostServer) Join(r types.JoinPostRequest) error {
	s.Users = append(s.Users, r.User)
	return nil
}

func (s *PostServer) broadcastPlayerJoined(joined types.User) {
	for _, user := range s.Users {
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
