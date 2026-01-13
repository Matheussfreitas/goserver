# 📚 RespondAI

**RespondAI** é uma API robusta e escalável desenvolvida em **Go**, projetada para auxiliar estudantes na fixação de conteúdo. Utilizando a inteligência artificial do **Google Gemini**, a aplicação transforma textos de estudo em questionários interativos e personalizados.

Este projeto segue os princípios da **Clean Architecture** e adota o **Standard Go Project Layout**, garantindo um código desacoplado, testável e de fácil manutenção.

---

## Funcionalidades

### Inteligência Artificial
*   **Geração de Quizzes**: Envie qualquer texto ou resumo e receba perguntas de múltipla escolha geradas por IA.
*   **Dificuldade Adaptável**: Configure o nível das questões entre *Fácil*, *Médio* e *Difícil*.
*   **Feedback Detalhado**: Explicações geradas pela IA para correções de respostas.

### Autenticação & Segurança
*   **Cadastro e Login**: Sistema completo de usuários.
*   **JWT (JSON Web Tokens)**: Proteção de rotas e identificação de usuários sem estado (stateless).
*   **Bcrypt**: Hashing seguro de senhas antes da persistência.

### Engenharia de Software
*   **Clean Architecture**: Separação clara entre Domínio, Casos de Uso (Service), Repositórios e Interface (Handlers).
*   **Injeção de Dependências**: Facilita testes e troca de implementações.
*   **Mux Padrão Moderno**: Utilização do roteador `http.ServeMux` do Go 1.22+.

---

## Arquitetura do Projeto

A estrutura de pastas reflete a separação de responsabilidades:

```
.
├── cmd/api/            # Ponto de entrada da aplicação (main.go)
├── internal/
│   ├── config/         # Carregamento de env vars e configurações
│   ├── database/       # Conexão com banco de dados (PostgreSQL)
│   ├── domain/         # Entidades e interfaces de negócio (Core)
│   ├── handler/        # Controladores HTTP (Parse de JSON, validação)
│   ├── middleware/     # Interceptadores (Auth, Logger)
│   ├── repository/     # Implementação do acesso a dados (SQL/pgx)
│   ├── service/        # Regras de negócio e orquestração
│   └── utils/          # Funções auxiliares (JWT, Parsers)
├── migrations/         # Scripts de migração de banco de dados
└── .env                # Variáveis de ambiente (não versionado)
```

---

## Tecnologias

*   **Linguagem**: Go (1.22+)
*   **Banco de Dados**: PostgreSQL
*   **Driver SQL**: pgx/v5
*   **AI SDK**: Google GenAI SDK (Gemini)
*   **Autenticação**: Golang-JWT
*   **Server**: `net/http` (Standard Lib)

---

## Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto seguindo o modelo abaixo:

```env
# Servidor
SERVER_PORT=8080

# Banco de Dados
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=goserver
# Ou se preferir usar url de conexão direta nos drivers que suportam:
# DB_URL=postgres://user:pass@localhost:5432/goserver

# Segurança
JWT_SECRET=sua_hash_secreta_super_segura

# Inteligência Artificial (Google AI Studio)
GEMINI_API_KEY=sua_api_key_do_google_gemini
```

---

## Como Rodar Localmente

### Pré-requisitos
*   Go instalado
*   PostgreSQL rodando e com o banco criado

### Passos

1.  **Clone o repositório**:
    ```bash
    git clone https://github.com/Matheussfreitas/goserver.git
    cd goserver
    ```

2.  **Instale as dependências**:
    ```bash
    go mod tidy
    ```

3.  **Configure o ambiente**:
    *   Copie o exemplo acima para um arquivo `.env`.
    *   Certifique-se que o banco de dados existe.

4.  **Execute a aplicação**:
    ```bash
    go run cmd/api/main.go
    ```

O servidor estará rodando em `http://localhost:8080`.

---

## Endpoints da API

### Autenticação (Público)

| Método | Caminho | Descrição | Payload Exemplo |
| :--- | :--- | :--- | :--- |
| `POST` | `/register` | Cria novo usuário | `{"email": "...", "password": "..."}` |
| `POST` | `/login` | Retorna JWT | `{"email": "...", "password": "..."}` |

### Quizzes (Protegido)
*Requer header `Authorization: Bearer <seu_token>`*

| Método | Caminho | Descrição |
| :--- | :--- | :--- |
| `POST` | `/quizzes/generate` | Gera um novo quiz. Payload: `{"content": "...", "difficulty": "Medium", "questions_count": 5}` |
| `GET` | `/quizzes` | Lista quizzes do usuário logado |
| `GET` | `/quizzes/{id}` | Detalhes de um quiz específico |
| `POST` | `/quizzes/{id}/submit` | Envia respostas para correção |

---

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

1.  Faça um fork do projeto
2.  Crie sua feature branch (`git checkout -b feature/MinhaFeature`)
3.  Commit suas mudanças (`git commit -m 'Adiciona: MinhaFeature'`)
4.  Push para a branch (`git push origin feature/MinhaFeature`)
5.  Abra um Pull Request
