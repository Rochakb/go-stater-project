FROM golang:1.19-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go build cmd/employee_app/main.go

EXPOSE 8000

CMD ["./cli_cmd","start"]