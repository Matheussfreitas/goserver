# Go Authentication Server

Um servidor HTTP simples construído em Go pura (apenas standard library) para demonstrar uma implementação limpa de autenticação.

## Sobre o Projeto

Este projeto tem como objetivo servir de base para estudos sobre desenvolvimento web com Go, focando em:
*   **API RESTful**: Endpoints que aceitam e retornam JSON.
*   **Arquitetura em Camadas**: Separação clara entre `Controller` (HTTP) e `Service` (Regras de Negócio).
*   **Tratamento de Erros**: Uso de *Sentinel Errors* para mapear erros de negócio para Status Codes HTTP corretos (401, 409, 500).
*   **Roteamento Moderno**: Uso do `http.ServeMux` com a sintaxe de métodos (ex: `POST /path`) disponível nas versões mais recentes do Go.

## Funcionalidades

*   **Cadastro (`POST /register`)**: Criação de novos usuários.
*   **Login (`POST /login`)**: Autenticação de usuários existentes.
*   *Nota: Atualmente utiliza armazenamento em memória (volátil).*

## Como Rodar

1.  Certifique-se de ter o [Go instalado](https://go.dev/dl/).
2.  Clone o repositório.
3.  Execute o servidor:

```bash
go run .
```

O servidor iniciará na porta `8080`.

## Testando a API

Exemplo de requisição para Cadastro:
```bash
curl -X POST http://localhost:8080/register \
   -H "Content-Type: application/json" \
   -d '{"email":"teste@exemplo.com", "password":"123"}'
```

Exemplo de requisição para Login:
```bash
curl -X POST http://localhost:8080/login \
   -H "Content-Type: application/json" \
   -d '{"email":"teste@exemplo.com", "password":"123"}'
```
