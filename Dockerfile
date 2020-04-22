FROM golang:1.10-stretch as builder
RUN mkdir /mnt/workspace
COPY . /mnt/workspace
WORKDIR /mnt/workspace
RUN go get -v -u github.com/gorilla/mux
RUN CGO_ENABLED=0 go build -o Provider .

FROM alpine:3.8
RUN apk -U add ca-certificates
EXPOSE 80
COPY --from=builder /mnt/workspace/Provider /
CMD ["/Provider"]
