#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT SDTL-PROVONE TRACE" << END_SCRIPT

blaze destroy --dataset kb --quiet
blaze create --dataset kb --quiet
blaze import --format jsonld --file ../data/single-command.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT AS N-TRIPLES" << END_SCRIPT

blaze export --format nt | sort

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S2 "EXPORT AS JSON-LD" << END_SCRIPT

blaze export --format jsonld

END_SCRIPT

# # *****************************************************************************
#
# bash ${RUNNER} S2 "EXPORT AS TURTLE" << END_SCRIPT
#
# blaze export --format ttl
#
# END_SCRIPT
#
#
# # *****************************************************************************
#
# bash ${RUNNER} S4 "EXPORT AS RDF/XML" << END_SCRIPT
#
# blaze export --format xml | xmllint - --c14n11
#
# END_SCRIPT
