package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/genai"
)

type MusicService struct {
	client *genai.Client
	model  string
}

func newMusicService(ctx context.Context) (*MusicService, error) {
	var clientConfig *genai.ClientConfig

	if project := os.Getenv("GOOGLE_CLOUD_PROJECT"); project != "" {
		clientConfig = &genai.ClientConfig{
			Project:  project,
			Location: "global",
			Backend:  genai.BackendVertexAI,
		}
	} else if apiKey := os.Getenv("GEMINI_API_KEY"); apiKey != "" {
		clientConfig = &genai.ClientConfig{
			APIKey:  apiKey,
			Backend: genai.BackendGeminiAPI,
		}
	}

	client, err := genai.NewClient(ctx, clientConfig)
	if err != nil {
		return nil, fmt.Errorf("falha ao inicializar o cliente GenAI: %w", err)
	}

	model := os.Getenv("LYRIA_MODEL")
	if model == "" {
		model = "lyria-3-clip-preview" // Modelo padrão para vinhetas de áudio
	}

	return &MusicService{
		client: client,
		model:  model,
	}, nil
}

func (s *MusicService) handleCompose(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Limite de 20MB para uploads
	if err := r.ParseMultipartForm(20 << 20); err != nil {
		http.Error(w, "falha ao processar formulário multipart", http.StatusBadRequest)
		return
	}

	prompt := r.FormValue("prompt")
	if prompt == "" {
		prompt = "Compose an upbeat, energetic soundtrack inspired by this image."
	}

	var parts []*genai.Part
	parts = append(parts, genai.NewPartFromText(prompt))

	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		imageData, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "falha ao ler arquivo de imagem", http.StatusInternalServerError)
			return
		}

		mimeType := header.Header.Get("Content-Type")
		if mimeType == "" {
			mimeType = "image/png"
		}

		parts = append(parts, &genai.Part{
			InlineData: &genai.Blob{
				MIMEType: mimeType,
				Data:     imageData,
			},
		})
	}

	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	defer cancel()

	log.Printf("Invocando modelo Lyria 3 (%s) com prompt: %q", s.model, prompt)

	contents := []*genai.Content{
		{
			Role:  "user",
			Parts: parts,
		},
	}

	resp, err := s.client.Models.GenerateContent(ctx, s.model, contents, nil)
	if err != nil {
		log.Printf("erro na geração musical com Lyria 3: %v", err)
		http.Error(w, fmt.Sprintf("falha na geração musical: %v", err), http.StatusInternalServerError)
		return
	}

	// Procura pelos bytes de áudio na resposta
	for _, cand := range resp.Candidates {
		if cand.Content == nil {
			continue
		}
		for _, part := range cand.Content.Parts {
			if part.InlineData != nil && len(part.InlineData.Data) > 0 {
				w.Header().Set("Content-Type", part.InlineData.MIMEType)
				w.Header().Set("Content-Disposition", "attachment; filename=\"soundtrack.mp3\"")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(part.InlineData.Data)
				return
			}
		}
	}

	http.Error(w, "nenhum dado de áudio gerado pelo modelo", http.StatusInternalServerError)
}

func (s *MusicService) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func main() {
	ctx := context.Background()

	service, err := newMusicService(ctx)
	if err != nil {
		log.Fatalf("erro fatal na inicialização: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", service.handleHealth)
	mux.HandleFunc("/api/compose", service.handleCompose)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	serverCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("Serviço de Música Lyria 3 rodando na porta :%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("falha no servidor HTTP: %v", err)
		}
	}()

	<-serverCtx.Done()
	log.Println("Encerrando servidor de forma graciosa...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("encerramento forçado do servidor: %v", err)
	}
	log.Println("Servidor finalizado com sucesso.")
}
