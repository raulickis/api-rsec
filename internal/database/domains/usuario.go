package domains

import (
	"time"
)

type Usuario struct {
	ID             int        `json:"id"                     gorm:"primary_key;auto_increment;not_null"`
	Nome           string     `json:"nome"`
	Documento      string     `json:"documento"              gorm:"index:idx_documento"`
	Email          string     `json:"email"                  gorm:"index:idx_email"`
	Ddd            *string    `json:"ddd"`
	Telefone       *string    `json:"telefone,omitempty"`
	FotoUrl        string     `json:"foto_url"`
	DataCadastro   time.Time  `json:"data_cadastro"          gorm:"DEFAULT:current_timestamp"`
	Enderecos      []Endereco `json:"enderecos,omitempty"`
}

