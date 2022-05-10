FROM golang:latest

RUN mkdir -p $GOPATH/src/github.com/wasuken/infomap.local
WORKDIR $GOPATH/src/github.com/wasuken/infomap.local
ADD . $GOPATH/src/github.com/wasuken/infomap.local

RUN go mod tidy

CMD ["go", "run", "main.go"]
