FROM golang:latest

WORKDIR /app

COPY ./main.go /app
COPY ./go.mod /app

RUN go build

EXPOSE 8080

CMD ["./http"]
