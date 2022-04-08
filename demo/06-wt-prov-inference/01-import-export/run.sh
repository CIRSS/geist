#!/usr/bin/env bash

# *****************************************************************************

bash_cell SETUP "IMPORT PROVONE TRACE" << END_CELL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet --infer owl
geist import --file ../data/wt-prov-rules.ttl
geist import --format jsonld --file ../data/branched-pipeline.jsonld

END_CELL

# *****************************************************************************

bash_cell S1 "EXPORT AS N-TRIPLES WITHOUT INFERENCE" << END_CELL

geist export --format nt --includeinferred=false | sort

END_CELL

# *****************************************************************************

bash_cell S2 "EXPORT AS N-TRIPLES WITH INFERENCE" << END_CELL

geist export --format nt --includeinferred=true | sort

END_CELL
