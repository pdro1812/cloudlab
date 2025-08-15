# Etapa 1: Build da aplicação Go
FROM golang:1.22-alpine AS builder

# Instala dependências do sistema necessárias
RUN apk add --no-cache git

# Cria diretório de trabalho
WORKDIR /app

# Copia os arquivos go.mod e go.sum primeiro (para aproveitar cache)
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila a aplicação em modo release (binário estático)
# O nome do output é 'myapp'
RUN go build -o myapp ./cmd/app

# Etapa 2: Imagem final mínima
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Copia binário compilado do builder para /myapp
COPY --from=builder /app/myapp /myapp

# Define comando padrão para executar o arquivo que foi copiado
CMD ["/myapp"]