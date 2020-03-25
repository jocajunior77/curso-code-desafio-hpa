FROM golang:alpine

WORKDIR /src/app
RUN apk add git

COPY ./src/app.go .

RUN go get -d -v \
 && GOOS=linux go build -ldflags="-s -w" app.go \
 && chmod +x app

ENTRYPOINT ["./app"]