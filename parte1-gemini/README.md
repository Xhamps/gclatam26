# Pratica 1: Hello World com Go GenAI SDK

Gerador de imagem com Nano Banana 2 Lite (`gemini-3.1-flash-lite-image`) utilizando `google.golang.org/genai`.

## Objetivos

* Inicializar o cliente do Go GenAI SDK.
* Realizar chamada de geracao de conteudo com modelo de imagem.

## Execucao

```bash
# Dependencias
go mod tidy

# Autenticacao
export GOOGLE_CLOUD_PROJECT="seu-projeto-gcp"

# Rodar com prompt padrao (salva em output.png)
go run main.go

# Rodar com prompt e arquivo de saida customizados
go run main.go -p "A capybara wearing sunglasses coding in Go" -o capybara.png
```

