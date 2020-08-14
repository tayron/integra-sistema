# Integra Sistema
Sistema de integração entre sistemas feito em Go

![Alt text](/imagens/cadastro-integracao.png?raw=true "Página cadastro de integração")

## Estrutura do banco desenvolvida
![Alt text](/imagens/der.png?raw=true "Diagrama de Entidade e Relacionamento")

## Tecnologias
* Go versão 1.14.4 linux/amd64
* https://github.com/ColorlibHQ/AdminLTE/releases/tag/v3.0.5
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
Permite integração via webhook, um sistema A faz a solicitação para uma API via POST no integrador, 
onde os parametros do sistema A são repassados para API do sistema B que por sua vez devolve o 
retorno da solicitação para o integrador que o integra ao sistema A.
![Alt text](/imagens/cadastro-integracao.png?raw=true "Página cadastro de integração")

Permite realizar DE-PARA das informações do sistema A para o sistema B na tela de cadastro de parâmetros
![Alt text](/imagens/editar-parametros.png?raw=true "Página de edição de parametros")

Permite visualizar log das requisições, parametros enviados pelo sistema A e o retorno do sistema B
![Alt text](/imagens/visualizar-logs.png?raw=true "Página de visualização de logs")

Visualização do parâmetro enviado pelo sistema A
![Alt text](/imagens/visualizar-log-parametro-integracao.png?raw=true "Parametro enviado pelo sistema A")

Visualização da reposta enviada pelo sistema B 
![Alt text](/imagens/visualizar-log-retorno-integracao.png?raw=true "Retorno enviado pelo sistema B")

### Usuário
Permite a criação de novos usuários para acessar o sistema, além de já possuir um usuário 
superadministrador cadastrado que não pode ser alterado ou excluído via sistema.
![Alt text](/imagens/cadastro-usuario.png?raw=true "Página de cadastro de usuário")

### Autenticação
Permite acesso ao sistema somente mediante a usuário e senha
![Alt text](/imagens/autenticacao.png?raw=true "Página de autenticação")

## Novos Recursos à implementar
### Fluxo de integração
* Permitir criar uma integração chame outra integração já existente no sistema, bastando informar nome da API
### Lista de dados como parametros
* Ao enviar dados via post, permitir o envio de lista de dados (Ex.: lista de produtos, clientes via post), no momento só dados simples são permitidos