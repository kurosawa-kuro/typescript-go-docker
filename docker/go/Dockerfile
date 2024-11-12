FROM golang:1.23.2
RUN apt-get update && apt-get install -y build-essential make

RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/cosmtrek/air@v1.28.0
RUN go install github.com/joho/godotenv/cmd/godotenv@latest

# migration
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate.linux-amd64 /usr/bin/migrate

WORKDIR /app

CMD ["air", "-c", "backend/.air.toml"]