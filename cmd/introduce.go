package cmd

import "fmt"

// Introduce exibe uma mensagem de introdução sobre o projeto web criado pelo usuário.
func introduce(projectName string) {
    fmt.Printf("Bem-vindo ao projeto web '%s'!\n", projectName)
    fmt.Println("Este projeto foi criado usando o Tupã e possui a seguinte estrutura:")
    fmt.Println()
    fmt.Println("app/")
    fmt.Println("├── cmd/                # Comandos da API")
    fmt.Println("│   └── api/           # Ponto de entrada da aplicação")
    fmt.Println("│       └── main.go    # Arquivo principal que inicia a aplicação")
    fmt.Println("├── internal/           # Lógica interna da aplicação")
    fmt.Println("│   └── model.go        # Definição dos modelos")
    fmt.Println("├── db/                 # Lógica de banco de dados")
    fmt.Println("│   └── db.go           # Funções para manipulação do banco de dados")
    fmt.Println("├── handler/            # Handlers HTTP")
    fmt.Println("│   ├── routes/         # Configuração das rotas")
    fmt.Println("│   │   └── routes.go   # Definições das rotas")
    fmt.Println("│   └── middleware/      # Middleware da aplicação")
    fmt.Println("│       └── middleware.go# Funções middleware")
    fmt.Println("├── web/                # Diretório para arquivos web")
    fmt.Println("├── templates/          # Templates HTML")
    fmt.Println("│   ├── form.html       # Template para adicionar modelos")
    fmt.Println("│   └── list.html       # Template para listar modelos")
    fmt.Println("└── css/                # Arquivos CSS")
    fmt.Println("    └── style.css       # Estilos da aplicação")
    
    fmt.Println()
    fmt.Println("Para começar a desenvolver, execute o comando 'go mod tidy' para baixar as dependências necessárias.")
    fmt.Println("Depois, você pode iniciar a aplicação com 'go run cmd/api/main.go'.")
}

