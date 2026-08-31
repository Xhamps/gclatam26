package main

import (
	"fmt"
	"net/http"
	"os"
)

// TODO: Prática 2 — Gerador de Música com Lyria 3 e Deploy no Cloud Run
//
// Seu objetivo nesta prática é desenvolver um microserviço HTTP em Go utilizando
// a CLI do Antigravity (`agy`) em modo de pareamento.
//
// Requisitos do Serviço:
// 1. Inicializar o cliente oficial do Go GenAI SDK (`google.golang.org/genai`)
//    suportando Gemini Enterprise (Vertex AI via GOOGLE_CLOUD_PROJECT) e
//    Google AI Studio (via GEMINI_API_KEY).
// 2. Criar um endpoint `GET /healthz` que retorne 200 OK.
// 3. Criar um endpoint `POST /api/compose` que receba:
//    - Um campo de formulário multipart `prompt` (descrição textual do estilo/clima da música)
//    - Um arquivo de imagem multipart opcional `image` (imagem de inspiração conceitual)
// 4. Invocar o modelo Lyria 3 (`lyria-3-clip-preview` ou `lyria-3-pro-preview`)
//    passando o prompt e a imagem como partes multimodais.
// 5. Extrair os bytes de áudio gerados e retornar como resposta HTTP (Content-Type: audio/mpeg ou audio/wav).
// 6. Suportar a porta configurada pela variável de ambiente PORT (padrão: 8080)
//    e encerramento gracioso (graceful shutdown com context/signal).
//
// Dica: Use o `agy` para implementar o serviço iterativamente!
// Exemplo de prompt para o agy:
//   "Implemente o servidor HTTP em main.go de acordo com os requisitos descritos nos comentários do arquivo."

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Servidor de partida da Prática 2 ouvindo em :%s\n", port)
	fmt.Println("Abra o terminal com 'agy' e implemente o gerador de música com Lyria 3!")

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK (Starter)"))
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Erro no servidor: %v\n", err)
	}
}
