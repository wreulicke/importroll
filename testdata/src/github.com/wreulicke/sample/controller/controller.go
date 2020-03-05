package controller

import (
	"github.com/wreulicke/sample/repository" // want "cannot import this package"
)

type Controller struct {
	Repository repository.Repository
}
