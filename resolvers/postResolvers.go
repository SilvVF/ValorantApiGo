package resolvers

import (
	"LFGbackend/graph/model"
	"context"
)

func CreatePostResolver(
	ctx context.Context,
	mode model.GameMode,
	player model.PlayerInput,
	need int,
	minRank model.Rank) (string, error) {
	return "", nil
}

func JoinPostResolver(ctx context.Context, player model.PlayerInput, id string) (bool, error) {
	return false, nil
}
