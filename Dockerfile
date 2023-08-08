FROM golang:1.21 as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -o /go/bin/app

FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/app /
ENTRYPOINT ["/app", "server"]
