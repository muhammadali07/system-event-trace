FROM golang:1.18-alpine

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Create a separate directory for go.mod and go.sum
RUN mkdir -p /go/src/app/modules

# Copy the go.mod and go.sum files to the modules directory
COPY ../services/acc/go.mod ../services/acc/go.sum /go/src/app/modules/

RUN apk add --no-cache git

# Move the go.mod and go.sum files to the working directory
WORKDIR /go/src/app
RUN ln -s modules/go.mod . && ln -s modules/go.sum .

RUN go mod download

# Copy the remaining source code
COPY ../services/acc/cmd ./
COPY ../services/acc/internal ./

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux go build -o main ../services/acc/main.go

EXPOSE 50051

# Run the executable
CMD ["/app/main"]