#!/usr/bin/env bash

# ==============================================================================
# GopherCon LATAM 2026: Workshop Environment Doctor Script
# Workshop: Gemini para Desenvolvedores Go & Mini Game Jam
# ==============================================================================

set -u

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
BOLD='\033[1m'
NC='\033[0m' # No Color

ERRORS=0
WARNINGS=0

echo -e "${BLUE}${BOLD}"
echo "=========================================================================="
echo "   GopherCon LATAM 2026: Workshop & Game Jam Environment Check            "
echo "=========================================================================="
echo -e "${NC}"

pass() {
    echo -e "  [${GREEN}PASS${NC}] $1"
}

fail() {
    echo -e "  [${RED}FAIL${NC}] $1"
    echo -e "         ${YELLOW}↳ FIX: $2${NC}\n"
    ERRORS=$((ERRORS + 1))
}

warn() {
    echo -e "  [${YELLOW}WARN${NC}] $1"
    echo -e "         ${YELLOW}↳ NOTE: $2${NC}\n"
    WARNINGS=$((WARNINGS + 1))
}

# 1. Check Go Compiler (1.24+)
echo -e "${BOLD}1. Verificando Compilador Go (1.24+)...${NC}"
if command -v go >/dev/null 2>&1; then
    GO_VER_STR=$(go version | awk '{print $3}' | sed 's/go//')
    GO_MAJOR=$(echo "$GO_VER_STR" | cut -d. -f1)
    GO_MINOR=$(echo "$GO_VER_STR" | cut -d. -f2)
    if [ "$GO_MAJOR" -gt 1 ] || { [ "$GO_MAJOR" -eq 1 ] && [ "$GO_MINOR" -ge 24 ]; }; then
        pass "Go versão $GO_VER_STR instalada."
    else
        fail "Go versão $GO_VER_STR instalada, mas Go 1.24+ é obrigatório." "Atualize o Go em https://go.dev/dl/ ou 'brew upgrade go'."
    fi
else
    fail "Compilador Go não encontrado." "Instale Go 1.24+ em https://go.dev/doc/install."
fi

# 2. Check Antigravity / agy CLI
echo -e "${BOLD}2. Verificando Antigravity CLI (agy)...${NC}"
if command -v agy >/dev/null 2>&1; then
    AGY_VER=$(agy --version 2>&1 | head -n1)
    pass "Antigravity CLI (agy) instalado ($AGY_VER)."
elif command -v antigravity >/dev/null 2>&1; then
    AGY_VER=$(antigravity --version 2>&1 | head -n1)
    pass "Antigravity CLI instalado ($AGY_VER)."
else
    warn "Antigravity CLI (agy) não encontrado no PATH." "Instale o agy CLI conforme as instruções em setup/README.md."
fi

# 3. Check Credenciais (GEMINI_API_KEY ou GCP ADC)
echo -e "${BOLD}3. Verificando Credenciais Gemini (API Key ou Google Cloud ADC)...${NC}"
HAS_AUTH=0
if [ -n "${GEMINI_API_KEY:-}" ] || [ -n "${GOOGLE_API_KEY:-}" ]; then
    pass "Chave GEMINI_API_KEY / GOOGLE_API_KEY configurada no ambiente."
    HAS_AUTH=1
fi

if command -v gcloud >/dev/null 2>&1; then
    if gcloud auth application-default print-access-token >/dev/null 2>&1; then
        pass "GCP Application Default Credentials (ADC) ativas."
        HAS_AUTH=1
    fi
fi

if [ $HAS_AUTH -eq 0 ]; then
    fail "Nenhuma credencial configurada." "Configure 'export GEMINI_API_KEY=...' ou execute 'gcloud auth application-default login'."
fi

# 4. Check Git
echo -e "${BOLD}4. Verificando Git...${NC}"
if command -v git >/dev/null 2>&1; then
    pass "git instalado ($(git --version))."
else
    fail "git não encontrado." "Instale o git para submeter o PR na Game Jam."
fi

# 5. Check uv / Python (Opcional para scripts de geração de mídia)
echo -e "${BOLD}5. Verificando Python / uv (Auxiliares de Mídia)...${NC}"
if command -v uv >/dev/null 2>&1; then
    pass "uv instalado ($(uv --version | head -n1))."
elif command -v python3 >/dev/null 2>&1; then
    pass "python3 instalado ($(python3 --version))."
else
    warn "uv/python3 não encontrado." "Instale o 'uv' se desejar rodar os scripts auxiliares de geração em linha de comando."
fi

# ==============================================================================
# Resumo Final
# ==============================================================================
echo -e "\n--------------------------------------------------------------------------"
if [ $ERRORS -eq 0 ]; then
    echo -e "${GREEN}${BOLD}🎉 AMBIENTE VERIFICADO COM SUCESSO! Tudo pronto para o workshop e a Game Jam.${NC}"
    if [ $WARNINGS -gt 0 ]; then
        echo -e "${YELLOW}Nota: $WARNINGS aviso(s) detectado(s), mas os componentes essenciais estão operacionais.${NC}"
    fi
    exit 0
else
    echo -e "${RED}${BOLD}❌ FALHA NA VERIFICAÇÃO: $ERRORS problema(s) detectado(s).${NC}"
    echo -e "${RED}Por favor, aplique as correções indicadas antes do início da sessão.${NC}"
    exit 1
fi
