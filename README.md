# Cheapest Trip #

O propósito do programa é fornecer a informação da rota mais barata independente do número de conexões necessárias para percorrê-la.
É ncessário inserir as rotas através de um arquivo de entrada.

## Input Example ##
```csv
GRU,BRC,10
BRC,SCL,5
GRU,CDG,75
GRU,SCL,20
GRU,ORL,56
ORL,CDG,5
SCL,ORL,20
```

### Execução ###
Execute o arquivo passando o arquivo de rotas como argumento.

##### Desenvolvimento
```sh
$ go run cmd/main.go input-routes.csv
```

##### Compilado
```sh
go build -o ./build/cheapest ./cmd
./build/cheapest input-file.csv
```
Após a mensagem solicitar a informação da rota, entrar com o valor no formato AAA-BBB:
```
please enter the route: GRU-CDG
best route: GRU - BRC - SCL - ORL - CDG > $40.00
```

### Testes ###
```sh
go test ./...
```

### API Rest
Após rodar o comando de execução a API estará disponível na porta 9000.
A API é muito simples e possue apenas um endpoint na url http://localhost:9000/routes

#### Buscar a rota mais barata
Enviar uma requisição GET e informar o parametro "route" na query string com a rota de partida e a rota de chegada separada por um hifen.

Exemplo:
GET/ http://localhost:9000/routes?route=GRU-CDG

Resposta:
```
HTTP/1.1 200 OK

{
  "departure": "GRU",
  "arrival": "CDG",
  "best_route": [
    "GRU",
    "BRC",
    "SCL",
    "ORL",
    "CDG"
  ],
  "cost": 40
}
```

#### Cadastro de rotas
Enviar uma requisição POST e passr a rota no body da requisição. O cabeçalho Content-Type deverá ter o valor application/json

Exemplo:
POST/ http://localhost:9000/routes
```
{
  "departure": "ABC",
  "arrival": "DEF",
  "cost": 5
}
```
Resposta:
```
HTTP/1.1 201 Created
```
### Estrutura de diretórios
`/cmd`
Principais aplicações do projeto.

`/internal/core`
Aplicações privadas e código de biblioteca.

`/internal/model`
Interfaces de dados

`/build`
Empacotamento e integração contínua.

### Notas
A estrutura de dados escolhida para receber os dados do arquivo csv foi um slice de slice de strings, por ter sido considerado a opção mais performática em relação a opção de slice de structs.

Foi decidido não manter os dados em memória forçando a leitura do arquivo a cada operação. Isso garantiu que qualquer alteração no arquivo, como por exemplo uma atualização pela api, fosse refletida nas operações.
