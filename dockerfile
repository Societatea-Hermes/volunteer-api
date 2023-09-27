FROM golang:alpine as builder

WORKDIR /tmp

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o ./volunteer-api .



FROM alpine

COPY --from=builder /tmp/volunteer-api /app/volunteer-api

EXPOSE 8080

CMD ["/app/volunteer-api"]
