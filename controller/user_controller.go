package controller

import (
	"fmt"

	"github.com/rizkix/wired/config"
	"github.com/rizkix/wired/model"
	"github.com/rizkix/wired/repo"
)

type Controller struct {
	Config config.Config
	Repo   repo.Repo
}

func New(config config.Config, userRepo repo.Repo) Controller {
	return Controller{Config: config, Repo: userRepo}
}

func (c *Controller) Get(ID string) model.Data {
	fmt.Println(c.Config.DbHost, "asdf")
	return c.Repo.Get(ID)
}
