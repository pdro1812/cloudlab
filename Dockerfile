# currency-converter/Dockerfile

# --- Estágio 1: Build ---
# Usamos uma imagem completa do Go para compilar a aplicação
FROM golang:1.22-alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos de gerenciamento de dependências primeiro
# Isso aproveita o cache do Docker. As dependências só serão baixadas novamente se o go.mod/sum mudar.
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o resto do código-fonte
COPY . .

# Compila a aplicação. As flags são importantes:
# - CGO_ENABLED=0: Cria um binário estaticamente linkado (não depende de libs do sistema)
# - o /app/server: O nome e local do arquivo de saída (nosso executável)
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/api/main.go


# --- Estágio 2: Final ---
# Começamos com uma imagem Alpine zerada, que é super leve
FROM alpine:latest

# Copia APENAS o executável compilado do estágio de build para a imagem final
COPY --from=builder /app/server /server

# Expõe a porta que a nossa aplicação usa (definida no .env como 8080)
EXPOSE 8080

# Define o comando que será executado quando o container iniciar.
ENTRYPOINT [ "/server" ]