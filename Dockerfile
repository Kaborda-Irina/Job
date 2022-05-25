FROM golang:1.18-alpine AS buildenv
WORKDIR /src
ADD . /src
RUN go mod download
#RUN go build -o sha256sum cmd/main.go
RUN go build -o sha256sum main.go

RUN chmod +x sha256sum

FROM alpine:latest
WORKDIR /app

COPY --from=buildenv /src/sha256sum .
COPY --from=buildenv /src/.env .
VOLUME /app

EXPOSE 9090

ENTRYPOINT ["/app/sha256sum"]
#CMD ["sh","-c","ls /"]
#CMD ["./sha256sum"]