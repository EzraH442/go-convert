FROM golang:1.21.4
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY *.go ./
# Build
RUN GOOS=linux go build -o ./build/server.go

EXPOSE 8090

# Run
CMD ["./build/server.go"]