package resolvers

import (
	"LFGbackend/graph/model"
	"LFGbackend/types"
	"context"
	"gorm.io/gorm"
	"log"
	"time"
)

func GetPlayerResolver(db *gorm.DB, _ context.Context, playerInput model.PlayerInput) (*model.Player, error) {
	name, tag := playerInput.Name, playerInput.Tag
	player := types.GormPlayer{}
	err := db.First(&player, "id = ?", name+tag).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &player.Player, nil
}

func UpsertPlayerResolver(db *gorm.DB, _ context.Context, playerInput model.PlayerInput) (*model.Player, error) {
	name, tag := playerInput.Name, playerInput.Tag
	data, ok := getPlayerData(name, tag)
	if !ok {
		return nil, nil
	}
	player := data.AsPlayer(name, tag)
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
