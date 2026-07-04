# Labs Auction

## Como rodar o projeto

Para subir o projeto basta rodar o seguinte commando

```
docker-compose up
```

## Chamando os endpoints
No projeto temos o arquivo `endpoint.http` com uma chamada de exemplo para criar um auction e outra para listar os auctions criados

```curl
POST http://localhost:8080/auction
Content-Type: application/json

{
    "product_name": "123456",
    "category": "Eletrônicos",
    "description": "Descrição do produto",
    "condition": 0
}


GET http://localhost:8080/auction?status=0
```
