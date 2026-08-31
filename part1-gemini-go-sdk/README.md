# Prática 1: Hello World com Go GenAI SDK (Cat Picture Generator)

Nesta prática, vamos conectar nossa aplicação Go à família Gemini utilizando o **Go GenAI SDK oficial** (`google.golang.org/genai`) e gerar uma imagem de um gato mago usando o modelo multimodal **Nano Banana 2 Lite** (`gemini-3.1-flash-lite-image`).

---

## 🎯 Objetivos
1. Inicializar o cliente unificado do Go GenAI SDK.
2. Fazer uma chamada multimodal usando `client.Models.GenerateContent`.
3. Extrair os bytes brutos (`InlineData`) da resposta e salvá-los localmente como arquivo de imagem (`cute_cat.png`).

---

## 🚀 Como Executar

### 1. Baixar Dependências
```bash
go mod tidy
```

### 2. Autenticação

**Opção 1: Google AI Studio (API Key)**
```bash
export GEMINI_API_KEY="sua-chave-api"
```

**Opção 2: Gemini Enterprise / Google Cloud (Vertex AI)**
```bash
export GOOGLE_CLOUD_PROJECT="seu-projeto-gcp"
```

### 3. Rodar o Programa
```bash
go run main.go
```

Após a execução, abra o arquivo `cute_cat.png` gerado no diretório!
