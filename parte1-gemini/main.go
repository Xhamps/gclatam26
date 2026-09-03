package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func main() {
	var prompt string
	var output string
	flag.StringVar(&prompt, "p", "Generate a high-resolution, cute image of a fluffy cat wearing a tiny wizard hat.", "Prompt para geracao da imagem")
	flag.StringVar(&output, "o", "output.png", "Caminho do arquivo de saida")
	flag.Parse()

	ctx := context.Background()

	// Configuração flexível: suporta Gemini Enterprise (Vertex AI) e Google AI Studio
	var clientConfig *genai.ClientConfig

	if project := os.Getenv("GOOGLE_CLOUD_PROJECT"); project != "" {
		clientConfig = &genai.ClientConfig{
			Project:  project,
			Location: "global", // Modelos Nano Banana são servidos globalmente
			Backend:  genai.BackendEnterprise,
		}
	} else {
		log.Fatalln("Para executar esse programa configure o seu GOOGLE_CLOUD_PROJECT de acordo com a sua conta GCP.")
	}

	client, err := genai.NewClient(ctx, clientConfig)
	if err != nil {
		log.Fatalf("falha ao criar o cliente GenAI: %v", err)
	}

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
				if err := os.WriteFile(output, part.InlineData.Data, 0644); err != nil {
					log.Fatalf("falha ao salvar imagem: %v", err)
				}
				fmt.Printf("Sucesso! Imagem salva em: %s\n", output)
				return
			}
		}
	}

	log.Fatal("nenhum dado de imagem retornado na resposta")
}
