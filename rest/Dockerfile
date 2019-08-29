FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /gorest .

FROM scratch
COPY --from=builder /gorest /gorest
COPY --from=builder /etc/passwd /etc/passwd

EXPOSE 8080

USER appuser
ENTRYPOINT ["/gorest"]
