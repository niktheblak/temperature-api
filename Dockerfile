FROM golang:1.17 as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -o /go/bin/app

FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/app /
ENTRYPOINT ["/app"]
