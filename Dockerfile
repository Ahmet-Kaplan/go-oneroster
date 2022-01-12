FROM golang:latest as builder
WORKDIR /go/src/usulroster
ADD . /go/src/usulroster

RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /go/bin/usulroster

FROM gcr.io/distroless/static
COPY --from=builder /go/bin/usulroster /
EXPOSE 80
CMD ["/usulroster"]
