#!/bin/bash

podman stop postgres
podman stop postgres_exporter
podman stop prometheus
podman stop grafana
podman rm postgres
podman rm postgres_exporter
podman rm prometheus
podman rm grafana
podman build -t postgres:latest -f Dockerfile-postgres .
podman build -t postgres_exporter:latest -f Dockerfile-postgres_exporter .
podman build -t prometheus:latest -f Dockerfile-prometheus .
podman build -t grafana:latest -f Dockerfile-grafana .
podman run --name postgres --ip=10.88.0.101 -p 5432:5432 -dt postgres
podman run --name postgres_exporter --ip=10.88.0.102 -p 9187:9187 -dt postgres_exporter
podman run --name prometheus --ip=10.88.0.103 -p 9090:9090 -dt prometheus
podman run --name grafana --ip=10.88.0.104 -p 3000:3000 -dt grafana

podman generate systemd --new --name postgres > /etc/systemd/system/postgres-podman.service
podman generate systemd --new --name postgres_exporter > /etc/systemd/system/postgres_exporter-podman.service
podman generate systemd --new --name prometheus > /etc/systemd/system/prometheus-podman.service
podman generate systemd --new --name grafana > /etc/systemd/system/grafana-podman.service
systemctl daemon-reload
systemctl enable postgres-podman.service
systemctl enable postgres_exporter-podman.service
systemctl enable prometheus-podman.service
systemctl enable grafana-podman.service
