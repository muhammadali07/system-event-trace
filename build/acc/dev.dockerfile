FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT [ "air", "-c", "./builds/golang/.air.toml" ]
