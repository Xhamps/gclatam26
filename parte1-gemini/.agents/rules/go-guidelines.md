# Diretrizes de Desenvolvimento Go para Agentes

Ao interagir e modificar o código neste repositório, siga rigorosamente as seguintes diretrizes:

## Padrões de Código Go
1. **Idiomatic Go:**
   - Siga as convenções de `Effective Go` e `Go Code Review Comments`.
   - Evite abstrações excessivas e pacotes utilitários genéricos (`util`, `common`).
   - Mantenha funções concisas e com responsabilidade única.
   - Evite side effects e side inputs
   - Evite estado mutavel sempre que possivel

2. **Tratamento de Erros:**
   - Sempre trate erros explicitamente. Nunca ignore retornos de erro com `_`.
   - Adicione contexto aos erros com `fmt.Errorf("contexto: %w", err)`.

3. **Validação e Loop de Qualidade:**
   - Sempre valide o código executando testes e linters (`go test`, `go vet`).
   - O escopo do teste ou lint deve ser o minimo necessario para maximizar paralelismo de agentes (e.g. ao trabalhar na package ./internal/foo execute testes em ./internal/foo e nao em ./...). Apenas execute ./... se for o unico agente trabalhando em uma tarefa
   - Mantenha dependências limpas e sincronizadas (`go mod tidy`).

4. Pacotes internos do projeto devem ser organizados em ./internal. Jamais use ./pkg a não ser que esteja trabalhando num projeto da família do kubernetes.