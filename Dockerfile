FROM golang:1.19-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o gophermart ./cmd/gophermart/main.go

CMD ["./gophermart"]