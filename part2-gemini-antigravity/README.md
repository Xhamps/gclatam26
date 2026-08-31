# Prática 2: Gerador de Música com Lyria 3 e Deploy no Cloud Run

Nesta prática, vamos utilizar a CLI **Antigravity (`agy`)** como copiloto de programação para desenvolver e empacotar um microserviço HTTP em Go que gera trilhas sonoras conceituais a partir de imagens utilizando o modelo especializado **Lyria 3** (`lyria-3-clip-preview` / `lyria-3-pro-preview`).

---

## 🎯 Objetivos
1. Utilizar a Antigravity CLI (`agy`) com regras e ferramentas do ecossistema `.agents/`.
2. Consumir o modelo Lyria 3 enviando imagens como inspiração e prompts de áudio.
3. Executar o loop agentivo de validação (`smart_build`).
4. Empacotar a aplicação via Docker e implantar no Google Cloud Run.

---

## 🚀 Execução Local

### 1. Baixar Dependências
```bash
go mod tidy
```

### 2. Rodar o Servidor
```bash
export GEMINI_API_KEY="sua-chave-api"
# ou export GOOGLE_CLOUD_PROJECT="seu-projeto-gcp"
go run main.go
```

### 3. Testar a Geração com cURL
```bash
# Geração passando imagem e prompt
curl -X POST http://localhost:8080/api/compose \
  -F "prompt=Compose a boss battle 16-bit chiptune theme" \
  -F "image=@../part1-gemini-go-sdk/cute_cat.png" \
  --output cat_battle_theme.mp3
```

---

## ☁️ Deploy no Google Cloud Run

```bash
# 1. Definir o ID do projeto GCP
gcloud config set project SEU_PROJECT_ID

# 2. Build e Deploy com Cloud Build + Cloud Run
gcloud run deploy lyria-music-service \
  --source . \
  --region us-central1 \
  --allow-unauthenticated
```
