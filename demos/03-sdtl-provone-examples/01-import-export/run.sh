#!/usr/bin/env bash

# *****************************************************************************

bash_cell SETUP << END_CELL

# IMPORT SDTL-PROVONE TRACE

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/single-command.jsonld

END_CELL

# *****************************************************************************

bash_cell S1 << END_CELL

# EXPORT AS N-TRIPLES

geist export --format nt | sort

END_CELL

# *****************************************************************************

bash_cell S2 << END_CELL

# EXPORT AS JSON-LD

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
