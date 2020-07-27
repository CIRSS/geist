#!/usr/bin/env bash

RUNNER='../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT SDTL-PROVONE TRACE" << END_SCRIPT

blazegraph drop
blazegraph import --format jsonld --file ../data/single-command.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT AS N-TRIPLES" << END_SCRIPT

blazegraph export --format nt | sort

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S2 "EXPORT AS TURTLE" << END_SCRIPT

blazegraph export --format ttl

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S3 "EXPORT AS JSON-LD" << END_SCRIPT

blazegraph export --format jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S4 "EXPORT AS RDF/XML" << END_SCRIPT

blazegraph export --format xml | xmllint - --c14n11

END_SCRIPT
