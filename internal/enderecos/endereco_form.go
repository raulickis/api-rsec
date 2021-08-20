package enderecos

import (
	"fmt"
	"github.com/raulickis/api-rsec/internal/utils"
	"gopkg.in/go-playground/validator.v9"
)

type EnderecoForm struct {
	UsuarioID      int              `json:"usuario_id"                 validate:"required"`
	Cep            string           `json:"cep"                        validate:"required,max=9,isvalidzipcode"`
	Rua            string           `json:"rua"                        validate:"required,max=100"`
	Numero         string           `json:"numero"                     validate:"required,max=10"`
	Complemento    string           `json:"complemento"                validate:"max=30"`
	Bairro         string           `json:"bairro"                     validate:"required,max=100"`
	Cidade         string           `json:"cidade"                     validate:"required,max=100"`
	Uf             string           `json:"uf"                         validate:"required,max=2,isfederativeunit"`
}

func (c *EnderecoForm) ValidarEndereco() []string {

	v, trans := utils.ValidatorForm()

	_ = v.RegisterValidation("isvalidzipcode", func(fl validator.FieldLevel) bool {return utils.IsZipCode(fl.Field().String())})
	_ = v.RegisterValidation("isfederativeunit", func(fl validator.FieldLevel) bool {return utils.IsFederativeUnit(fl.Field().String())})

	err := v.Struct(c)
	var mapErrors []string
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e.Translate(trans))
			mapErrors = append(mapErrors, e.Translate(trans))
		}
	}

	return mapErrors
}
