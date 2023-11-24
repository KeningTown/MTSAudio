FROM golang

WORKDIR /app/dbwriter

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

#download psql
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go build -o ./main ./cmd/mtsaudio/main.go
