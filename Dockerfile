# imagem base
FROM golang:latest

# diretório de trabalho
RUN mkdir /app

WORKDIR /app

# copia o código da API para o diretório de trabalho
COPY . .

# compila o código para um arquivo executável
RUN go build src/main.go

# expõe a porta 8080
EXPOSE 8080

# define o comando para iniciar a API
CMD [ "./main" ]