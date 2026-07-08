FROM golang:alpine

WORKDIR /ViniciusSilva-golang-coding-interview

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./ViniciusSilva-golang-coding-interview

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./ViniciusSilva-golang-coding-interview"