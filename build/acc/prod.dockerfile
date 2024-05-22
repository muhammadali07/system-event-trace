FROM golang:1.22-alpine as build

WORKDIR /app
COPY . .

# check unit test with code coverage treshold

RUN go build -o main ./src/main.go

FROM scratch
COPY --from=build /app/main /main

ENTRYPOINT ["/main"]
