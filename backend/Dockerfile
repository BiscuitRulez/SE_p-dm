FROM golang:1.23-alpine
WORKDIR /go/src/github.com/sut67/team10
COPY backend ./
RUN apk add build-base
RUN go mod tidy
RUN go build -o main .
EXPOSE 8000
CMD ["./main"]
