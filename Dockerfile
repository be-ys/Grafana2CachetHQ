FROM golang:alpine AS builder

COPY . /go/src/cachetconnector

WORKDIR /go/src/cachetconnector

#RUN go get
ENV GO111MODULE=off
RUN apk add --no-cache git

#RUN go mod init
WORKDIR /go/src/cachetconnector/src
ENV GOPATH="/go/src/cachetconnector"

RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/zhashkevych/scheduler

RUN go build -o /go/bin/cachetconnector .


FROM alpine AS run

RUN mkdir /app/
COPY --from=builder /go/bin/cachetconnector /app/cachetconnector

WORKDIR /app
ENTRYPOINT ["./cachetconnector"]

EXPOSE 8080
