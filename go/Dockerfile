FROM golang:latest
WORKDIR /go/src/github.com/patrickshan/knative-sample/go/
COPY . .
RUN CGO_ENABLED=0 go build -v -o app

FROM scratch
COPY --from=0 /go/src/github.com/patrickshan/knative-sample/go/app .
ENTRYPOINT ["/app"]
