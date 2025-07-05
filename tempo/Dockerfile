FROM ubuntu:25.04 AS variable

RUN apt update && apt install -y gettext
COPY ./config.yaml ./config.yaml

ARG ACCESS_KEY
ENV ACCESS_KEY=$ACCESS_KEY
ARG SECRET_KEY
ENV SECRET_KEY=$SECRET_KEY

RUN envsubst < config.yaml > updated-config.yaml

FROM grafana/tempo:main-3349f00

COPY --from=variable ./updated-config.yaml /etc/config.yaml
