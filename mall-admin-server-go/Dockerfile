FROM golang:latest

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go build -o main .

EXPOSE 8081
ENTRYPOINT ["./main"]
