FROM golang:1.24 AS builder
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server server.go

FROM scratch
COPY --from=builder /src/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]