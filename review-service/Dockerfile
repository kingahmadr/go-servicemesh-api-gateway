# --- Tahap 1: Build ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod dan sum lalu download dependency
COPY go.mod ./
RUN go mod download

# Copy sisa source code
COPY . .

# Build aplikasi menjadi static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# --- Tahap 2: Final Image ---
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary yang sudah di-build dari tahap builder
COPY --from=builder /app/main .

# Jalankan aplikasi
CMD ["./main"]