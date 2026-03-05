# Hello Monitor

Projeto simples em Go para monitorar disponibilidade de sites via terminal.

## Funcionalidades

- Menu interativo no terminal
- Monitoramento de sites com `HTTP GET`
- Registro de status em `log.txt`
- Exibição dos logs salvos

## Requisitos

- Go 1.20+ (recomendado)
- Conexão com internet para testar os sites

## Como executar

### 1. Rodar com Go

```bash
go run hello.go
```

### 2. Compilar e executar

```bash
go build -o hello hello.go
./hello
```

No Windows:

```powershell
go build -o hello.exe hello.go
.\hello.exe
```

## Como usar

Ao iniciar, o programa mostra um menu:

- `1` Inicia o monitoramento
- `2` Exibe os logs
- `0` Encerra o programa

O monitoramento atual:

- Faz `3` ciclos (`monitoramentos`)
- Aguarda `5` segundos entre ciclos (`delay`)
- Testa os sites definidos no código

## Estrutura do projeto

- `hello.go`: código-fonte principal
- `sites.txt`: lista de sites (há função para leitura, mas não está sendo usada no fluxo atual)
- `log.txt`: arquivo de logs gerado/atualizado pelo programa
- `hello.exe`: binário compilado

## Exemplo de log

```txt
04/03/2026 23:02:28 - https://google.com.br - online: true
```

## Observações

- O programa atualmente usa uma lista fixa de sites em `iniciarMonitoramento()`.
- Existe a função `leSitesDoArquivo()`, pronta para usar `sites.txt` se você quiser evoluir o projeto.
- Há mensagens com acentuação incorreta no terminal por codificação; salvar o arquivo em UTF-8 costuma resolver.
