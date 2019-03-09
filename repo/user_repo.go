package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/rizkix/wired/config"

	"github.com/rizkix/wired/model"
)

type Repo struct {
	Conn   *gorm.DB
	Config config.Config
}

func New(config config.Config, db *gorm.DB) (Repo, error) {
	return Repo{Config: config, Conn: db}, nil
}

func (r *Repo) Get(ID string) model.Data {
	a := model.Data{}

	r.Conn.Table("users_permissions").Select("resourceId").First(&a)
	return a
}
