#!/usr/bin/env bash

# GopherCon LATAM 2026: Script de Verificacao de Ambiente
# Workshop: Gemini para Desenvolvedores Go

set -u

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
BOLD='\033[1m'
NC='\033[0m'

ERRORS=0
WARNINGS=0

echo -e "${BLUE}${BOLD}"
echo "=========================================================================="
echo "   GopherCon LATAM 2026: Workshop Environment Check                       "
echo "=========================================================================="
echo -e "${NC}"

pass() {
    echo -e "  [${GREEN}✅ PASS${NC}] $1"
}

fail() {
    echo -e "  [${RED}❌ FAIL${NC}] $1"
    echo -e "         ${YELLOW}> FIX: $2${NC}\n"
    ERRORS=$((ERRORS + 1))
}

warn() {
    echo -e "  [${YELLOW}⚠️  WARN${NC}] $1"
    echo -e "         ${YELLOW}> NOTE: $2${NC}\n"
    WARNINGS=$((WARNINGS + 1))
}

# 1. Compilador Go (1.26+)
echo -e "${BOLD}1. Verificando Compilador Go (1.26+)...${NC}"
if command -v go >/dev/null 2>&1; then
    GO_VER_STR=$(go version | awk '{print $3}' | sed 's/go//')
    GO_MAJOR=$(echo "$GO_VER_STR" | cut -d. -f1)
    GO_MINOR=$(echo "$GO_VER_STR" | cut -d. -f2)
    if [ "$GO_MAJOR" -gt 1 ] || { [ "$GO_MAJOR" -eq 1 ] && [ "$GO_MINOR" -ge 26 ]; }; then
        pass "Go versao $GO_VER_STR instalada."
    else
        fail "Go versao $GO_VER_STR instalada, mas Go 1.26+ e obrigatorio." "Atualize o Go em https://go.dev/dl/ ou 'brew upgrade go'."
    fi
else
    fail "Compilador Go nao encontrado." "Instale Go 1.26+ em https://go.dev/doc/install."
fi

# 2. Antigravity CLI (agy)
echo -e "${BOLD}2. Verificando Antigravity CLI (agy)...${NC}"
if command -v agy >/dev/null 2>&1; then
    AGY_VER=$(agy --version 2>&1 | head -n1)
    pass "Antigravity CLI (agy) instalado ($AGY_VER)."
elif command -v antigravity >/dev/null 2>&1; then
    AGY_VER=$(antigravity --version 2>&1 | head -n1)
    pass "Antigravity CLI instalado ($AGY_VER)."
else
    fail "Antigravity CLI (agy) nao encontrado no PATH." "Instale o agy CLI conforme instrucoes do workshop."
fi

# 3. Kungfu CLI (Opcional)
echo -e "${BOLD}3. Verificando Kungfu CLI (Opcional)...${NC}"
if command -v kungfu >/dev/null 2>&1; then
    KUNGFU_VER=$(kungfu version 2>&1 | head -n1)
    pass "Kungfu CLI instalado ($KUNGFU_VER)."
else
    warn "Kungfu CLI nao encontrado." "Opcional. Para instalar: 'go install github.com/danicat/kungfu@latest'."
fi

# 4. Google Cloud SDK e ADC
echo -e "${BOLD}4. Verificando Google Cloud SDK e Credenciais (ADC)...${NC}"
if command -v gcloud >/dev/null 2>&1; then
    pass "gcloud CLI instalado."
    if gcloud auth application-default print-access-token >/dev/null 2>&1; then
        pass "GCP Application Default Credentials (ADC) ativas."
    else
        fail "GCP ADC ausente ou expirado." "Execute 'gcloud auth application-default login'."
    fi
else
    fail "gcloud CLI nao encontrado." "Instale o Google Cloud SDK: https://cloud.google.com/sdk/docs/install."
fi

# 5. Variavel GOOGLE_CLOUD_PROJECT
echo -e "${BOLD}5. Verificando variavel GOOGLE_CLOUD_PROJECT...${NC}"
if [ -n "${GOOGLE_CLOUD_PROJECT:-}" ]; then
    pass "GOOGLE_CLOUD_PROJECT definido: $GOOGLE_CLOUD_PROJECT"
else
    fail "GOOGLE_CLOUD_PROJECT nao definido no ambiente." "Defina 'export GOOGLE_CLOUD_PROJECT=seu-projeto-gcp'."
fi

# 6. Git
echo -e "${BOLD}6. Verificando Git...${NC}"
if command -v git >/dev/null 2>&1; then
    pass "git instalado ($(git --version))."
else
    fail "git nao encontrado." "Instale o git para versionamento e submissao via PR."
fi

# 7. Node.js / npx (Necessario para MCPs e Skills)
echo -e "${BOLD}7. Verificando npx (Node.js)...${NC}"
if command -v npx >/dev/null 2>&1; then
    NPX_VER=$(npx --version 2>/dev/null || npx --version 2>&1 | tail -n1)
    pass "npx instalado ($NPX_VER)."
else
    fail "npx nao encontrado." "Instale Node.js (que inclui npx) em https://nodejs.org/ ou 'brew install node'."
fi

# 8. uv / Python (Opcional)
echo -e "${BOLD}8. Verificando uv / Python...${NC}"
if command -v uv >/dev/null 2>&1; then
    pass "uv instalado ($(uv --version | head -n1))."
elif command -v python3 >/dev/null 2>&1; then
    pass "python3 instalado ($(python3 --version))."
else
    warn "uv ou python3 nao encontrado." "Opcional, util para scripts auxiliares."
fi

# Resumo Final
echo -e "\n--------------------------------------------------------------------------"
if [ $ERRORS -eq 0 ]; then
    echo -e "${GREEN}${BOLD}✅ AMBIENTE VERIFICADO COM SUCESSO!${NC}"
    if [ $WARNINGS -gt 0 ]; then
        echo -e "${YELLOW}Avisos: $WARNINGS aviso(s) detectado(s).${NC}"
    fi
    exit 0
else
    echo -e "${RED}${BOLD}❌ FALHA NA VERIFICACAO: $ERRORS erro(s) detectado(s).${NC}"
    echo -e "${RED}Aplique as correcoes indicadas acima antes de prosseguir.${NC}"
    exit 1
fi
