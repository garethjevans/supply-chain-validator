FROM --platform=${BUILDPLATFORM} ubuntu:23.10
LABEL maintainer="Gareth Evans <gareth@bryncynfelin.co.uk>"

ARG TARGETOS
ARG TARGETARCH
ARG TARGETPLATFORM

RUN apt-get update && apt install wget -y

ENV YTT_VERSION 0.40.4
RUN wget https://github.com/carvel-dev/ytt/releases/download/v${YTT_VERSION}/ytt-${TARGETOS}-${TARGETARCH} && \
    mv ytt-${TARGETOS}-${TARGETARCH} /usr/bin/ytt && \
    chmod a+x /usr/bin/ytt

COPY dist/scv-${TARGETOS}_${TARGETOS}_${TARGETARCH}_v1/scv /usr/bin/scv

ENTRYPOINT [ "/usr/bin/scv" ]

CMD ["--help"]
