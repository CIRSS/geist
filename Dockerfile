FROM docker.io/cirss/repro-template

COPY exports /repro/exports

USER repro

RUN repro.require geist exports --code --demo
RUN repro.require blaze 0.2.6 ${CIRSS_RELEASE}
RUN repro.require blazegraph-service master ${CIRSS_BRANCH}

RUN repro.atstart blazegraph-service.start

CMD  /bin/bash -il

