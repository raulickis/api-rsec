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
curl http://localhost:9990/cadastro/usuario -H 'token: ZjFhNTUyZmI2YjYzNDI0ZmRmNDUzZDAx' 
```

2 - Importar no Postman o arquivo **rsec.postman_collection.json**

3 - Executar os endpoints para cadastro de usuarios e enderecos


## Servidor WEB

Auto-iniciar o docker-compose com o Linux:
```
sudo cp docker-compose-rsec.service /etc/systemd/system/docker-compose-rsec.service
sudo systemctl enable docker-compose-rsec
sudo systemctl start docker-compose-rsec
sudo systemctl status docker-compose-rsec
```
 
Como configurar o Nginx de um servidor para redirecionar para a api:
```
location /rsec {
    rewrite ^/rsec(.*) $1 break;
    proxy_pass http://127.0.0.1:9990;
}
```

