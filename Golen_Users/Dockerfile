FROM golang:1.26

WORKDIR /app

# Copy go mod files first (important for caching)
COPY go.mod go.sum ./

RUN go mod tidy

# Copy rest of the code
COPY . .

RUN go build -o main ./cmd

EXPOSE 8002

CMD ["./main"]