# Guia de Configuração do Ambiente — Workshop GopherCon LATAM 2026

Este guia cobre todos os passos necessários para configurar seu computador antes ou durante o workshop.

---

## 1. Instalação do Go

Certifique-se de estar utilizando o **Go 1.24+**:

```bash
go version
```

Se precisar instalar ou atualizar:
* Acesse: [go.dev/dl](https://go.dev/dl/)
* No macOS (Homebrew): `brew install go`
* No Linux: siga o tarball oficial de `go.dev`

---

## 2. Configuração de Credenciais da Gemini API

Você pode utilizar qualquer uma das duas opções abaixo:

### Opção A: Google AI Studio (Recomendado para início rápido)
1. Crie uma chave de API gratuita em: [aistudio.google.com](https://aistudio.google.com)
2. Exporte a variável de ambiente no seu terminal:
   ```bash
   export GEMINI_API_KEY="sua-chave-aqui"
   ```

### Opção B: Gemini Enterprise / Google Cloud (Vertex AI)
1. Certifique-se de ter o Google Cloud SDK instalado (`gcloud`).
2. Efetue o login com Application Default Credentials (ADC):
   ```bash
   gcloud auth application-default login
   ```
3. Exporte a variável com o ID do seu projeto Google Cloud:
   ```bash
   export GOOGLE_CLOUD_PROJECT="seu-projeto-gcp-aqui"
   ```

---

## 3. Instalação da Antigravity CLI (`agy`)

Para as práticas da Parte 2:
1. Siga as instruções oficiais de instalação em [antigravity.google](https://antigravity.google).
2. Valide a instalação no terminal:
   ```bash
   agy --version
   ```

---

## 4. Servidores MCP e Ferramentas Recomendadas

Para habilitar a experiência agentiva completa para Go:

```bash
# 1. MCP Oficial do gopls (já incluído no ecossistema Go)
go install golang.org/x/tools/gopls@latest

# 2. Godoctor MCP
go install github.com/danicat/godoctor@latest

# 3. Linter oficial da comunidade
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```
