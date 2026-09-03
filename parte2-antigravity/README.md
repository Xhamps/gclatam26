# Prática 2: Desenvolvendo com Antigravity — Gerador de Música com Lyria 3

Neste laboratório prático, **você e o Antigravity (`agy`) desenvolverão o código do zero**.

O objetivo é construir um microserviço HTTP em Go que gera trilhas sonoras conceituais a partir de imagens utilizando o modelo especializado **Lyria 3** (`lyria-3-clip-preview` ou `lyria-3-pro-preview`) e implantá-lo no **Google Cloud Run**.

## Pré-requisitos

Instalar Antigravity 2.0 e/ou Antigravity CLI baixando do site oficial https://antigravity.google.

Um projeto GCP com a API AI Platform (aiplatform.googleapis.com) habilitada:
```sh
gcloud services enable aiplatform.googleapis.com
```

## Especificação

Seu microserviço em Go deve atender aos seguintes requisitos:

1. **Cliente GenAI:**
   - Inicializar o cliente oficial `google.golang.org/genai`.
   - Suportar autenticação via `GOOGLE_CLOUD_PROJECT` (Gemini Enterprise / Vertex AI).

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