# Dockerfile
FROM amazonlinux:2

# Instale as dependências
RUN yum install -y \
    golang \
    unzip \
    zip \
    curl \
    less \
    groff \
    python3 \
    && pip3 install awscli

WORKDIR /app

# Copie o arquivo go.mod para inicializar as dependências
COPY go.mod ./

# Baixe as dependências e gere go.sum
RUN go mod tidy

# Copie o restante dos arquivos da aplicação
COPY . .

# Conceda permissão de execução aos scripts
RUN chmod +x /app/init.sh /app/entrypoint.sh

# Baixe novamente as dependências após copiar todos os arquivos
RUN go mod tidy

# Construa a aplicação
RUN go build -o /poc-go-dynamodb

EXPOSE 8080

# Use o script de entrada para iniciar o contêiner
ENTRYPOINT ["/bin/sh", "/app/entrypoint.sh"]
