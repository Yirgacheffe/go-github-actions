# STEP 1
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

ENV GOPATH=/go
WORKDIR $GOPATH/src/olint-go

COPY . .
RUN go get -d -v

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build ./cmd/general/main.go -ldflags="-w -s" -o $GOPATH/bin/olint-go

# STEP 2
FROM scratch

COPY --from=builder /go/bin/olint-go /opt/app/olint-go
ENTRYPOINT ["/opt/app/olint-go"]
