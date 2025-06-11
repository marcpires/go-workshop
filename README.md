# Go Workshop: Leitor de feed
Este repositório contêm os códigos referentes aos encontros do workshop de Golang no LHC

## Requisitos
- Git
- Git lfs
- Go >= 1.23
- VSCode com [extensão Go](https://marketplace.visualstudio.com/items?itemName=golang.Go) ou a IDE de sua preferência

## Estrutura do projeto

```sh
├── cmd
│   ├── main.go -- Ponto de entrada da aplicação
│   └── shadow
│       └── client.go --- Exemplo de mascaramento de variáveis
├── data
│   └── data.json -- Listya de feeds
├── go.mod
├── go.sum
├── internal
   ├── matchers
   │   └── rss.go -- Matcher para bus ca de feeds rss
   └── search
       ├── default.go -- Matcher padrão
       ├── feed.go -- Suporte para leitura de arquivos json
       ├── match.go -- Interface para suportar diferentes tipos de MAtcher
       └── search.go -- Lógica principal para realizar buscas
```

### Organização dos ramos:
Cada exemplo será organizado na ramificação correspondente ao tema apresentado nos encontros:

exemplo:
```sh
feat/01/project-code-organization
```

**Nota**: Nome das ramificações e comentários no código serão apresentados em Inglês

## Architetura

![Feed Reader Architecture](./assets/architecture.png)

