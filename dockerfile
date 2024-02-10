FROM golang:alpine
WORKDIR /app
ADD . .
RUN go build -o /app
CMD ['/app']
