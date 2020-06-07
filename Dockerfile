FROM golang:latest

RUN mkdir -p /go/src/github.com/dennisssdev/yoink

WORKDIR /go/src/github.com/dennisssdev/yoink

COPY . /go/src/github.com/dennisssdev/yoink

RUN go install ./cmd/epicstorecommunicatorserver

CMD /go/bin/epicstorecommunicatorserver

EXPOSE 8034