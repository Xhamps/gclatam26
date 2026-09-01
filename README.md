# Gemini para Desenvolvedores Go 🐾 🚀
### Workshop Oficial — GopherCon LATAM 2026
**Data:** Quinta-feira, 3 de Setembro de 2026  
**Duração:** 4 horas (3 blocos práticos com intervalos)

---

## 🎯 Sobre o Workshop

Este repositório contém os materiais de apoio, guias de setup, configurações agentivas e instruções para o workshop **"Gemini para Desenvolvedores Go"** e para a **GopherCon LATAM Mini Game Jam**.

O objetivo do workshop é capacitar desenvolvedores Go a construir aplicações nativas de IA e utilizar o ecossistema agentivo do Antigravity (`agy`), culminando em um hackathon prático onde os participantes constroem e submetem jogos 2D em Go utilizando modelos multimodais do Gemini (Nano Banana e Lyria 3).

---

## 📅 Grade e Cronograma (4 Horas)

| Horário / Duração | Bloco | Tema | Prática |
| :--- | :--- | :--- | :--- |
| **50 min** | **Parte 1** | **A Família de Modelos Gemini**<br>• Nomenclatura e catálogo (Gemini 3.7 Flash, Nano Banana 2 Lite, Veo, Lyria 3, Gemma)<br>• Features: Thinking Mode, Function Calling (4 etapas), Grounding, Structured Outputs<br>• Superfícies de API (GenerateContent, Interactions, Live, Batch, Managed Agents)<br>• Autenticação: Google AI Studio vs. Gemini Enterprise | **Prática 1 (Smoke Test):** Validação do ambiente e credenciais gerando a foto do gato mago com Go GenAI SDK e Nano Banana 2 Lite (`gemini-3.1-flash-lite-image`). |
| **10 min** | *Intervalo* | *Coffee Break & Suporte de Setup* | — |
| **50 min** | **Parte 2** | **Desenvolvimento com Gemini e Antigravity**<br>• Go na Era da IA: Afinidade Agentiva e o Loop Agentivo Determinístico<br>• O ecossistema Antigravity: Antigravity 2.0 (Agent Manager), `agy` CLI (em Go), Antigravity IDE<br>• Customizações no workspace (`.agents/`): Regras, MCPs (`gopls`, `godoctor`), Skills, Hooks e Subagentes | **Prática 2:** Desenvolvimento assistido com `agy` do microserviço de trilhas sonoras com Lyria 3 + deploy no Cloud Run. |
| **15 min** | *Intervalo* | *Coffee Break & Formação de Times* | — |
| **110 min** | **Parte 3** | **GopherCon LATAM Mini Game Jam**<br>• **Kickoff (15 min):** Introdução ao Ebitengine v2, apresentação das skills de game dev (`.agents/skills/`) e anúncio do tema.<br>• **Jam Hands-on (70 min):** Desenvolvimento dos jogos em Go usando `agy`, Nano Banana e Lyria 3.<br>• **Showcase & Julgamento (25 min):** Demonstração ao vivo dos jogos, votação e premiação! | **Prática 3 (Hackathon):** Criação e submissão de um jogo 2D via Pull Request (Fork + PR). |
| **5 min** | *Encerramento* | *Q&A, Fotos & Teaser do Keynote* | — |

---

## 📂 Estrutura do Repositório

```text
.
├── README.md                      # Guia geral do workshop
├── check_env.sh                   # Script de verificação prévia do ambiente
├── setup/                         # Instruções de setup de ambiente e credenciais
│   └── README.md
├── .agents/                       # Customizações e Skills do Antigravity
│   ├── mcp_config.json            # Configuração dos servidores MCP recomendados
│   ├── rules/                     # Regras de código Go para o agente
│   └── skills/                    # Skills especializadas (vibe-game-developer, ebitengineer, etc.)
├── part1-gemini-go-sdk/           # Prática 1: Smoke Test & Cat Picture Generator
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── README.md
├── part2-gemini-antigravity/      # Prática 2: Desafio Lyria 3 + Cloud Run
│   ├── go.mod
│   ├── main.go
│   └── README.md
├── part3/                         # Bloco 3: Visão Geral da Game Jam
│   └── README.md
└── mini-game-jam/                 # Área do Hackathon e Submissões
    └── README.md                  # Regras da Jam e Guia de Submissão via Fork + PR
```

---

## 🛠️ Como Começar

1. **Valide seu ambiente com o script:**
   ```bash
   ./check_env.sh
   ```
2. **Consulte o [Guia de Setup](setup/README.md)** se faltar alguma dependência ou credencial.
3. **Durante a Game Jam (Parte 3):** Faça o fork do repositório e siga as instruções em [mini-game-jam/README.md](mini-game-jam/README.md).

---

## 📖 Leituras & Recursos Complementares

* [Gemini para Desenvolvedores Go: A Família de Modelos Gemini](https://danicat.dev/posts/gemini-for-go-developers-part-1-model-family/)
* [Gemini para Desenvolvedores Go: Programando com o Gemini](https://danicat.dev/posts/gemini-for-go-developers-part-2-coding-with-gemini/)
* [Catálogo de Skills para Antigravity (`skills.danicat.dev`)](https://skills.danicat.dev)
* [Documentação do Go GenAI SDK (`google.golang.org/genai`)](https://pkg.go.dev/google.golang.org/genai)
* [Documentação Oficial do Ebitengine v2](https://ebitengine.org)
