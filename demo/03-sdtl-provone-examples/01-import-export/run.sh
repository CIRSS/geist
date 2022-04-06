#!/usr/bin/env bash

# *****************************************************************************

run_cell SETUP "IMPORT SDTL-PROVONE TRACE" << END_CELL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/single-command.jsonld

END_CELL

# *****************************************************************************

run_cell S1 "EXPORT AS N-TRIPLES" << END_CELL

geist export --format nt | sort

END_CELL

# *****************************************************************************

run_cell S2 "EXPORT AS JSON-LD" << END_CELL

geist export --format jsonld

END_CELL

# # *****************************************************************************
#
# run_cell S2 "EXPORT AS TURTLE" << END_CELL
#
# geist export --format ttl
#
# END_CELL
#
#
# # *****************************************************************************
#
# run_cell S4 "EXPORT AS RDF/XML" << END_CELL
#
# geist export --format xml | xmllint - --c14n11
#
# END_CELL
