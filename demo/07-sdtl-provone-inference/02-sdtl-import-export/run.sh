#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} DUMP-1 "IMPORT SDTL AS JSON-LD AND EXPORT AS N-TRIPLES" << END_SCRIPT

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/compute-sdtl.jsonld
geist export --format nt | sort

END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} DUMP-2 "IMPORT SDTL OWL FILE AND EXPORT AS N-TRIPLES" << END_SCRIPT

# geist destroy --dataset kb --quiet
# geist create --dataset kb --quiet
# geist import --format xml --file ../data/sdtl.owl
# geist export --format nt

# END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} DUMP-3 "EXPORT SDTL OWL AS TURTLE" << END_SCRIPT

# geist destroy --dataset kb --quiet
# geist create --dataset kb --quiet
# geist import --format xml --file ../data/sdtl.owl
# geist export --format ttl

# END_SCRIPT

# *****************************************************************************

# bash ${RUNNER} DUMP-3 "IMPORT TOMMY'S JSON-LD FILE AND EXPORT AS N-TRIPLES" << END_SCRIPT

# geist destroy --dataset kb --quiet
# geist create --dataset kb --quiet
# geist import --format jsonld --file ../data/rdf.jsonld
# geist export --format nt | sort

# END_SCRIPT
