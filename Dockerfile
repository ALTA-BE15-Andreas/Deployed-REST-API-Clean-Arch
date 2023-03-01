FROM golang:1.19-alpine

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o rest-api-clean-arch

CMD ["./rest-api-clean-arch"]