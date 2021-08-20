# api-rsec

API em **GO v1.14** responsável por registrar métricas de dispositivos de segurança IoT.


Para fazer o download das depências e limpar dependências não utilizadas:

``` 
go14 mod tidy
```

## Para executar localmente:
 
1 - Subir o banco de dados
```
docker-compose up postgres
```

2 - Subir o aplicativo localmente pelo GO

```
go run cmd/main.go
```

## Para preparar a versão para implantação:
 
1 - Compilar o programa e gerar a Imagem Docker
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd
docker build -t raulickis/api-rsec .
```

2 - Executar tudo pelo Docker Compose

```
docker-compose up
```

## Como testar a aplicacão:

1 - (Opcional) Para fazer um teste rápido, executar o comando **curl** abaixo:
```
curl http://localhost:9990/cadastro/usuario
```

2 - Importar no Postman o arquivo **rsec.postman_collection.json**

3 - Executar os endpoints para cadastro de usuarios e enderecos


## Instalar o Elastic Agent
```
go get -u go.elastic.co/apm
export ELASTIC_APM_SERVER_URL=http://elastic-apm-server-apm-http.default:8200
```
Declarar a variável de ambiente ELASTIC_APM_SERVER_URL nas configurações e instumentar o código
https://www.elastic.co/guide/en/apm/agent/go/current/getting-started.html
https://www.elastic.co/guide/en/apm/agent/go/current/builtin-modules.html

Frameworks Autocontidos para ElasticAPM:
```
import (
	"go.elastic.co/apm/module/apmgin"
)

func main() {
	engine := gin.New()
	engine.Use(apmgin.Middleware(engine))
	...
}



import (
	"go.elastic.co/apm/module/apmgorm"
	_ "go.elastic.co/apm/module/apmgorm/dialects/postgres"
)

func main() {
	db, err := apmgorm.Open("postgres", "")
	...
	db = apmgorm.WithContext(ctx, db)
	db.Find(...) // creates a "SELECT FROM <foo>" span
}



import (
	"net/http"

	"github.com/go-redis/redis"

	"go.elastic.co/apm/module/apmgoredis"
)

var redisClient *redis.Client // initialized at program startup

func handleRequest(w http.ResponseWriter, req *http.Request) {
	// Wrap and bind redisClient to the request context. If the HTTP
	// server is instrumented with Elastic APM (e.g. with apmhttp),
	// Redis commands will be reported as spans within the request's
	// transaction.
	client := apmgoredis.Wrap(redisClient).WithContext(req.Context())
	...
}

