FROM golang:1.19-alpine3.16
Run mkdir /app
WORKDIR /app
COPY . /app
Run go build -o main .
CMD ["/app/main"]
