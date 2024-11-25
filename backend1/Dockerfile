FROM golang:1.23.3-alpine

RUN mkdir app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

ARG PORT=:3333 

ENV PORT=${PORT}

COPY . .

RUN go build -o main .

CMD ["./main"]