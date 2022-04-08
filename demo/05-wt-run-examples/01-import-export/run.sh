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
