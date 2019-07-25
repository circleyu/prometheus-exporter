FROM golang:latest AS buildStage
WORKDIR /go/src/prometheus-exporter-training
COPY . .
RUN CGO_ENABLED=0 go build

FROM scratch
WORKDIR /app
COPY --from=buildStage /go/src/prometheus-exporter-training/prometheus-exporter-training .
EXPOSE 9001
ENTRYPOINT ["/app/prometheus-exporter-training"]