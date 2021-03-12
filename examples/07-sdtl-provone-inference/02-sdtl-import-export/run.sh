#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} DUMP-1 "IMPORT SDTL AS JSON-LD AND EXPORT AS N-TRIPLES" << END_SCRIPT

blazegraph destroy --dataset kb --quiet
blazegraph create --dataset kb --quiet
blazegraph import --format jsonld --file ../data/compute-sdtl.jsonld
blazegraph export --format nt | sort

END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} DUMP-2 "IMPORT SDTL OWL FILE AND EXPORT AS N-TRIPLES" << END_SCRIPT

# blazegraph destroy --dataset kb --quiet
# blazegraph create --dataset kb --quiet
# blazegraph import --format xml --file ../data/sdtl.owl
# blazegraph export --format nt

# END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} DUMP-3 "EXPORT SDTL OWL AS TURTLE" << END_SCRIPT

# blazegraph destroy --dataset kb --quiet
# blazegraph create --dataset kb --quiet
# blazegraph import --format xml --file ../data/sdtl.owl
# blazegraph export --format ttl

# END_SCRIPT

# *****************************************************************************

# bash ${RUNNER} DUMP-3 "IMPORT TOMMY'S JSON-LD FILE AND EXPORT AS N-TRIPLES" << END_SCRIPT

# blazegraph destroy --dataset kb --quiet
# blazegraph create --dataset kb --quiet
# blazegraph import --format jsonld --file ../data/rdf.jsonld
# blazegraph export --format nt | sort

# END_SCRIPT
