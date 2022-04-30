FROM golang:1.17-alpine AS build

COPY cn-camp-work1 /go/src/project
WORKDIR /go/src/project/

RUN go build -o /bin/httpserver

ENTRYPOINT ["/bin/httpserver"]
