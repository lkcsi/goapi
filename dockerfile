FROM golang:alpine
WORKDIR /app
ADD . .
CMD go build
CMD ./goapi