FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /mysql-db-ex

EXPOSE 8080

CMD [ "/mysql-db-ex" ]