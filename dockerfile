FROM golang:alpine
WORKDIR /app
ADD . .
RUN go build -o /goapi 
CMD ['/goapi']