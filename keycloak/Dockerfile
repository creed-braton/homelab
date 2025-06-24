FROM quay.io/keycloak/keycloak:26.2.5 AS builder

ENV KC_HEALTH_ENABLED=true
ENV KC_METRICS_ENABLED=true
ENV KC_DB=postgres

RUN /opt/keycloak/bin/kc.sh build

FROM ubuntu:25.04 AS variable

RUN apt update && apt install -y gettext
COPY ./config.json ./config.json

ARG KC_URI
ENV KC_URI=$KC_URI
ARG GF_URI 
ENV GF_URI=$GF_URI
ARG GF_CLIENT_SECRET
ENV GF_CLIENT_SECRET=$GF_CLIENT_SECRET

# set environment variables explicitly, otherwise overwrites ${client_id} (breaks realm config)
RUN envsubst '${KC_URI} ${GF_URI} ${GF_CLIENT_SECRET}}' < config.json > updated-config.json

FROM quay.io/keycloak/keycloak:26.2.5

COPY --from=builder /opt/keycloak/ /opt/keycloak/
COPY --from=variable ./updated-config.json /opt/keycloak/data/import/config.json

ENTRYPOINT ["/opt/keycloak/bin/kc.sh"]
