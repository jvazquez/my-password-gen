FROM golang:1.15  as build
RUN mkdir -p /go/src/github.com/my-password-gen
WORKDIR /go/src/github.com/my-password-gen
COPY . .
RUN go get -v all

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo \
    -o /go/bin/generate-password \
    /go/src/github.com/my-password-gen

FROM scratch
COPY --from=build /go/bin/my-password-gen /my-password-gen
ENTRYPOINT ["/my-password-gen"]
