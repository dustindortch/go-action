FROM golang:1.24 AS builder

RUN apt-get update && apt-get -y install upx

ENV GO111MODULE=on CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.mod
COPY main.go main.go

RUN go get
RUN go build \
  -a \
  -trimpath \
  -ldflags "-s -w -extldflags '-static'" \
  -installsuffix cgo \
  -o build/action
RUN strip build/action
RUN upx -q -9 build/action

FROM scratch

COPY --from=builder /app/build/action /bin/action

ENTRYPOINT ["/bin/action"]