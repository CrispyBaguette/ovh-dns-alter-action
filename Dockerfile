FROM golang:1.16 as build-env

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /app

RUN go build -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /
CMD ["/app"]