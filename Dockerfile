# Step 1 build executable binary
FROM golang:alpine AS builder

# Install git
COPY . $GOPATH/src/whitelist_ceria
WORKDIR $GOPATH/src/whitelist_ceria

#build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/app .
# RUN GO111MODULE=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/app .

# Step 2 build a small image
# start from scratch
FROM alpine:3.18

# Copy our static executable
COPY --from=builder /go/bin/app /go/bin/app
COPY kerberos /go/bin

RUN apk update
RUN apk add krb5-pkinit krb5-dev krb5
RUN cat go/bin/krb5.conf

ENTRYPOINT ["/go/bin/app"]