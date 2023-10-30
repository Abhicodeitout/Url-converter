FROM golang:1.20.5

WORKDIR /app

COPY . /app

RUN go build main

EXPOSE 8080

CMD ["./main"]
