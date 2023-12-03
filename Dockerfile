FROM golang:1.21.4
WORKDIR /app

RUN apt-get update && apt-get -y install libheif-dev

COPY go.mod go.sum ./

RUN go mod download
COPY *.go ./
# Build
RUN GOOS=linux go build -o ./build/server.go

EXPOSE 8090

# Run
CMD ["./build/server.go"]