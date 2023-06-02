FROM golang:1.19.3

WORKDIR /usr/src/app

COPY . .

RUN go install github.com/cosmtrek/air@latest

RUN go mod tidy

CMD docker run -v .:/usr/src/app -p 3000:3000  air main.go -b 0.0.0.0  docker run -v postgres-db:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_USER=silvvf -e POSTGRES_PASSWORD=silvvfpostgresdbpassword777 -e POSTGRES_DB=silvvfvalorantaapigo postgres:alpine
