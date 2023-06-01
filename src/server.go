package src

import (
	"LFGbackend/graph/model"
	"LFGbackend/types"
)

type Server struct {
	connections    <-chan types.JoinRequest
	disconnections <-chan types.LeaveRequest
	posts          map[string]model.Post
	users          map[string]model.Player
}

func NewServer() *Server {
	return &Server{
		connections:    make(chan types.JoinRequest, 2),
		disconnections: make(chan types.LeaveRequest, 2),
		posts:          make(map[string]model.Post, 0),
		users:          make(map[string]model.Player, 0),
	}
}

func (s *Server) collect() {
	for {
		select {
		case r := <-s.connections:
			s.users[r.UserId] = r.Player
			post, exists := s.posts[r.PostId]
			if exists {
				post.Players = append(post.Players, &r.Player)
			}
		case r := <-s.disconnections:
			delete(s.users, r.UserId)
			post, exists := s.posts[r.PostId]
			if exists {
				post.Players = remove(&r.Player, post.Players)
			}
		default:
		}
	}
}

func remove(player *model.Player, slice []*model.Player) []*model.Player {
	for i, v := range slice {
		if v == player {
			return append(slice[0:i], slice[i+1:]...)
		}
	}
	return slice
}
