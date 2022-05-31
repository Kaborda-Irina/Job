FROM golang:1.18-alpine AS buildenv
WORKDIR /src
ADD . /src
RUN go mod download
#RUN go build -o sha256sum cmd/main.go
RUN go build -o appCron app_check/main.go

RUN chmod +x appCron

FROM alpine:latest
WORKDIR /app

COPY --from=buildenv /src/appCron .
COPY --from=buildenv /src/.env .
VOLUME /app

EXPOSE 9090

ENTRYPOINT ["/app/appCron"]