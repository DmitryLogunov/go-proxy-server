FROM golang:alpine AS builder

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app

RUN apk --no-cache add git
RUN go get -u github.com/justinas/alice
RUN go get -u github.com/dgrijalva/jwt-go
RUN go get -u github.com/fatih/color
RUN go get -u gopkg.in/yaml.v2
RUN go build -o proxy-server

FROM alpine

WORKDIR /usr/src/app

COPY --from=builder /go/src/app/proxy-server .

CMD ["./proxy-server"]

EXPOSE 80
