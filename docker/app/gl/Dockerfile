FROM golang:alpine
WORKDIR /services/gl
COPY . /app
RUN go build -o main .
CMD ["/app/main"]
