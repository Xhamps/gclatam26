# Estágio de compilação
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copia dependências
COPY go.mod go.sum* ./
RUN go mod download || true

# Copia código-fonte e compila binário estático
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server main.go

# Estágio final de execução leve
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/
COPY --from=builder /app/server .

# Cloud Run define a variável de ambiente PORT (padrão: 8080)
ENV PORT=8080
EXPOSE 8080

CMD ["./server"]
