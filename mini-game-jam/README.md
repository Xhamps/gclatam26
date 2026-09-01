# 🎮 GopherCon LATAM 2026 Mini Game Jam

Bem-vindos à **GopherCon LATAM 2026 Mini Game Jam**!

Durante o terceiro bloco do workshop **"Gemini para Desenvolvedores Go"**, os participantes competirão individualmente ou em duplas criando um jogo 2D em Go utilizando **Ebitengine v2**, **Antigravity (`agy`)** e os modelos de IA da família Gemini (**Nano Banana** para arte e **Lyria 3** para música/áudio).

---

## ⏱️ Formato e Regras da Jam

* **Duração do Desenvolvimento:** ~70 minutos.
* **Formação dos Times:** Individual ou duplas.
* **Engine / Tecnologias:** Go 1.24+, Ebitengine v2 (`github.com/hajimehoshi/ebitengine/v2`), Antigravity (`agy`), Gemini API (Nano Banana & Lyria 3).
* **Tema:** Revelado ao vivo no kickoff da Parte 3!

---

## 🚀 Como Participar e Fluxo de Submissão (Fork + PR)

As submissões serão avaliadas ao vivo via **Pull Request no GitHub**. Siga o fluxo abaixo:

### Passo 1: Fork e Clone
1. Faça um **Fork** deste repositório para o seu perfil no GitHub.
2. Clone o seu fork na sua máquina:
   ```bash
   git clone https://github.com/SEU-USUARIO/gclatam26.git
   cd gclatam26
   ```

### Passo 2: Criar o Diretório do Seu Jogo
Crie uma pasta dedicada para o seu jogo dentro de `mini-game-jam/`:
```bash
mkdir -p mini-game-jam/nome-do-seu-jogo
cd mini-game-jam/nome-do-seu-jogo
```

### Passo 3: Inicializar o Módulo Go
```bash
go mod init nome-do-seu-jogo
go get github.com/hajimehoshi/ebitengine/v2@latest
```

### Passo 4: Desenvolver com Antigravity (`agy`)
Inicie o `agy` na raiz do projeto ou no diretório do jogo e utilize as skills disponíveis em `.agents/skills/`:
* Inicie o brainstorm digitando: `/vibe-game-developer` ou `/game-design` (para conduzir o `/grill-me`).
* Gere a arquitetura e o game loop com `/ebitengineer`.
* Gere sprites e texturas com `/nano-banana`.
* Crie efeitos sonoros e músicas de fundo com `/lyria` ou `/procedural-composer`.

### Passo 5: Documentar o Jogo
Crie um arquivo `README.md` dentro de `mini-game-jam/nome-do-seu-jogo/` contendo:
* **Título do Jogo**
* **Autor(es) / GitHub Handles**
* **Sinopse e Mecânicas**
* **Controles** (ex: Setas / WASD / Espaço)
* **Instruções de Execução** (ex: `go run .`)
* **Screenshot ou GIF** (opcional, mas altamente recomendado!)

### Passo 6: Submissão via Pull Request (PR)
1. Faça commit e push para o seu fork:
   ```bash
   git add mini-game-jam/nome-do-seu-jogo
   git commit -m "Add Game Jam entry: Nome do Jogo"
   git push origin main
   ```
2. Abra um **Pull Request (PR)** para o repositório upstream principal com o título:
   ```text
   [Game Jam] Nome do Jogo - @seu-usuario
   ```

---

## 🏆 Critérios de Avaliação

Os jogos serão jogados e demonstrados ao vivo na sessão de **Showcase & Playtest** e avaliados nos seguintes critérios:

1. **🎨 Criatividade & Fidelidade ao Tema (40%):** Originalidade da ideia e interpretação do tema da Jam.
2. **🕹️ Jogabilidade & Diversão (30%):** O jogo é divertido, responsivo e funcional?
3. **🤖 Uso de IA e Antigravity (20%):** Aplicação criativa de geração de código (`agy`), arte (Nano Banana) e áudio (Lyria 3).
4. **🚀 Polish e Execução Técnica (10%):** Fluidez do game loop em Go e ausência de bugs impeditivos.

---

## 🎁 Premiação

O vencedor ou dupla vencedora da Game Jam receberá o **Prêmio Oficial da GopherCon LATAM Mini Game Jam** durante o encerramento do workshop!
