FROM golang as builder

WORKDIR /go/src/computer-api

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 1337

CMD ["computer-api"]