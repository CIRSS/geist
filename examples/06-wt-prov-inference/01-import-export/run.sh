#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT PROVONE TRACE" << END_SCRIPT

blaze destroy --dataset kb --quiet
blaze create --dataset kb --quiet --infer owl
blaze import --file ../data/wt-prov-rules.ttl
blaze import --format jsonld --file ../data/branched-pipeline.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT AS N-TRIPLES" << END_SCRIPT

blaze export --format nt | sort

END_SCRIPT
