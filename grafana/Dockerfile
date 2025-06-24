FROM ubuntu:25.04 AS variable

RUN apt update && apt install -y gettext
COPY ./grafana.ini ./grafana.ini

ARG DB_PASS
ENV DB_PASS=$DB_PASS
ARG GF_URI 
ENV GF_URI=$GF_URI
ARG KC_URI
ENV KC_URI=$KC_URI
ARG GF_CLIENT_SECRET
ENV GF_CLIENT_SECRET=$GF_CLIENT_SECRET

RUN envsubst < grafana.ini > updated-grafana.ini

FROM grafana/grafana:12.0.2

COPY --from=variable ./updated-grafana.ini /etc/grafana/grafana.ini
COPY ./datasource.yaml /etc/grafana/provisioning/datasources/datasource.yaml
COPY ./dashboard.yaml /etc/grafana/provisioning/dashboards/dashboard.yaml
