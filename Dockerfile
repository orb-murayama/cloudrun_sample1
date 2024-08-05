# Use the official Golang image to create a build artifact.
# https://hub.docker.com/_/golang
FROM golang:1.20 AS builder

# Copy local code to the container image.
WORKDIR /app

# Goモジュールファイルと.envファイルをコピー
#COPY go.mod go.sum ./
#COPY .env ./

# 依存関係をインストール
#RUN go mod download

COPY . .

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o server

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server

# Run the web service on container startup.
CMD ["/server"]
