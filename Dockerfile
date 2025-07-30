# ---------------------
# Stage 1 — build
# ---------------------
FROM golang:latest AS builder

WORKDIR /app
COPY . .

RUN go build -o bigscanner ./cmd/app/main.go

# ---------------------
# Stage 2 — runtime
# ---------------------
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/bigscanner .

ENTRYPOINT ["./bigscanner"]

