Entendi! Aqui está um README para sua resposta ao desafio:

---

# Desafio: Cliente e Servidor em Go para Consulta de Cotação do Dólar

Este é o resultado do desafio proposto para construir dois sistemas em Go: `client.go` e `server.go`, para consulta de cotação do dólar.

## Descrição do Desafio

O desafio consiste em criar um cliente e um servidor em Go, onde o cliente faz uma solicitação HTTP para o servidor, que por sua vez consome uma API de câmbio Dólar-Real e retorna o valor da cotação atual apenas no formato JSON para o cliente. O servidor deve registrar cada cotação recebida em um banco de dados SQLite, com timeouts definidos para chamadas de API e operações de banco de dados.

### Requisitos

- **client.go**:
  - Realiza uma solicitação HTTP para o servidor, solicitando a cotação do dólar.
  - Define um timeout máximo de 300ms para receber a resposta do servidor.
  - Salva a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}.

- **server.go**:
  - Consome a API de câmbio Dólar-Real disponível em https://economia.awesomeapi.com.br/json/last/USD-BRL.
  - Retorna a cotação atual no formato JSON para o cliente, incluindo apenas o campo "bid" do JSON recebido.
  - Utiliza o endpoint `/cotacao` e a porta 8080 para o servidor HTTP.
  - Registra cada cotação recebida no banco de dados SQLite.
  - Define timeouts máximos de 200ms para chamadas de API e 10ms para operações de banco de dados.

## Executando a Solução

1. Certifique-se de ter Go instalado em sua máquina.
2. Clone este repositório: `git clone https://github.com/guirialli/Client-Server-API-Resposta/`
3. Navegue até o diretório do projeto
4. Execute `go run main.go` para iniciar a solução.
   
## Organização do Proejeto

- **client**: Contém o código relacionado ao cliente que faz a solicitação HTTP para o servidor. O arquivo `cliente.go` contém a lógica para fazer essa solicitação.

- **controllers**: Este diretório provavelmente contém controladores responsáveis por manipular as solicitações HTTP recebidas pelo servidor. O arquivo `contabil_controllers.go` contem controladores específicos para operações relacionadas à contabilidade.

- **dao**: Contém os objetos de acesso a dados (DAO, Data Access Objects) responsáveis por interagir com o banco de dados. O arquivo `contabil_dao.go`  contém informações para realizar a pesistência no banco sqlite.

- **database**: Pode conter arquivos relacionados à inicialização e configuração do banco de dados. O arquivo `init.go` contém código para inicializar e realizar a migração do banco dados; também sendo resposavel por guardar a conexão.

- **models**: Contém as definições de modelos de dados usados no projeto. Por exemplo, o arquivo `contabil.go` define a estrutura de dados para registros contábeis e o arquivo `usdbrl.go` define a estrutura de dados para a cotação do dólar em relação ao real.

- **routers**: Este diretório provavelmente contém arquivos relacionados à definição de rotas para o servidor HTTP. O arquivo `contabil_routers.go` específica para operações relacionadas à contabilidade.

- **server**: Contém o código relacionado ao servidor HTTP. O arquivo `server.go` provavelmente contém a lógica principal do servidor, incluindo a configuração do roteamento e a inicialização do servidor.

- **services**: Comtém arquivos relacionados à lógica de negócios ou serviços utilizados no projeto. O arquivo `economia_awesomeapi.go` comtém a lógica para consumir a API de câmbio Dólar-Real.

- **contabil.db**: Este é o banco de dados SQLite utilizado pelo projeto para armazenar os registros contábeis.

- **cotacao.txt**: Este arquivo é usado para armazenar a cotação atual do dólar em formato de texto.

- **go.mod** e **go.sum**: Estes arquivos são usados pelo Go Modules para gerenciar as dependências do projeto.

- **main.go**: É o arquivo de entrada para o programa, onde a execução principal é iniciada relacionado ao desafio.



## Notas Adicionais

- Certifique-se de fornecer o link do seu repositório ao finalizar o desafio para a correção.

---
