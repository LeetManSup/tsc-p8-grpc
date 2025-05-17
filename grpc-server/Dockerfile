FROM golang:1.24-alpine

WORKDIR /app

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY .. .

WORKDIR /app/grpc-server

RUN go build -o grpc-server .

EXPOSE 50051

CMD ["./grpc-server"]
