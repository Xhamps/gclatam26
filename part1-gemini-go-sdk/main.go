package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	// Configuração flexível: suporta Gemini Enterprise (Vertex AI) e Google AI Studio
	var clientConfig *genai.ClientConfig

	if project := os.Getenv("GOOGLE_CLOUD_PROJECT"); project != "" {
		clientConfig = &genai.ClientConfig{
			Project:  project,
			Location: "global", // Modelos Nano Banana são servidos globalmente
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
		log.Fatalf("falha ao criar o cliente GenAI: %v", err)
	}

	prompt := "Generate a high-resolution, cute image of a fluffy cat wearing a tiny wizard hat."

	fmt.Println("Enviando requisição para Nano Banana 2 Lite (gemini-3.1-flash-lite-image)...")

	// Chamada ao modelo Nano Banana 2 Lite
	resp, err := client.Models.GenerateContent(
		ctx,
		"gemini-3.1-flash-lite-image",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Fatalf("falha ao gerar imagem: %v", err)
	}

	// Extração dos bytes da imagem retornados pelo modelo
	for _, candidate := range resp.Candidates {
		if candidate.Content == nil {
			continue
		}
		for _, part := range candidate.Content.Parts {
			if part.InlineData != nil && part.InlineData.Data != nil {
				filename := "cute_cat.png"
				if err := os.WriteFile(filename, part.InlineData.Data, 0644); err != nil {
					log.Fatalf("falha ao salvar imagem: %v", err)
				}
				fmt.Printf("Sucesso! Imagem do gato mago salva em: %s\n", filename)
				return
			}
		}
	}

	log.Fatal("nenhum dado de imagem retornado na resposta")
}
