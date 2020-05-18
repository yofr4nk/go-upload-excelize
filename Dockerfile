FROM golang:latest

WORKDIR /go/src/github.com/yofr4nk/go-upload-excelize

COPY ./ /go/src/github.com/yofr4nk/go-upload-excelize

RUN go get github.com/githubnemo/CompileDaemon

CMD CompileDaemon --build="go build main.go" --command=./main
