FROM --platform=${BUILDPLATFORM} alpine:3.17.2

ARG TARGETOS
ARG TARGETARCH
ARG TARGETPLATFORM

LABEL maintainer="Gareth Evans <gareth@bryncynfelin.co.uk>"
COPY dist/scv-${TARGETOS}_${TARGETOS}_${TARGETARCH}/scv /usr/bin/scv

ENTRYPOINT [ "/usr/bin/scv" ]

CMD ["--help"]
