TARGETS_DIR=.repro/Makefile.targets

default_target: help

##
## # Aliases for targets in this Makefile.
##

include repro.config
include ${TARGETS_DIR}/Makefile.setup
include ${TARGETS_DIR}/Makefile.examples
include ${TARGETS_DIR}/Makefile.code
include ${TARGETS_DIR}/Makefile.image
include ${TARGETS_DIR}/Makefile.docker
include ${TARGETS_DIR}/Makefile.help

