# Build GO Project
FROM golang:1.14.0-alpine
ENV GO111MODULE=on
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod ./
RUN go mod download
RUN go clean --modcache
COPY . .
RUN go build -o main

# Reduce Size without GO Image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["/app/main"]