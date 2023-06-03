package srv

import (
	"LFGbackend/graph/model"
	"LFGbackend/types"
	"errors"
	"sync"
	"time"
)

type Server struct {
	Posts            map[string]*Post
	ClientIdToPostId map[string]string
	ClientIdToPlayer map[string]*model.Player
	mutex            sync.Mutex
}

func NewServer() *Server {

	server := &Server{
		Posts:            make(map[string]*Post),
		ClientIdToPostId: make(map[string]string),
		ClientIdToPlayer: make(map[string]*model.Player),
		mutex:            sync.Mutex{},
	}
	go func() {
		for {
			time.Sleep(time.Hour * 3)
			for _, p := range server.Posts {
				if len(p.users) <= 0 || p.Closed {
					delete(server.Posts, p.Id)
				}
			}
		}
	}()
	return server
}

func (s *Server) CreatePost(post model.Post, clientId string) {
	ps := &Post{
		Id:       post.ID,
		Needed:   post.Needed,
		MinRank:  post.MinRank,
		Players:  make([]*model.Player, 0),
		Messages: make([]*model.Message, 0),
		users:    make(map[string]*types.User),
		mutex:    sync.Mutex{},
		creator:  clientId,
	}
	s.Posts[post.ID] = ps
}

func (s *Server) DeletePlayer(session *types.PostSession) {
	delete(s.ClientIdToPlayer, session.ClientId)
}

func (s *Server) LeavePost(id string, user *types.User) {
	go func() {
		post := s.getPostForUser(id)
		if post != nil {
			post.Leave(user)
			s.mutex.Lock()
			delete(s.ClientIdToPostId, user.Info.ClientId)
			if len(post.Players) <= 0 {
				delete(s.Posts, post.Id)
			}
			s.mutex.Unlock()
		}
	}()
}

func (s *Server) JoinPost(id string, user *types.User) error {
	s.mutex.Lock()
	post := s.Posts[id]
	s.mutex.Unlock()
	if post != nil {
		return errors.New("post was not found for given id")
	}
	err := post.Join(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) getPostForUser(id string) *Post {
	s.mutex.Lock()
	postId, ok := s.ClientIdToPostId[id]
	s.mutex.Unlock()
	if !ok {
		return nil
	}
	return s.Posts[postId]
}

type Post struct {
	Id       string
	Needed   int
	MinRank  model.Rank
	Players  []*model.Player
	Messages []*model.Message
	Closed   bool
	creator  string
	users    map[string]*types.User
	mutex    sync.Mutex
}

func (p *Post) SendMessage(message *model.Message) {
	p.Messages = append(p.Messages, message)
	go p.broadcastState()
}

func (p *Post) Join(joiner *types.User) error {
	if p.Needed > len(p.Players) {
		return errors.New("room is already full")
	}
	if GetValue(p.MinRank) < GetValue(model.Rank(joiner.Info.Player.Rank)) {
		return errors.New("rank to low")
	}
	p.mutex.Lock()
	p.Needed--
	p.users[joiner.Info.ClientId] = joiner
	p.mutex.Unlock()
	p.Players = append(p.Players, joiner.Info.Player)

	go p.broadcastState()

	return nil
}

func (p *Post) Leave(leaver *types.User) {

	remove := func(player *model.Player, players []*model.Player) []*model.Player {
		filtered := make([]*model.Player, 0, len(players))
		for _, p := range players {
			if p.Name+p.Tag != player.Name+player.Tag {
				filtered = append(filtered, p)
			}
		}
		return filtered
	}

	p.mutex.Lock()
	p.Needed++
	delete(p.users, leaver.Info.ClientId)
	p.mutex.Unlock()

	p.Players = remove(leaver.Info.Player, p.Players)

	go p.broadcastState()
}

func (p *Post) broadcastState() {
	creatorHere := false
	for _, user := range p.users {
		if user.Info.ClientId == p.creator {
			creatorHere = true
			p.Closed = true
		}
	}
	p.mutex.Lock()
	for _, user := range p.users {
		user.State <- &model.Post{
			ID:       p.Id,
			Players:  p.Players,
			Needed:   p.Needed,
			MinRank:  p.MinRank,
			Messages: p.Messages,
			Closed:   len(p.users) > 0 || !creatorHere,
		}
	}
	p.mutex.Unlock()
}

func GetValue(r model.Rank) int {
	switch r {
	case "RADIANT":
		return 25
	case "IMMORTAL3":
		return 24
	case "IMMORTAL2":
		return 23
	case "IMMORTAL1":
		return 22
	case "ASCENDANT3":
		return 21
	case "ASCENDANT2":
		return 20
	case "ASCENDANT1":
		return 19
	case "DIAMOND3":
		return 18
	case "DIAMOND2":
		return 17
	case "DIAMOND1":
		return 16
	case "PLAT3":
		return 15
	case "PLAT2":
		return 14
	case "PLAT1":
		return 13
	case "GOLD3":
		return 12
	case "GOLD2":
		return 11
	case "GOLD1":
		return 10
	case "SILVER3":
		return 9
	case "SILVER2":
		return 8
	case "SILVER1":
		return 7
	case "BRONZE3":
		return 6
	case "BRONZE2":
		return 5
	case "BRONZE1":
		return 4
	case "IRON3":
		return 3
	case "IRON2":
		return 2
	case "IRON1":
		return 1
	case "UNRANKED":
		return 0
	}
	return 0
}
