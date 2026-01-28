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

Crie um arquivo `.env` na raiz do projeto. Você pode usar o arquivo `.env.example` como base (se houver) ou seguir o modelo abaixo:

```env
# Server
PORT=8080

# Segurança
JWT_SECRET=sua_secret_key_super_segura

# Banco de Dados
DATABASE_URL=postgres://user:password@localhost:5432/goserver?sslmode=disable

# Variáveis opcionais se estiver usando Docker Compose para definir o container do banco:
POSTGRES_USER=postgres
POSTGRES_PASSWORD=docker
POSTGRES_DB=goserver

# Inteligência Artificial (Google AI Studio)
# Obtenha sua chave em: https://aistudio.google.com/
GEMINI_API_KEY=sua_api_key_aqui
```

---

## Como Rodar Localmente

### Pré-requisitos
*   Go instalado (1.22+)
*   Docker e Docker Compose instalados

### Passo a Passo

1.  **Configure o .env**:
    Crie o arquivo `.env` com as variáveis listadas acima.

2.  **Suba a Infraestrutura**:
    Utilize o Docker Compose para subir o banco de dados e aplicar as migrações automaticamente.
    ```bash
    docker-compose up -d
    ```

3.  **Rode a Aplicação**:
    ```bash
    go run cmd/api/main.go
    ```
    O servidor iniciará (padrão porta 8080).

4.  **(Opcional) Migrações Manuais**:
    Se precisar rodar migrações manualmente via Makefile:
    ```bash
    make migrate-up
    ```

---

## Endpoints da API

### Autenticação (Público)

| Método | Caminho | Descrição | Payload Exemplo |
| :--- | :--- | :--- | :--- |
| `POST` | `/register` | Cria novo usuário | `{"email": "user@example.com", "password": "password123"}` |
| `POST` | `/login` | Retorna JWT | `{"email": "user@example.com", "password": "password123"}` |

### Quizzes (Protegido)
*Requer header `Authorization: Bearer <seu_token>`*

| Método | Caminho | Descrição | Referência do Payload |
| :--- | :--- | :--- | :--- |
| `POST` | `/quizzes/create` | Gera um novo quiz com IA | `{"tema": "Golang Interfaces", "numQuestoes": 5, "dificuldade": "medium"}` |
| `GET` | `/quizzes` | Lista todos os quizzes do usuário | - |
| `GET` | `/quizzes/{id}` | Busca detalhes de um quiz e suas questões | - |
| `POST` | `/quizzes/submit` | Envia respostas de um quiz | `{"quiz_id": "uuid...", "answers": [{"question_id": "uuid...", "user_choice": 1}]}` |

---

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

1.  Faça um fork do projeto
2.  Crie sua feature branch (`git checkout -b feature/MinhaFeature`)
3.  Commit suas mudanças (`git commit -m 'Adiciona: MinhaFeature'`)
4.  Push para a branch (`git push origin feature/MinhaFeature`)
5.  Abra um Pull Request
