FROM golang:1.17-alpine3.15
WORKDIR /go/src/github.com/edfan0930/aha
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o aha 
EXPOSE 8080
ENTRYPOINT [ "./aha" ]