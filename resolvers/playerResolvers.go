package resolvers

import (
	"LFGbackend/graph/model"
	"LFGbackend/srv"
	"LFGbackend/types"
	"context"
	"gorm.io/gorm"
	"log"
	"time"
)

func GetPlayerResolver(db *gorm.DB, _ context.Context, playerInputs []*model.PlayerInput) ([]*model.Player, error) {
	players := make([]*model.Player, 0, len(playerInputs))
	for _, playerInput := range playerInputs {
		name, tag := playerInput.Name, playerInput.Tag
		player := types.GormPlayer{}
		err := db.First(&player, "id = ?", name+tag).Error
		if err != nil {
			log.Println(err)
			return nil, err
		}
		players = append(players, player.AsPlayer())
	}
	return players, nil
}

func UpsertPlayerResolver(s *srv.Server, db *gorm.DB, ctx context.Context, playerInput model.PlayerInput) (*model.Player, error) {

	session := srv.GetSession(ctx)

	name, tag := playerInput.Name, playerInput.Tag
	data, ok := getPlayerData(name, tag)
	if !ok {
		return nil, nil
	}

	player := data.AsPlayer(name, tag)

	s.ClientIdToPlayer[session.ClientId] = &player

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
