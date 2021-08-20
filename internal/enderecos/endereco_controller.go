package enderecos

import (
	"github.com/raulickis/api-rsec/internal/database/domains"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string
	Errors  []string
}

func ListarEnderecos(ctx *gin.Context) {

	var err error
	enderecos := &[]domains.Endereco{}

	if ctx.Query("user_id") != "" {
		userId, err := strconv.Atoi(ctx.Query("user_id"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		enderecos, err = ListAddressesByUser(ctx.Request.Context(), userId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

	} else {
		enderecos, err = ListAddresses(ctx.Request.Context())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	ctx.JSON(http.StatusOK, enderecos)
}

func ObterEndereco(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Id Inválido")
		return
	}

	endereco, err := GetAddress(ctx.Request.Context(), id)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, "Registro não encontrado")
		} else {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}

	ctx.JSON(http.StatusOK, endereco)
}

func InserirEndereco(ctx *gin.Context) {

	enderecoForm := EnderecoForm{}
	err := ctx.ShouldBind(&enderecoForm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	mapErrors := enderecoForm.ValidarEndereco()
	if len(mapErrors) > 0 {
		response := ResponseError{
			Message           : "Alguns dados informados não foram aceitos, por favor verifique e tente novamente.",
			Errors            : mapErrors,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	endereco := domains.Endereco{}
	endereco.UsuarioID   = enderecoForm.UsuarioID
	endereco.Cep         = enderecoForm.Cep
	endereco.Cep         = enderecoForm.Cep
	endereco.Rua         = enderecoForm.Rua
	endereco.Numero      = enderecoForm.Numero
	endereco.Complemento = &enderecoForm.Complemento
	endereco.Bairro      = enderecoForm.Bairro
	endereco.Cidade      = enderecoForm.Cidade
	endereco.Uf          = enderecoForm.Uf

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	enderecoCriado, err := RegisterAddress(ctx.Request.Context(), &endereco)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, enderecoCriado)
}

func AtualizarEndereco(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Id Inválido")
		return
	}

	endereco, err := GetAddress(ctx.Request.Context(), id)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, "Registro não encontrado")
		} else {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}

	enderecoForm := EnderecoForm{}
	err = ctx.ShouldBind(&enderecoForm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	mapErrors := enderecoForm.ValidarEndereco()
	if len(mapErrors) > 0 {
		response := ResponseError{
			Message           : "Alguns dados informados não foram aceitos, por favor verifique e tente novamente.",
			Errors            : mapErrors,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	endereco.Cep         = enderecoForm.Cep
	endereco.Cep         = enderecoForm.Cep
	endereco.Rua         = enderecoForm.Rua
	endereco.Numero      = enderecoForm.Numero
	endereco.Complemento = &enderecoForm.Complemento
	endereco.Bairro      = enderecoForm.Bairro
	endereco.Cidade      = enderecoForm.Cidade
	endereco.Uf          = enderecoForm.Uf

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	enderecoAtualizado, err := UpdateAddress(ctx.Request.Context(), endereco)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, enderecoAtualizado)
}

func ExcluirEndereco(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Id Inválido")
		return
	}

	enderecoDb, err := GetAddress(ctx.Request.Context(), id)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, "Registro não encontrado")
		} else {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}

	err = DeleteAddress(ctx.Request.Context(), enderecoDb)

	ctx.JSON(http.StatusOK, "Registro deletado")
}