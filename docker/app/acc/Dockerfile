FROM golang:alpine
WORKDIR /services/acc
COPY . /app
RUN go build -o main .
CMD ["/app/main"]
