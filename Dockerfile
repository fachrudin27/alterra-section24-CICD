FROM golang:latest as clean

WORKDIR /app
COPY . .

RUN go build -tags netgo -o main.app .

FROM alpine:latest

COPY --from=clean /app/main.app .
COPY --from=clean /app/.env .

CMD ["/main.app"]