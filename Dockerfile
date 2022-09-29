FROM cirss/repro-parent:latest

COPY exports /repro/exports

ADD ${REPRO_DIST}/boot-setup /repro/dist/
RUN bash /repro/dist/boot-setup

USER repro

RUN repro.require geist exports --code --demos
RUN repro.require blaze 0.2.7 ${CIRSS_RELEASE}
RUN repro.require blazegraph-service master ${CIRSS}

CMD  /bin/bash -il

