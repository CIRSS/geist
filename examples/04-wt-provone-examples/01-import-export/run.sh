#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT PROVONE TRACE" << END_SCRIPT

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/branched-pipeline.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT AS N-TRIPLES" << END_SCRIPT

geist export --format nt | sort

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S2 "EXPORT AS JSON-LD" << END_SCRIPT

geist export --format jsonld

END_SCRIPT

# # *****************************************************************************
#
# bash ${RUNNER} S2 "EXPORT AS TURTLE" << END_SCRIPT
#
# geist export --format ttl
#
# END_SCRIPT
#
#
# # *****************************************************************************
#
# bash ${RUNNER} S4 "EXPORT AS RDF/XML" << END_SCRIPT
#
# geist export --format xml | xmllint - --c14n11
#
# END_SCRIPT
