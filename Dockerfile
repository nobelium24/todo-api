FROM golang:1.22.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN chmod +x ./wait-for-it.sh

RUN ls -la /app 

WORKDIR /app/cmd/app

RUN go build -o /app/main .

WORKDIR /app

EXPOSE 8080

CMD ["./wait-for-it.sh", "db:5432", "--","./main"]