# Go Authentication Server

Este é um projeto robusto de servidor HTTP em Go, reestruturado seguindo o padrão **Standard Go Project Layout**. Ele implementa um sistema completo de autenticação e gerenciamento de usuários utilizando **Clean Architecture** para garantir escalabilidade, testabilidade e manutenibilidade.

## 🏗 Arquitetura

O projeto está organizado para separar responsabilidades de forma clara:

*   **`cmd/api`**: Ponto de entrada da aplicação (`main.go`).
*   **`internal/config`**: Gerenciamento de variáves de ambiente e configurações.
*   **`internal/database`**: Configuração e conexão com o banco de dados (PostgreSQL).
*   **`internal/domain`**: Definições das entidades principais do sistema (Structs).
*   **`internal/repository`**: Camada de acesso a dados (SQL queries usando `pgx`).
*   **`internal/service`**: Regras de negócio da aplicação (Hashing de senha, validações).
*   **`internal/handler`**: Camada de transporte HTTP (Controllers, Roteamento, Parse de JSON).
*   **`internal/middleware`**: Interceptadores de requisições (ex: Proteção de rotas com JWT).
*   **`internal/utils`**: Funções utilitárias (ex: Geração e validação de JWT).

## 🚀 Tecnologias Utilizadas

*   **Go 1.22+**: Linguagem principal.
*   **PostgreSQL**: Banco de dados relacional.
*   **pgx/v5**: Driver de alta performance para Postgres.
*   **Golang-JWT**: Geração e validação de tokens JWT.
*   **Bcrypt**: Hashing seguro de senhas.
*   **net/http**: Servidor HTTP padrão do Go (com `ServeMux` moderno).

## ⚙️ Configuração

Antes de rodar, certifique-se de configurar o arquivo `.env` na raiz do projeto com as seguintes chaves:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=goserver
JWT_SECRET=sua_chave_secreta_super_segura
SERVER_PORT=8080
```

## 🏃 Como Rodar

1.  **Clone o repositório** e entre na pasta.
2.  **Baixe as dependências**:
    ```bash
    go mod tidy
    ```
3.  **Execute a aplicação**:
    ```bash
    go run cmd/api/main.go
    ```

O servidor iniciará na porta definida no `.env` (ex: `8080`).

## 📡 Endpoints

### Público

*   **`POST /register`**: Criação de novos usuários.
    *   Body: `{"email": "...", "password": "..."}`
*   **`POST /login`**: Autenticação de usuários. Retorna um Token JWT.
    *   Body: `{"email": "...", "password": "..."}`

### Protegido

*   As rotas protegidas exigem o header: `Authorization: Bearer <TOKEN>`
