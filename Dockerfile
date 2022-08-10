FROM cirss/repro-parent:latest

COPY exports /repro/exports

ADD ${REPRO_DIST}/setup /repro/dist/
RUN bash /repro/dist/setup

USER repro

RUN repro.require geist exports --code --demos
RUN repro.require blaze 0.2.6 ${CIRSS_RELEASE}
RUN repro.require blazegraph-service master ${CIRSS}

CMD  /bin/bash -il

