FROM golang:1.20.7

WORKDIR /app

COPY ./src .

RUN go build -o app main.go

EXPOSE 7070

CMD ["./app"]
