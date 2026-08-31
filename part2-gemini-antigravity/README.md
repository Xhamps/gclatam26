# Prática 2: Desenvolvendo com Antigravity — Gerador de Música com Lyria 3

Neste laboratório prático, **você e o Antigravity (`agy`) desenvolverão o código do zero**.

O objetivo é construir um microserviço HTTP em Go que gera trilhas sonoras conceituais a partir de imagens utilizando o modelo especializado **Lyria 3** (`lyria-3-clip-preview` ou `lyria-3-pro-preview`) e implantá-lo no **Google Cloud Run**.

---

## 🎯 Especificação do Desafio

Seu microserviço em Go deve atender aos seguintes requisitos:

1. **Cliente GenAI:**
   - Inicializar o cliente oficial `google.golang.org/genai`.
   - Suportar autenticação tanto via `GEMINI_API_KEY` (Google AI Studio) quanto via `GOOGLE_CLOUD_PROJECT` (Gemini Enterprise / Vertex AI).

2. **Endpoints HTTP:**
   - `GET /healthz`: Retorna status `200 OK`.
   - `POST /api/compose`:
     - Recebe um formulário `multipart/form-data`.
     - Campo `prompt` (string): Descrição do clima musical, gênero ou emoção desejada (ex: *"16-bit retro boss battle"*).
     - Campo `image` (arquivo opcional): Imagem PNG/JPEG para servir de inspiração conceitual para o modelo.

3. **Integração com Lyria 3:**
   - Montar o payload multimodal com as partes de texto e imagem.
   - Invocar o modelo `lyria-3-clip-preview` (ou `lyria-3-pro-preview`) via `client.Models.GenerateContent`.
   - Extrair os bytes do arquivo de áudio retornado e responder com o cabeçalho apropriado (`Content-Type: audio/mpeg` ou `audio/wav`).

4. **Produção e Cloud Run:**
   - Respeitar a variável de ambiente `PORT` (padrão: 8080).
   - Encerramento gracioso (*graceful shutdown*) capturando sinais `SIGINT` e `SIGTERM`.
   - `Dockerfile` multi-stage leve para deploy no Cloud Run.

---

## 🛠️ Passo a Passo com o Antigravity CLI (`agy`)

### 1. Iniciar a sessão com o `agy`
No diretório `part2-gemini-antigravity`, abra o terminal e inicie o `agy`:

```bash
agy
```

### 2. Parear com o Agente para Escrever o Código
Envie uma instrução clara para o agente no terminal do `agy`. Por exemplo:

> *"Leia os requisitos e os comentários em `main.go` e implemente o microserviço HTTP com o SDK `google.golang.org/genai` consumindo o modelo Lyria 3 (`lyria-3-clip-preview`)."*

### 3. Executar o Loop de Validação Agentiva
Peça ao agente para verificar a compilação, formatação e tipos:

> *"Rode `go mod tidy`, valide a compilação e verifique se há problemas com `go vet`."*

*(Se tiver o `godoctor` configurado em `.agents/mcp_config.json`, você pode pedir: `"Execute um smart_build com godoctor e corrija quaisquer apontamentos."`)*

### 4. Testar a Aplicação Localmente
Em outro terminal:

```bash
# Inicie o servidor
go run main.go

# Teste com cURL passando a imagem do gato mago gerada na Prática 1:
curl -X POST http://localhost:8080/api/compose \
  -F "prompt=Compose an epic orchestral wizard theme" \
  -F "image=@../part1-gemini-go-sdk/cute_cat.png" \
  --output cat_soundtrack.mp3
```

### 5. Criar o Dockerfile e Deploy no Cloud Run
Peça ao `agy` para criar o Dockerfile e o comando de deploy:

> *"Crie um Dockerfile multi-stage com Go 1.24 e alpine para deploy no Google Cloud Run."*

Em seguida, faça o deploy no GCP:

```bash
gcloud run deploy lyria-music-service \
  --source . \
  --region us-central1 \
  --allow-unauthenticated
```
