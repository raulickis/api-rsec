package usuarios

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

func ListarUsuarios(ctx *gin.Context) {

	usuarios, err := ListUsers(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, usuarios)
}

func ObterUsuario(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Id Inválido")
		return
	}

	usuario, err := GetUser(ctx.Request.Context(), id)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, "Registro não encontrado cara!")
		} else {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}

	ctx.JSON(http.StatusOK, usuario)
}

func InserirUsuario(ctx *gin.Context) {

	usuario := domains.Usuario{}
	//err := ctx.ShouldBind(&usuario)


	// Opcional - Inicio da Validacao de Campos

	form := UsuarioForm{}
	err := ctx.ShouldBind(&form)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	mapErrors := form.ValidarUsuario()
	if len(mapErrors) > 0 {
		response := ResponseError{
			Message           : "Alguns dados informados não foram aceitos, por favor verifique e tente novamente.",
			Errors            : mapErrors,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	usuario.Nome      = form.Nome
	usuario.Documento = form.Documento
	usuario.Email     = form.Email
	usuario.Ddd       = &form.Ddd
	usuario.Telefone  = &form.Telefone
	usuario.FotoUrl   = form.FotoUrl
	for _, enderecoForm := range  form.Enderecos {
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
		endereco.Cep         = enderecoForm.Cep
		endereco.Rua         = enderecoForm.Rua
		endereco.Numero      = enderecoForm.Numero
		endereco.Complemento = &enderecoForm.Complemento
		endereco.Bairro      = enderecoForm.Bairro
		endereco.Cidade      = enderecoForm.Cidade
		endereco.Uf          = enderecoForm.Uf
		usuario.Enderecos =  append(usuario.Enderecos, endereco)
	}

	// Opcional - Fim da Validacao de Campos

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	usuarioCriado, err := RegisterUser(ctx.Request.Context(), &usuario)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, usuarioCriado)
}

func AtualizarUsuario(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Id Inválido")
		return
	}

	usuario, err := GetUser(ctx.Request.Context(), id)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, "Registro não encontrado")
		} else {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}

	//err = ctx.ShouldBind(&usuario)


	// Opcional - Inicio da Validacao de Campos

	form := UsuarioForm{}
	err = ctx.ShouldBind(&form)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	mapErrors := form.ValidarUsuario()
	if len(mapErrors) > 0 {
		response := ResponseError{
			Message           : "Alguns dados informados não foram aceitos, por favor verifique e tente novamente.",
			Errors            : mapErrors,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	usuario.Nome      = form.Nome
	usuario.Documento = form.Documento
	usuario.Email     = form.Email
	usuario.Ddd       = &form.Ddd
	usuario.Telefone  = &form.Telefone
	usuario.FotoUrl   = form.FotoUrl

	// Opcional - Fim da Validacao de Campos


	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	usuarioAtualizado, err := UpdateUser(ctx.Request.Context(), usuario)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, usuarioAtualizado)
}

func ExcluirUsuario(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Id Inválido")
		return
	}

	usuarioDb, err := GetUser(ctx.Request.Context(), id)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, "Registro não encontrado")
		} else {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}

	err = DeleteUser(ctx.Request.Context(), usuarioDb)

	ctx.JSON(http.StatusOK, "Registro deletado")
}