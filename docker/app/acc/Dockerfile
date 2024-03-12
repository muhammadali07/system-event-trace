FROM golang:1.18-alpine

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

COPY ../services/acc/go.mod ../services/acc/go.sum ./

RUN apk add --no-cache git

RUN go mod download

# Menyalin semua directory yang berisi kode sumber aplikasi
COPY ../services/acc/cmd ./
COPY ../services/acc/internal ./

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux go build -o main ../services/acc/main.go

EXPOSE 50051

# Run the executable
CMD ["/app/main"]