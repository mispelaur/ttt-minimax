FROM golang:alpine

WORKDIR /go/src/github.com/mispelaur/ttt-minimax
COPY main.go ./

RUN go build -o main .
CMD ["./main"]