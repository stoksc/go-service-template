FROM golang:1.13.0-stretch AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make bin

WORKDIR /dist

RUN cp /build/bin/* .


FROM debian:stable

COPY --chown=0:0 --from=builder /dist /

ENTRYPOINT ["/hello"]