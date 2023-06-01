package graph

import (
	"LFGbackend/src"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Db     *gorm.DB
	Server *src.Server
}
