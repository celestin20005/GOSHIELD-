
# Utilise une image légère de Go pour compiler
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o goshield .

# Utilise une image ultra-légère pour exécuter
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/goshield .
# Expose le port de ton application (ex: 8080)
EXPOSE 8080
CMD ["./goshield"]

