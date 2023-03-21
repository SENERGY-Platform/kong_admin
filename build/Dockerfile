FROM golang:1.18 AS builder
ENV GO111MODULE=on

COPY . /go/src/app
WORKDIR /go/src/app
RUN git log -1 --oneline > version.txt

WORKDIR /go/src/app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o ../app

FROM golang:1.18 AS prod
WORKDIR /root

COPY --from=builder /go/src/app/app .
COPY --from=builder /go/src/app/version.txt .

EXPOSE 8080

ENTRYPOINT [ "/root/app" ]