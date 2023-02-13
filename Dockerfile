FROM cirss/repro-parent:latest

# copy exports into new Docker image
COPY exports /repro/exports

# copy the repro boot setup script from the distribution and run it
ADD ${REPRO_DIST}/boot-setup /repro/dist/
RUN bash /repro/dist/boot-setup

USER repro

# install required external repro modules
RUN repro.require blaze 0.2.7 ${CIRSS_RELEASE}
RUN repro.require blazegraph-service master ${CIRSS}

# install contents of the exports directory as a repro module
RUN repro.require geist exports --code --demos

CMD  /bin/bash -il

