FROM golang:latest

WORKDIR /go-upload-excelize

COPY ./ /go-upload-excelize

RUN go get github.com/githubnemo/CompileDaemon

CMD CompileDaemon --build="go build main.go" --command=./main
