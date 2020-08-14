# Integra Sistema
Sistema de integração entre sistemas feito em Go

![Alt text](/imagens/cadastro-integracao.png?raw=true "Página cadastro de integração")
![Alt text](/imagens/visualizar-log-retorno-integracao.png?raw=true "Página listagem log das integrações")

## Estrutura do banco desenvolvida
![Alt text](/imagens/der.png?raw=true "Diagrama de Entidade e Relacionamento")

## Tecnologias
* Go versão 1.14.4 linux/amd64
* Bootstrap v4.5.0 (https://getbootstrap.com/)
* https://icons.getbootstrap.com/

## Dependências
* github.com/joho/godotenv
* github.com/gorilla/mux
* golang.org/x/crypto/bcrypt
* github.com/go-sql-driver/mysql

Para instalar dependência basta executar o comando go get -u link-repositório, exemplo: : ```go get -u github.com/go-sql-driver/mysql```

## Configurando a aplicação 
Altere o arquivo .env na raiz do projeto conforme exemplo abaixo
```
NOME_SISTEMA=Integra Sistema
VERSAO_SISTEMA=1.0

PORTA_SERVIDOR=3003

DB_LOCALHOST=servidor_mysql_local
DB_PORTA=3306
DB_USUARIO=root
DB_SENHA=yakTLS&70c52
DB_BANCO=integra_sistema
```

## Administrador
**Usuário:** integra-sistema
**Senha:** NftK2O7y

## Execução
```go run *.go```

## Recursos

### Integração
Permite integração via webhook, onde se pode informar dados da api de destino, nome da integração, 
método de envio e nome da api no integrador que será chamado pelo sistema que irá fazer a solicitação.
![Alt text](/imagens/cadastro-integracao.png?raw=true "Página cadastro de integração")

Pode se cadastrar parametros (DE-PARA) dos dados da integração.
![Alt text](/imagens/editar-parametros.png?raw=true "Página de edição de parametros")

Pode se visualizar log das requisições
![Alt text](/imagens/visualizar-logs.png?raw=true "Página de visualização de logs")

### Usuário
Permite a criação de novos usuários para acessar o sistema, 
além de já possuir um usuário superadministrador cadastrado 
previamente no banco de dados
![Alt text](/imagens/cadastro-usuario.png?raw=true "Página de cadastro de usuário")

### Autenticação
Permite acesso ao sistema somente mediante a usuário e senha
![Alt text](/imagens/autenticacao.png?raw=true "Página de autenticação")

## Novos Recursos à implementar
* Permitir que uma integração chame outra integração já previamente cadastrada permitindo criar fluxos de trabalho
* Ao enviar dados via post, permitir o envio de lista de dados (Ex.: lista de produtos, clientes via post), no momento só dados simples são permitidos