FROM golang:1.25.6

WORKDIR /build
COPY . ./

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o warehouse app/main.go

RUN chmod +x warehouse

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=0 /build/warehouse ./

