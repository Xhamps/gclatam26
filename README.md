# Gemini para Desenvolvedores Go 🐾 🚀
### Workshop Oficial — GopherCon LATAM 2026
**Data:** Quinta-feira, 3 de Setembro de 2026  
**Duração:** 3h (3 blocos de 50 min com intervalos)

---

## 🎯 Sobre o Workshop

Este repositório contém todo o material de apoio, códigos de exemplo, configurações agentivas e laboratórios práticos do workshop **"Gemini para Desenvolvedores Go"** apresentado no **GopherCon LATAM 2026**.

O objetivo do workshop é capacitar desenvolvedores Go a construir aplicações nativas de IA e utilizar ferramentas agentivas modernas com a família de modelos Gemini e o ecossistema Antigravity.

---

## 📅 Grade e Cronograma

| Horário / Duração | Bloco | Tema | Prática |
| :--- | :--- | :--- | :--- |
| **50 min** | **Parte 1** | **A Família de Modelos Gemini**<br>• Nomenclatura e catálogo (Gemini 3.7 Flash, Nano Banana 2 Lite, Veo, Lyria, Gemma)<br>• Features: Thinking Mode, Function Calling (4 etapas), Grounding, Structured Outputs<br>• Superfícies de API (GenerateContent, Interactions, Live, Batch, Managed Agents)<br>• Autenticação: Google AI Studio vs. Gemini Enterprise | **Prática 1:** Hello World & Gerador de fotos de gatinhos mágicos com Go GenAI SDK e Nano Banana 2 Lite (`gemini-3.1-flash-lite-image`) |
| **10 min** | *Intervalo* | *Coffee Break & Networking* | — |
| **50 min** | **Parte 2** | **Desenvolvimento com Gemini e Antigravity**<br>• Go na Era da IA: Afinidade Agentiva e o Loop Agentivo Determinístico<br>• O ecossistema Antigravity: Antigravity 2.0 (Agent Manager), `agy` CLI (em Go), Antigravity IDE<br>• Customizações no workspace (`.agents/`): Regras, MCPs (`gopls`, `godoctor`), Skills, Hooks e Subagentes | **Prática 2:** Desenvolvimento assistido com `agy` de um gerador de trilhas sonoras com Lyria 3 a partir de imagens + deploy no Cloud Run |
| **10 min** | *Intervalo* | *Coffee Break* | — |
| **50 min** | **Parte 3** | *(Em aberto / TBD)* | **Prática 3:** *(Em aberto / TBD)* |

---

## 📂 Estrutura do Repositório

```text
.
├── README.md                      # Este guia
├── setup/                         # Instruções de setup de ambiente e credenciais
│   └── README.md
├── .agents/                       # Customizações do workspace Antigravity
│   ├── mcp_config.json            # Servidores MCP recomendados (gopls, godoctor, etc.)
│   └── rules/
│       └── go-guidelines.md       # Regras de engenharia e estilo Go
├── part1-gemini-go-sdk/           # Prática 1: Cat Picture Generator (Nano Banana 2 Lite)
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── README.md
├── part2-gemini-antigravity/      # Prática 2: Gerador de Música com Lyria 3 + Cloud Run
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── README.md
└── part3/                         # Bloco 3 (Em aberto)
    └── README.md
```

---

## 🛠️ Pré-requisitos Rápidos

1. **Go 1.24+**: Instalado e configurado no seu `PATH`.
2. **Antigravity CLI (`agy`)**: Instalado para as práticas da Parte 2.
3. **Credenciais de API**:
   - Chave do **Google AI Studio** (`GEMINI_API_KEY`), **ou**
   - Projeto no **Google Cloud** com Application Default Credentials configuradas (`gcloud auth application-default login` e `GOOGLE_CLOUD_PROJECT`).

Consulte o [Guia de Setup](setup/README.md) para o passo a passo completo de configuração de ambiente.

---

## 📖 Leituras Complementares

* [Gemini para Desenvolvedores Go: A Família de Modelos Gemini](https://danicat.dev/posts/gemini-for-go-developers-part-1-model-family/)
* [Gemini para Desenvolvedores Go: Programando com o Gemini](https://danicat.dev/posts/gemini-for-go-developers-part-2-coding-with-gemini/)
* [Documentação do Go GenAI SDK (`google.golang.org/genai`)](https://pkg.go.dev/google.golang.org/genai)
* [Documentação Oficial do Antigravity](https://antigravity.google/docs)
