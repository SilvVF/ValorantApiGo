package resolvers

import (
	"LFGbackend/graph/model"
	"LFGbackend/srv"
	"LFGbackend/types"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func GetPostsResolver(
	server *srv.Server,
	_ context.Context,
	page int,
	count int,
) ([]*model.PostInfo, error) {

	start := page - 1*count
	end := page * count

	var posts = make([]*model.PostInfo, 0, count)
	var i = 0

	for _, p := range server.Posts {
		switch {
		case i >= start:
			names := make([]string, 0)
			for _, player := range p.Players {
				names = append(names, player.Name+player.Tag)
			}
			post := &model.PostInfo{
				ID:          p.Id,
				Needed:      p.Needed,
				MinRank:     p.MinRank,
				PlayerNames: names,
				Closed:      p.Needed == 0,
			}
			posts = append(posts, post)
		case i > end:
			return posts, nil
		}
		i++
	}
	return posts, nil
}

func CreatePostResolver(
	ctx context.Context,
	server *srv.Server,
	_ *gorm.DB,
	_ model.GameMode,
	_ model.PlayerInput,
	need int,
	minRank model.Rank) (<-chan *model.Post, error) {

	id := uuid.NewString()
	clientId := srv.GetSession(ctx).ClientId
	player, ok := server.ClientIdToPlayer[clientId]
	if !ok {
		return nil, errors.New("couldn't player with clientId" + clientId)
	}

	server.CreatePost(model.Post{
		ID:       id,
		Players:  nil,
		Needed:   need,
		MinRank:  minRank,
		Messages: nil,
		Closed:   false,
	}, clientId)

	user := &types.User{
		Info: types.UserInfo{
			ClientId: clientId,
			Player:   player,
		},
		State: make(chan *model.Post),
	}
	err := server.JoinPost(id, user)
	if err != nil {
		return nil, err
	}

	postState := make(chan *model.Post)

	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				server.LeavePost(id, user)
				return
			case state := <-user.State:
				postState <- state
			}
		}
	}()
	return postState, nil
}

func JoinPostResolver(
	ctx context.Context,
	server *srv.Server,
	_ model.PlayerInput,
	id string,
) (<-chan *model.Post, error) {

	session := srv.GetSession(ctx)
	player, ok := server.ClientIdToPlayer[session.ClientId]

	if !ok || session == nil {
		return nil, errors.New("player not attached to session")
	}

	postState := make(chan *model.Post)
	user := &types.User{
		Info: types.UserInfo{
			ClientId: session.ClientId,
			Player:   player,
		},
		State: make(chan *model.Post),
	}
	err := server.JoinPost(id, user)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				server.LeavePost(id, user)
				return
			case state := <-user.State:
				postState <- state
			}
		}
	}()
	return postState, nil
}

func SendMessageResolver(m string, ctx context.Context, s *srv.Server) (bool, error) {

	session := srv.GetSession(ctx)
	if session == nil {
		return false, errors.New("session not found")
	}
	postId, ok := s.ClientIdToPostId[session.ClientId]
	if !ok {
		return false, errors.New("unable to find joined post for clientId")
	}
	post, ok := s.Posts[postId]
	if !ok {
		return false, errors.New("post does not exist anymore")
	}
	post.SendMessage(&model.Message{
		Sender: s.ClientIdToPlayer[session.ClientId],
		Text:   m,
		SentAt: time.Now().String(),
	})
	return true, nil
}
