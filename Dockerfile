FROM golang:alpine
WORKDIR /workspace/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o lk-user-provider
EXPOSE 8080
CMD ["./lk-user-provider"]
