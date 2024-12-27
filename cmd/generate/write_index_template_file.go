package generate

import "os"

func WriteIndexTemplateFile(indexTemplateFilePath string) error {
	/*
		Content to write in index.html
	*/

	packageContent := []byte(`<!DOCTYPE html>
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tupã - Framework para Criação de Projetos Web em Go</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css">
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }
        header {
            background: #007bff;
            color: white;
            padding: 20px 0;
            text-align: center;
        }
        h1 {
            margin: 0;
            font-size: 2.5em;
        }
        .container {
            width: 80%;
            margin: auto;
            overflow: hidden;
        }
        section {
            background: white;
            padding: 20px;
            margin-top: 20px;
            border-radius: 5px;
            box-shadow: 0 2px 5px rgba(0,0,0,0.1);
        }
        h2 {
            color: #007bff;
        }
        ul {
            list-style-type: none;
            padding-left: 0;
        }
        li {
            background: #e9ecef;
            margin: 10px 0;
            padding: 10px;
            border-radius: 5px;
        }
        footer {
            text-align:center; 
            margin-top:20px; 
            padding:20px; 
            background:#007bff; 
            color:white; 
            position:relative; 
        }
    </style>
</head>
<body>
    <header>
        <h1>Tupã - Framework para Criação de Projetos Web em Go</h1>
    </header>
    <div class="container">
        <section>
            <h2>Descrição</h2>
            <p>Tupã é um framework leve e eficiente projetado para criar a estrutura de projetos web utilizando a linguagem Go. Com uma arquitetura bem definida e separada em camadas, Tupã permite que desenvolvedores gerem rapidamente uma aplicação web organizada, seguindo as melhores práticas de desenvolvimento.</p>
        </section>
        
        <section>
            <h2>Características Principais</h2>
            <ul>
                <li><strong>Criação Rápida de Projetos:</strong> Com um simples comando, você pode gerar toda a estrutura necessária para um novo projeto web, incluindo diretórios e arquivos essenciais.</li>
                <li><strong>Leveza:</strong> O Tupã não depende de pacotes de terceiros, tornando-o uma escolha ideal para projetos que buscam simplicidade e eficiência.</li>
                <li><strong>Estrutura Modular:</strong> O framework cria uma estrutura modular para o seu projeto, permitindo que cada parte da aplicação (como modelos, handlers e templates) seja facilmente gerenciada e mantida.</li>
                <li><strong>Melhores Práticas de Arquitetura:</strong> Tupã segue as melhores práticas em arquitetura de código, garantindo que todas as camadas da aplicação estejam bem definidas e separadas. Isso facilita a manutenção e a escalabilidade do projeto.</li>
                <li><strong>Suporte a Templates HTML:</strong> O framework inclui suporte nativo para templates HTML, permitindo a renderização dinâmica de páginas web com dados do servidor.</li>
                <li><strong>Manipulação Simples de Rotas:</strong> Com um sistema de rotas intuitivo, Tupã permite que você defina facilmente as URLs da sua aplicação e associe-as a funções específicas para manipulação de requisições.</li>
                <li><strong>Integração com Banco de Dados:</strong> Embora o Tupã não inclua um ORM, ele é projetado para ser facilmente integrado com qualquer solução de banco de dados, permitindo que você armazene e recupere dados rapidamente.</li>
            </ul>
        </section>

        <footer>
            <p>&copy; 2024 Tupã Framework. Todos os direitos reservados.</p>
        </footer>
    </div>
</body>
</html>
`)

	indexTemplateFile, err := os.OpenFile(indexTemplateFilePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer indexTemplateFile.Close()

	_, err = indexTemplateFile.Write(packageContent)
	if err != nil {
		return err
	}
	return nil
}
