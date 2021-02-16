FROM golang:1.14

WORKDIR /app

COPY . .

RUN go mod download && go mod vendor && go mod verify

RUN GOOS=linux GOARCH=amd64 go build -o contract-check

FROM alpine as run

WORKDIR /app

COPY --from=0 /app/contract-check /app/check

CMD ["/app/check"]

