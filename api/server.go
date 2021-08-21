package api

import (
	"github.com/gin-gonic/gin"
	"github.com/raulickis/api-rsec/config"
	"github.com/raulickis/api-rsec/internal/enderecos"
	"github.com/raulickis/api-rsec/internal/usuarios"
	"net/http"

	//"github.com/raulickis/api-rsec/tools"
)

func Run() {
	r := SetupRoutes()
	_ = r.Run(":" + config.APP_PORT)
}

func SetupRoutes() *gin.Engine {

	router := gin.New()
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health/check"),
		gin.Recovery(),
	)
	router.GET("/health/check", CheckHealth)

	router.Use(Interceptor())

	// API
	cadastroRouter := router.Group("/cadastro/usuario")
	{
		cadastroRouter.POST("", usuarios.InserirUsuario)
		cadastroRouter.GET("", usuarios.ListarUsuarios)
		cadastroRouter.GET("/:id", usuarios.ObterUsuario)
		cadastroRouter.PUT("/:id", usuarios.AtualizarUsuario)
		cadastroRouter.DELETE("/:id", usuarios.ExcluirUsuario)
	}

	cadastroEnderecosRouter := router.Group("/cadastro/endereco")
	{
		cadastroEnderecosRouter.POST("", enderecos.InserirEndereco)
		cadastroEnderecosRouter.GET("", enderecos.ListarEnderecos)
		cadastroEnderecosRouter.GET("/:id", enderecos.ObterEndereco)
		cadastroEnderecosRouter.PUT("/:id", enderecos.AtualizarEndereco)
		cadastroEnderecosRouter.DELETE("/:id", enderecos.ExcluirEndereco)
	}

	return router
}

func CheckHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func Interceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Access Token Required"})
			ctx.Abort()
			return
		}
		if token != config.TOKEN {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Access Token Invalid"})
			ctx.Abort()
			return
		}
	}
}