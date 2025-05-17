FROM golang:1.24-alpine

WORKDIR /app

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY .. .

WORKDIR /app/grpc-client

RUN go build -o grpc-client .

CMD ["./grpc-client"]
