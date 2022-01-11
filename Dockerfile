FROM golang:1.17

WORKDIR /go/src/app


COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

RUN chmod +x ./main

CMD ["./main"]