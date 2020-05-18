FROM golang:1.13

WORKDIR /go/src/si-convert
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["si-convert"]