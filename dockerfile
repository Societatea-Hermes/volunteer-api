FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./volunteer-api cmd/api/*.go



FROM alpine

COPY --from=builder /app/volunteer-api /app/volunteer-api

EXPOSE 8080

CMD ["/app/volunteer-api"]
