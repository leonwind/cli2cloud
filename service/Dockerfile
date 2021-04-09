FROM golang:1.16.3-alpine3.13
COPY . /service
WORKDIR /service
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]

