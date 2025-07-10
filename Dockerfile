FROM golang:1.24-alpine AS builder

WORKDIR /build-dir

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/first-debug/lk-tools/schema-fetcher@latest

RUN /go/bin/schema-fetcher -url first-debug/lk-graphql-schemas/master/schemas/user-provider/schema.graphql -output graph/schema.graphqls

RUN go run github.com/99designs/gqlgen

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /build-dir/server ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build-dir/server .

# -v ./config:/config
# --env-file .env

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

EXPOSE 8080

CMD [ "/app/server" ]
