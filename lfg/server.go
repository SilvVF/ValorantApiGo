package lfg

import (
	"LFGbackend/types"
	"errors"
	"github.com/google/uuid"
	"sync"
)

type LfgServer struct {
	posts            map[string]*PostServer
	clientIdToPostId map[string]string
	mutex            sync.Mutex
}

func (s *LfgServer) CreatePost(needed int, gameMode string, minRank string) string {
	postId := uuid.NewString()
	postServer := NewPostServer(postId, needed, minRank, gameMode)
	s.mutex.Lock()
	s.posts[postId] = postServer
	s.mutex.Unlock()
	return postId
}

func (s *LfgServer) GetPosts() []*PostServer {
	postList := make([]*PostServer, 0, len(s.posts))
	s.mutex.Lock()
	for _, p := range s.posts {
		postList = append(postList, p)
	}
	s.mutex.Unlock()
	return postList
}

func NewLfgServer() *LfgServer {
	server := &LfgServer{
		posts:            make(map[string]*PostServer),
		clientIdToPostId: make(map[string]string),
		mutex:            sync.Mutex{},
	}
	server.posts["dev"] = NewPostServer("dev", 3, "Unranked", "Ranked")
	return server
}

func (s *LfgServer) JoinPost(request types.JoinPostRequest) error {
	s.mutex.Lock()
	post, ok := s.posts[request.PostId]
	s.mutex.Unlock()
	if !ok {
		return errors.New("post with id:" + request.PostId + " does not exist")
	}
	err := post.Join(request)
	if err != nil {
		s.clientIdToPostId[request.ClientId] = request.PostId
	}
	return err
}
