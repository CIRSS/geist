#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT PROVONE TRACE" << END_SCRIPT

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet --infer owl
geist import --file ../data/wt-prov-rules.ttl
geist import --format jsonld --file ../data/branched-pipeline.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT AS N-TRIPLES WITHOUT INFERENCE" << END_SCRIPT

geist export --format nt --includeinferred=false | sort

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S2 "EXPORT AS N-TRIPLES WITH INFERENCE" << END_SCRIPT

geist export --format nt --includeinferred=true | sort

END_SCRIPT
