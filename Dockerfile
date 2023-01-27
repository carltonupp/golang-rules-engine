FROM golang:alpine

WORKDIR /app

COPY . ./

RUN go get -d -v ./...

RUN go build -o /rules_engine

EXPOSE 8080

ENTRYPOINT [ "/rules_engine" ]
