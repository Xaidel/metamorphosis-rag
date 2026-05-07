FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/app ./cmd

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /out/app /usr/local/bin/app

EXPOSE 6969

ENTRYPOINT [ "/usr/local/bin/app" ]