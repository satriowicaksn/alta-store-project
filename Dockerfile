FROM golang:1.16.6-alpine as debug
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main
EXPOSE $PORT
CMD ["/app/main"]

