FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./

RUN go build -o svc

ENV PORT 8080

EXPOSE $PORT

RUN find . -name "*.go" -type f -delete

CMD ["./svc"]