package lfg

import (
	"LFGbackend/types"
	"errors"
	"sync"
)

type LfgServer struct {
	posts            map[string]*PostServer
	clientIdToPostId map[string]string
	mutex            sync.Mutex
}

func NewLfgServer() *LfgServer {
	return &LfgServer{
		posts:            make(map[string]*PostServer),
		clientIdToPostId: make(map[string]string),
		mutex:            sync.Mutex{},
	}
}

func (s *LfgServer) JoinPost(request types.JoinPostRequest) error {
	s.mutex.Lock()
	post, ok := s.posts[request.PostId]
	if !ok {
		return errors.New("post with id:" + request.PostId + " does not exist")
	}
	post.Join(request)
	s.mutex.Unlock()
	return nil
}

func (s *LfgServer) Start() {

}
