FROM golang:1.19

RUN export GOROOT="/public"

WORKDIR /public

COPY ./app ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]