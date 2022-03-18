FROM docker.io/cirss/repro-template

COPY .repro .repro

USER repro

# install required repro modules
RUN repro.require geist exported --dev --demo
RUN repro.require blazegraph-service 0.2.6 ${CIRSS_RELEASE}

RUN repro.atstart start-blazegraph

CMD  /bin/bash -il

