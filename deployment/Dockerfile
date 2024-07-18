FROM golang:1.22-alpine

RUN apk add --update curl

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /server

EXPOSE 8080

CMD ["/server"]