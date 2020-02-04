FROM grafana/grafana
RUN grafana-cli plugins install vertamedia-clickhouse-datasource
ADD ./provisioning /etc/grafana/provisioning
ADD ./config.ini /etc/grafana/config.ini
ADD ./dashboards /var/lib/grafana/dashboards