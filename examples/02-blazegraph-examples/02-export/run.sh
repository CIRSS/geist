#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "INITIALIZE BLAZEGRAPH INSTANCE WITH ADDRESS BOOK" << END_SCRIPT

blazegraph destroy --dataset kb --quiet
blazegraph create --dataset kb --quiet
blazegraph import --file ../data/address-book.jsonld --format jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT ADDRESS BOOK AS JSON-LD" << END_SCRIPT

blazegraph export --format jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT ADDRESS BOOK AS TURTLE" << END_SCRIPT

blazegraph export --format ttl

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT ADDRESS BOOK AS N-TRIPLES" << END_SCRIPT

blazegraph export --format nt | sort

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "EXPORT ADDRESS BOOK AS RDF-XML" << END_SCRIPT

blazegraph export --format xml

END_SCRIPT
