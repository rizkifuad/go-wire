package controller

import (
	"github.com/rizkix/wired/model"
	"github.com/rizkix/wired/repo"
)

type Controller struct {
	Repo repo.Repo
}

func New(userRepo repo.Repo) Controller {
	return Controller{Repo: userRepo}
}

func (c *Controller) Get(ID string) model.Data {
	return c.Repo.Get(ID)
}
