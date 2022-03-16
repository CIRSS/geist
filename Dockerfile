FROM docker.io/cirss/go-dev

COPY .repro .repro

USER repro

# install required repro modules
RUN repro.require geist exported --dev --demo
RUN repro.require blazegraph-service 0.2.6 ${CIRSS_RELEASE}

RUN repro.atstart start-blazegraph

RUN repro.setenv REPRO_BINARY_SHORT_NAME geist

CMD  /bin/bash -il

