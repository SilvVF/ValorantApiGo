package resolvers

import (
	"LFGbackend/graph/model"
	"LFGbackend/src"
	"LFGbackend/types"
	"context"
	"errors"
	"gorm.io/gorm"
)

func CreatePostResolver(
	ctx context.Context,
	server *src.Server,
	db *gorm.DB,
	mode model.GameMode,
	player model.PlayerInput,
	need int,
	minRank model.Rank) (string, error) {
	return "", nil
}

func JoinPostResolver(
	ctx context.Context,
	server *src.Server,
	input model.PlayerInput,
	id string,
) (<-chan *model.Post, error) {

	session := src.SessionContext(ctx)

	if session != nil {
		if session.Player == nil {
			return nil, errors.New("player not attached to session")
		}
		postState := make(chan *model.Post)
		change := make(chan *model.Post)

		user := &types.User{
			Info: types.UserInfo{
				ClientId: session.ClientId,
				Player:   session.Player,
			},
		}

		err := server.JoinPost(
			id,
			user,
		)
		if err != nil {
			return nil, err
		}
		go func() {
			for {
				select {
				case <-ctx.Done():
					server.LeavePost(id, user)
					return
				case state := <-change:
					postState <- state
				}
			}
		}()
		return postState, nil
	}
	return nil, errors.New("session id was not initialized")
}
