# Tupã - Framework para Criação de Projetos Web em Go

## Descrição

Tupã é um framework leve e eficiente projetado para criar a estrutura de projetos web utilizando a linguagem Go. Com uma arquitetura bem definida e separada em camadas, Tupã permite que desenvolvedores gerem rapidamente uma aplicação web organizada, seguindo as melhores práticas de desenvolvimento.

## Características Principais

- **Criação Rápida de Projetos**: Com um simples comando, você pode gerar toda a estrutura necessária para um novo projeto web, incluindo diretórios e arquivos essenciais.

- **Leveza**: O Tupã não depende de pacotes de terceiros, tornando-o uma escolha ideal para projetos que buscam simplicidade e eficiência.

- **Estrutura Modular**: O framework cria uma estrutura modular para o seu projeto, permitindo que cada parte da aplicação (como modelos, handlers e templates) seja facilmente gerenciada e mantida.

- **Melhores Práticas de Arquitetura**: Tupã segue as melhores práticas em arquitetura de código, garantindo que todas as camadas da aplicação estejam bem definidas e separadas. Isso facilita a manutenção e a escalabilidade do projeto.

- **Suporte a Templates HTML**: O framework inclui suporte nativo para templates HTML, permitindo a renderização dinâmica de páginas web com dados do servidor.

- **Manipulação Simples de Rotas**: Com um sistema de rotas intuitivo, Tupã permite que você defina facilmente as URLs da sua aplicação e associe-as a funções específicas para manipulação de requisições.

- **Integração com Banco de Dados**: Embora o Tupã não inclua um ORM, ele é projetado para ser facilmente integrado com qualquer solução de banco de dados, permitindo que você armazene e recupere dados rapidamente.

## Estrutura do Projeto Criado

Ao usar o Tupã para criar um novo projeto, a estrutura gerada será semelhante a esta:

app/
│
├── cmd/ # Comandos da CLI
│ ├── api/ # API do projeto
│ │ └── main.go # Ponto de entrada da aplicação
│
├── internal/ # Lógica interna da aplicação
│ ├── db/ # Lógica de banco de dados
│ │ └── db.go # Funções para manipulação do banco de dados
│ ├── handler/ # Handlers HTTP
│ │ ├── routes/ # Configuração das rotas
│ │ │ └── routes.go # Definições das rotas
│ │ └── middleware/ # Middleware da aplicação
│ │ └── middleware.go # Funções middleware
│ └── model.go # Definição dos modelos
│
├── web/ # Diretório para arquivos web
│ ├── templates/ # Templates HTML
│ │ ├── form.html # Template para adicionar modelos
│ │ └── list.html # Template para listar modelos
│ └── css/ # Arquivos CSS
│ └── style.css # Estilos da aplicação
