FROM golang:latest  AS build-env

# Dependencies
WORKDIR /build
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 go build -ldflags '-w -s' -o /main

# Build runtime container
FROM scratch
COPY --from=build-env /main /main
EXPOSE 80
CMD ["/main"]