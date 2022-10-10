FROM golang:1.19-alpine3.16
Run mkdir /app
ADD . /app
RUN go mod init
Run go build -o main .
CMD ["/app/main"]
