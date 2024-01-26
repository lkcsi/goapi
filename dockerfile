FROM golang:latest
RUN apt update
RUN apt install -y vim
WORKDIR /app
ADD . .
CMD go run .