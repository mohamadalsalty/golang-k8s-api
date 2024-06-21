# Build stage
FROM golang:1.22.4 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-w -s" -o main .

FROM scratch

WORKDIR /app

COPY --from=builder /app/main /app/main

ENV KUBECONFIG=/home/app/.kube/config

RUN chmod +x /app/main

EXPOSE 8080

RUN adduser -D app
USER app

ENTRYPOINT ["/app/main"]
