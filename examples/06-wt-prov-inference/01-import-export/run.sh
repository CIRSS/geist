#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT PROVONE TRACE" << END_SCRIPT

blazegraph destroy --dataset kb
blazegraph create --dataset kb --quiet --infer owl
blazegraph import --file ../data/wt-prov-rules.ttl
blazegraph import --format jsonld --file ../data/branched-pipeline.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT AS N-TRIPLES" << END_SCRIPT

blazegraph export --format nt | sort

END_SCRIPT
