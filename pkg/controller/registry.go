package controller

import "github.com/jinzhu/gorm"

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() AppController {
	return AppController{
		Auth: r.NewAuthController(r.db),
	}
}
