FROM docker.io/cirss/repro-template

COPY exports /repro/exports

USER repro

RUN repro.require geist exports --dev --demo
RUN repro.require blazegraph-service master ${CIRSS_BRANCH}

RUN repro.atstart blazegraph-service.start

CMD  /bin/bash -il

