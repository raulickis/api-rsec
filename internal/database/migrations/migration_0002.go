package migrations

import (
	"github.com/raulickis/api-rsec/internal/database/domains"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var migration_0002 = gormigrate.Migration{
	ID: "migration_0002",
	Migrate: func(tx *gorm.DB) error {

		// Criacao dos primeiros registros para testes

		// Avatars disponíveis em https://pixabay.com/images/search/avatar

		var ddd string
		var telefone string
		var complemento string
		var usuario domains.Usuario
		var endereco domains.Endereco
		var enderecos []domains.Endereco

		ddd = "11"
		telefone = "96969-6969"
		usuario = domains.Usuario{Nome:"Joao Cunha", Email:"joao.cunha@gmail.com", Documento:"280.947.150-90", FotoUrl:"https://cdn.pixabay.com/photo/2016/04/01/10/11/avatar-1299805__340.png", Ddd:&ddd, Telefone:&telefone}
		tx.NewRecord(usuario)
		tx.Create(&usuario)

		enderecos = []domains.Endereco{}

		complemento = "apto 64B"
		endereco = domains.Endereco{Cep:"06132-000", Rua: "Avenida Paulista", Numero:"1270", Complemento: &complemento, Bairro:"Jardim Paulista", Cidade:"São Paulo", Uf: "SP"}
		enderecos = append(enderecos, endereco)

		endereco = domains.Endereco{Cep:"05444-090", Rua: "Avenida Rebouças", Numero:"3270", Bairro:"Jardim Europa", Cidade:"São Paulo", Uf: "SP"}
		enderecos = append(enderecos, endereco)

		ddd = "11"
		telefone = "97778-1111"
		usuario = domains.Usuario{Enderecos: enderecos, Nome:"Maria das Dores", Email:"maria.dores@gmail.com", Documento:"658.426.670-29", FotoUrl:"https://cdn.pixabay.com/photo/2015/12/13/20/43/doll-1091702__340.jpg", Ddd:&ddd, Telefone:&telefone}
		tx.NewRecord(usuario)
		tx.Create(&usuario)


		ddd = "43"
		telefone = "3200-1223"
		usuario = domains.Usuario{Nome:"Thiago Alencar", Email:"joao.cunha@gmail.com", Documento:"737.626.700-05", FotoUrl:"https://cdn.pixabay.com/photo/2018/04/28/13/18/man-3357275__340.png", Ddd:&ddd, Telefone:&telefone}
		tx.NewRecord(usuario)
		tx.Create(&usuario)

		enderecos = []domains.Endereco{}

		complemento = "casa"
		endereco = domains.Endereco{Cep:"74445-000", Rua: "Rua Moxarim Iosef Hala", Numero:"123", Complemento: &complemento, Bairro:"Jardim Panama", Cidade:"Campo Grande", Uf: "MS"}
		enderecos = append(enderecos, endereco)

		ddd = "67"
		telefone = "98675-1234"
		usuario = domains.Usuario{Enderecos:enderecos, Nome:"Marcos Paulo", Email:"raulickis@gmail.com", Documento:"348.006.000-57", FotoUrl:"https://cdn.pixabay.com/photo/2017/12/16/06/41/avatar-3022215__340.jpg", Ddd:&ddd, Telefone:&telefone}
		tx.NewRecord(usuario)
		tx.Create(&usuario)


		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}