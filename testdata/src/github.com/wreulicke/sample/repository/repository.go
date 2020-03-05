package repository

import (
	"github.com/wreulicke/sample/model"
	"github.com/wreulicke/sample/service" // want "cannot import this package"
)

type Repository interface {
	FindModel() model.Model
	Invalid() service.Service
}
