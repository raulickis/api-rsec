package domains

import (
	"time"
)

type Endereco struct {
	ID             int       `json:"id"                       gorm:"primary_key;auto_increment;not_null"`
	UsuarioID      int       `json:"usuario_id"`
	Usuario        Usuario   `json:"-"                        gorm:"constraint:OnDelete:CASCADE;"`
	Cep            string    `json:"cep"                      gorm:"index:idx_cep"`
	Rua            string    `json:"rua"`
	Numero         string    `json:"numero"`
	Complemento    *string   `json:"complemento,omitempty"`
	Bairro         string    `json:"bairro"`
	Cidade         string    `json:"cidade"`
	Uf             string    `json:"uf"`
	DataCadastro   time.Time `json:"data_cadastro"             gorm:"DEFAULT:current_timestamp"`
}
