package resolvers

import (
	"LFGbackend/graph/model"
	"context"
)

func GetPlayerResolver(ctx context.Context, playerInput model.PlayerInput) (*model.Player, error) {
	name, tag := playerInput.Name, playerInput.Tag
	player := getPlayerData(name, tag).AsPlayer(name, tag)
	return &player, nil
}

func UpsertPlayerResolver(ctx context.Context, playerInput model.PlayerInput) (*model.Player, error) {
	name, tag := playerInput.Name, playerInput.Tag
	player := getPlayerData(name, tag).AsPlayer(name, tag)
	return &player, nil
}
