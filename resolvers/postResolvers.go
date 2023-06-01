package resolvers

import (
	"LFGbackend/graph/model"
	"LFGbackend/src"
	"context"
	"errors"
)

func CreatePostResolver(
	ctx context.Context,
	mode model.GameMode,
	player model.PlayerInput,
	need int,
	minRank model.Rank) (string, error) {
	return "", nil
}

func JoinPostResolver(ctx context.Context, player model.PlayerInput, id string) (<-chan *model.Post, error) {
	session := src.SessionContext(ctx)
	if session != nil {
		session.PostId = id

	}
	return nil, errors.New("session id was not initialized")
}
