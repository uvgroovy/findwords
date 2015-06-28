FROM golang

RUN apt-get update && apt-get install -y wamerican


ADD . /go/src/github.com/uvgroovy/findwords

RUN go install github.com/uvgroovy/findwords

WORKDIR /go/src/github.com/uvgroovy/findwords/

ENTRYPOINT ["/go/bin/findwords", "-server"]

EXPOSE 8080
