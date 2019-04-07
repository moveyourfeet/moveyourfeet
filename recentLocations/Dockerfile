FROM golang:latest  AS build-env

RUN go get -u github.com/golang/dep/cmd/dep && \
    mkdir -p /go/src/github.com/georace/recentLocations 

WORKDIR /go/src/github.com/georace/recentLocations

COPY Gopkg.lock Gopkg.toml ./

RUN dep ensure -vendor-only

COPY . ./

RUN CGO_ENABLED=0 go build -ldflags '-w -s' -o /main

FROM alpine

COPY --from=build-env /main /main
EXPOSE 80

CMD ["/main"]