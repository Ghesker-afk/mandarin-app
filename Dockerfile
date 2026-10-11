# --- Estágio 1: Build da aplicação ---
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Baixa as dependências do projeto
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código fonte
COPY . .

# Compila o binário em Go (desativando CGO para gerar um binário estático puro)
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server ./cmd/server/main.go

# --- Estágio 2: Imagem Final de Produção ---
FROM alpine:3.19

WORKDIR /root/

# Instala certificados ca-certificates caso precise de chamadas HTTPS externas no futuro
RUN apk --no-cache add ca-certificates

# Copia o binário compilado do estágio anterior
COPY --from=builder /app/bin/server .

# Expõe a porta em que o servidor HTTP vai rodar
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./server"]