package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/raulickis/api-rsec/config"
	db "github.com/raulickis/api-rsec/tools"
	"sync"
)

type Repository struct {
	dbPostgres *gorm.DB
	once       sync.Once
}

//StartPostgres start the DB
func (r *Repository) Start() error {
	err := db.LoadGormPostGres(
		config.DB_USER,
		config.DB_PASS,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
		false)
	return err
}

//StopPostgres stop the DB
func (r *Repository) Stop() {
	defer r.dbPostgres.Close()
}

// GetInstance returns a unique instance of gorm.DB
func (r *Repository) GetInstance() *gorm.DB {
	r.once.Do(func() {
		var err error
		r.dbPostgres, err = db.GetGormDb()
		if err != nil {
			panic(err.Error())
		}
		r.dbPostgres.SingularTable(true)
		r.dbPostgres.LogMode(true)
	})
	return r.dbPostgres
}
