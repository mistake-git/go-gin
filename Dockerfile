FROM golang:1.18.1-alpine

WORKDIR /app

COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

CMD ["go", "run", "main.go"]