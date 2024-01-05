# TODO Inject dependencies

FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o artifact .

FROM alpine:latest
COPY --from=builder /app/artifact /app/main

# k8s configmap
#COPY environment/config.yaml /environment/config.yaml

EXPOSE 8080
ENTRYPOINT ["/app/main"]