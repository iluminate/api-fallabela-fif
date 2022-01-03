FROM golang:1.17 as builder
RUN mkdir /resources
ADD . /resources/
WORKDIR /resources
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN go mod download
RUN go test ./test/...
RUN go build -a -installsuffix cgo -o runner cmd/api/server/main.go

FROM scratch
COPY --from=builder /resources/configs /configs
COPY --from=builder /resources/runner .
CMD ["./runner"]