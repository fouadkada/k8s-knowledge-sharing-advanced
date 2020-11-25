FROM golang:1.15 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY cmd/ cmd/
COPY internal/ internal/

RUN CGO_ENABLED=0 GOOS=linux go build -o /app.o ./cmd/main.go

FROM alpine:3.12
COPY --from=builder /app.o ./app.o
RUN chmod +x ./app.o
CMD ["./app.o"]
