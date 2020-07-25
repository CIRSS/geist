#!/usr/bin/env bash

RUNNER='../common/run_script_example.sh'

bash ${RUNNER} SETUP "INITIALIZE BLAZEGRAPH INSTANCE WITH ADDRESS BOOK" << END_SCRIPT
blazegraph drop
blazegraph import --file ../data/address-book.jsonld --format jsonld
END_SCRIPT

bash ${RUNNER} S1 "EXPORT ADDRESS BOOK AS JSON-LD" << END_SCRIPT
blazegraph export --format jsonld
END_SCRIPT
