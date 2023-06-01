package resolvers

import (
	"LFGbackend/graph/model"
	"LFGbackend/types"
	"context"
	"gorm.io/gorm"
	"log"
	"time"
)

func GetPlayerResolver(db *gorm.DB, ctx context.Context, playerInput model.PlayerInput) (*model.Player, error) {
	name, tag := playerInput.Name, playerInput.Tag
	var player = types.GormPlayer{}
	db.Where("name = ? AND tag = ?", name, tag).First(&player)
	return &player.Player, nil
}

func UpsertPlayerResolver(db *gorm.DB, ctx context.Context, playerInput model.PlayerInput) (*model.Player, error) {
	name, tag := playerInput.Name, playerInput.Tag
	player := getPlayerData(name, tag).AsPlayer(name, tag)
	gormPlayer := types.GormPlayer{
		Id:        name + tag,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Player:    player,
	}
	err := db.Save(gormPlayer).Error
	if err != nil {
		log.Println(err)
	}
	return &player, nil
}
