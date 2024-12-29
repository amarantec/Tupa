package cmd

import "fmt"

// Introduce exibe uma mensagem de introdução sobre o projeto web criado pelo usuário.
func introduce(projectName string) {
	fmt.Printf("Bem-vindo ao projeto web '%s'!\n", projectName)
	fmt.Println("Este projeto foi criado usando o Tupã e possui a seguinte estrutura:")
	fmt.Println()
	fmt.Println("app/")
	fmt.Println("├── cmd/                  # Comandos da API")
	fmt.Println("│   └── api/              # Ponto de entrada da aplicação")
	fmt.Println("│       └── main.go       # Arquivo principal que inicia a aplicação")
	fmt.Println("├── internal/             # Lógica interna da aplicação")
	fmt.Println("│   └── model.go          # Definição dos modelos")
	fmt.Println("├── db/                   # Lógica de banco de dados")
	fmt.Println("│   └── db.go             # Função que abre a conexão com o banco de dados")
	fmt.Println("├── handler/              # Handlers HTTP")
	fmt.Println("│   ├── routes/           # Configuração das rotas")
	fmt.Println("│   │   └── routes.go     # Definições das rotas")
	fmt.Println("│   └── middleware/       # Middleware da aplicação")
	fmt.Println("│       └── middleware.go # Funções middleware")
	fmt.Println("├── web/                  # Diretório para arquivos web")
	fmt.Println("├── config/               # Diretório para arquivos de configuração")
	fmt.Println("│   └── .env              # Coloque suas configurações para a conexão com o database")
	fmt.Println("├── deployment/           # Diretório para armazera seus Dockerfile's")
	fmt.Println("│   └── Dockerfile        # Dockerfile contendo instruções simples para contruir uma imagem do projeto")
	fmt.Println("├── templates/            # Templates HTML")
	fmt.Println("│   ├── form.html         # Template para adicionar modelos")
	fmt.Println("│   └── list.html         # Template para listar modelos")
	fmt.Println("└── css/                  # Arquivos CSS")
	fmt.Println("    └── style.css         # Estilos da aplicação")

	fmt.Println()
	fmt.Println("Execute o projeto com './build.sh'")
	fmt.Println("Ou")
	fmt.Println("Use o Dockerfile no deployment/")
}
