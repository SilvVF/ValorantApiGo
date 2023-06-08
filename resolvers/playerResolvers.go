package resolvers

import (
	"LFGbackend/graph/model"
	"LFGbackend/lfg"
	"LFGbackend/middleware"
	"LFGbackend/trn"
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

func SignInAsPlayerResolver(ctx context.Context, db *gorm.DB, playerInput model.PlayerInput, manager *lfg.SessionManager) (*model.Player, error) {

	client := middleware.ForContext(ctx)

	name, tag := playerInput.Name, playerInput.Tag
	data, ok := trn.GetPlayerData(name, tag)
	if !ok {
		return nil, nil
	}

	player := data.AsPlayer(name, tag)
	saved := types.GormPlayer{}

	err := db.First(&saved, "id = ?", name+tag).Error

	if err != nil {
		err = createPlayerInDb(db, player, name+tag)
		if err != nil {
			return nil, err
		}
	} else {
		db.Save(&types.GormPlayer{
			Id:        saved.Id,
			CreatedAt: saved.CreatedAt,
			UpdatedAt: time.Now(),
			Player:    player,
		})
	}

	manager.Set(client.Id, player)

	return &player, nil
}

func createPlayerInDb(db *gorm.DB, player model.Player, key string) error {
	gormPlayer := types.GormPlayer{
		Id:        key,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Player:    player,
	}
	err := db.Create(gormPlayer).Error
	return err
}
