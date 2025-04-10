# Dockerfile
FROM golang:1.23.4 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN go build -o server ./cmd/server

EXPOSE 50051

CMD ["/app/server"]
