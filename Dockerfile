FROM cirss/repro-parent:latest

COPY exports /repro/exports

ADD ${REPRO_DIST}/setup /repro/dist/
RUN bash /repro/dist/setup

USER repro

RUN repro.require repro master ${REPROS_DEV}
RUN repro.require geist exports --code --demo
RUN repro.require blaze 0.2.6 ${CIRSS_RELEASE}
RUN repro.require blazegraph-service master ${CIRSS}

RUN repro.atstart blazegraph-service.start

CMD  /bin/bash -il

