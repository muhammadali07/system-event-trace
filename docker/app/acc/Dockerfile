FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Menyalin semua directory yang berisi kode sumber aplikasi
COPY cmd ./
COPY internal ./

RUN go build -o main

CMD ["./main"]