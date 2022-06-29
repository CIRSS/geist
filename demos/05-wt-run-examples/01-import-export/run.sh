#!/usr/bin/env bash

# *****************************************************************************

bash_cell SETUP << END_CELL

# IMPORT PROVONE TRACE

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/branched-pipeline.jsonld

END_CELL

# *****************************************************************************

bash_cell S1 << END_CELL

# EXPORT AS N-TRIPLES

geist export --format nt | sort

END_CELL
