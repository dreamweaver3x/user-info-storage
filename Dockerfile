
FROM golang:1.16 as build

ENV GO111MODULE="on"
ENV CGO_ENABLED="0"
ENV GOOS="linux"

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /build/app.bin -a -ldflags "-s -w" cmd/main.go

# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static
WORKDIR /opt/service
COPY --from=build /build/app.bin .
EXPOSE 9090
CMD ["./app.bin"]