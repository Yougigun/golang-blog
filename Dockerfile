FROM golang:1.16-alpine

WORKDIR /app
COPY . .
RUN go build -o /golang-blog
EXPOSE 3000
ENTRYPOINT ["/golang-blog"]