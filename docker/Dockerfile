# Imagem base para construção
FROM golang:1.24-alpine AS builder

# Diretório de trabalho
WORKDIR /app

# Copiar os arquivos de dependência
COPY ../api/go.mod ../api/go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o código fonte
COPY ../api/ .

# Construir a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Imagem final mais leve
FROM alpine:latest

# Adicionar certificados CA para possíveis conexões HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o binário compilado da etapa de construção
COPY --from=builder /app/main .

# Expor a porta
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"]