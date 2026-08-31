# Diretrizes de Desenvolvimento Go para Agentes

Ao interagir e modificar o código neste repositório, siga rigorosamente as seguintes diretrizes:

## Padrões de Código Go
1. **Idiomatic Go:**
   - Siga as convenções de `Effective Go` e `Go Code Review Comments`.
   - Evite abstrações excessivas e pacotes utilitários genéricos (`util`, `common`).
   - Mantenha funções concisas e com responsabilidade única.

2. **Tratamento de Erros:**
   - Sempre trate erros explicitamente. Nunca ignore retornos de erro com `_`.
   - Adicione contexto aos erros com `fmt.Errorf("contexto: %w", err)`.

3. **Validação e Loop de Qualidade:**
   - Sempre valide o código executando testes e linters (`go test ./...`, `go vet ./...`).
   - Mantenha dependências limpas e sincronizadas (`go mod tidy`).
