FROM golang:alpine

WORKDIR /app

COPY . .

EXPOSE 3000

ENTRYPOINT [ "go", "run", "main.go" ]