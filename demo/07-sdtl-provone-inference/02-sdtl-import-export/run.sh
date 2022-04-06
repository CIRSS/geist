#!/usr/bin/env bash

# *****************************************************************************

run_cell DUMP-1 "IMPORT SDTL AS JSON-LD AND EXPORT AS N-TRIPLES" << END_CELL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/compute-sdtl.jsonld
geist export --format nt | sort

END_CELL

# # *****************************************************************************

# run_cell DUMP-2 "IMPORT SDTL OWL FILE AND EXPORT AS N-TRIPLES" << END_CELL

# geist destroy --dataset kb --quiet
# geist create --dataset kb --quiet
# geist import --format xml --file ../data/sdtl.owl
# geist export --format nt

# END_CELL

# # *****************************************************************************

# run_cell DUMP-3 "EXPORT SDTL OWL AS TURTLE" << END_CELL

# geist destroy --dataset kb --quiet
# geist create --dataset kb --quiet
# geist import --format xml --file ../data/sdtl.owl
# geist export --format ttl

# END_CELL

# *****************************************************************************

# run_cell DUMP-3 "IMPORT TOMMY'S JSON-LD FILE AND EXPORT AS N-TRIPLES" << END_CELL

# geist destroy --dataset kb --quiet
# geist create --dataset kb --quiet
# geist import --format jsonld --file ../data/rdf.jsonld
# geist export --format nt | sort

# END_CELL
