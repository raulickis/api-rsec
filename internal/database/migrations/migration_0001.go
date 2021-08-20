package migrations

import (
	"github.com/raulickis/api-rsec/internal/database/domains"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var migration_0001 = gormigrate.Migration{
	ID: "migration_0001",
	Migrate: func(tx *gorm.DB) error {

		// Criacao automatica das tabelas

		tx.AutoMigrate(
			&domains.Usuario{},
			&domains.Endereco{},
		)

		// Criacao do relacionamento no banco de dados entre as tabelas usuario e endereco (Foreign Key)

		tx.Model(&domains.Endereco{}).AddForeignKey("usuario_id", "usuario(id)", "CASCADE", "CASCADE")

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}