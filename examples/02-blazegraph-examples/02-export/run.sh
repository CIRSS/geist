#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "INITIALIZE BLAZEGRAPH INSTANCE WITH ADDRESS BOOK" << END_SCRIPT

blaze destroy --dataset kb --quiet
blaze create --dataset kb --quiet
blaze import --file ../data/address-book.jsonld --format jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT ADDRESS BOOK AS JSON-LD" << END_SCRIPT

blaze export --format jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT ADDRESS BOOK AS TURTLE" << END_SCRIPT

blaze export --format ttl

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT ADDRESS BOOK AS N-TRIPLES" << END_SCRIPT

blaze export --format nt | sort

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT ADDRESS BOOK AS RDF-XML" << END_SCRIPT

blaze export --format xml

END_SCRIPT
