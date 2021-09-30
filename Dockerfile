FROM alpine:3.9
WORKDIR /app

COPY k8cronjob /app/
ENTRYPOINT ["./k8cronjob"]
