package migrations

import (
	"log"
	db "github.com/raulickis/api-rsec/internal/database"
	"gopkg.in/gormigrate.v1"
)

var migrations = []*gormigrate.Migration{
	&migration_0001,
	&migration_0002,
}

func RunMigrations() {
	postgres := &db.Repository{}
	db := postgres.GetInstance()
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Não conseguimos realizar a criação da estrutura de dados: %v", err)
	}
	log.Printf("Criação da estrutura de dados finalizada!")
}
