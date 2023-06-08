package graph

import (
	"LFGbackend/lfg"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Db             *gorm.DB
	Server         *lfg.LfgServer
	SessionManager *lfg.SessionManager
}
