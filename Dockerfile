FROM golang:1.17-alpine AS build

ENV CGO_ENABLED=0

COPY cn-camp-work1 /go/src/project
WORKDIR /go/src/project/

RUN GOOS=linux go build -o /bin/httpserver

FROM scratch
COPY --from=build /bin/httpserver /bin/httpserver
ENTRYPOINT ["/bin/httpserver"]
