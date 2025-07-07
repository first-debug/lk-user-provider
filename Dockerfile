FROM golang:1.24-alpine AS builder


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN /go/bin/schema-fetcher --url https://raw.githubusercontent.com/first-debug/lk-graphql-schemas/master/schemas/user-provider/schema.graphql --output graph/schema.graphqls
RUN go generate ./...


RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/server ./cmd/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

#TODO: Конфиг подгружать через том (-v) 
COPY config/config_local.yml ./config/config_local.yml
#TODO: .env файл подгружать при запуске контейнера:  docker run -p 8080:8080 --env-file ./.env lk-user-service
COPY .env /.env

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/server"]

CMD ["--config", "/app/config/config_local.yml"]
