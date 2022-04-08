#!/usr/bin/env bash

# *****************************************************************************

bash_cell SETUP "IMPORT PROVONE TRACE" << END_CELL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/branched-pipeline.jsonld

END_CELL

# *****************************************************************************

bash_cell S1 "EXPORT AS N-TRIPLES" << END_CELL

geist export --format nt | sort

END_CELL

# *****************************************************************************

bash_cell S2 "EXPORT AS JSON-LD" << END_CELL

geist export --format jsonld

END_CELL

# # *****************************************************************************
#
# bash_cell S2 "EXPORT AS TURTLE" << END_CELL
#
# geist export --format ttl
#
# END_CELL
#
#
# # *****************************************************************************
#
# bash_cell S4 "EXPORT AS RDF/XML" << END_CELL
#
# geist export --format xml | xmllint - --c14n11
#
# END_CELL
