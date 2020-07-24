#!/usr/bin/env bash

 runner='../common/run_example.sh'
# data_file='../data/address-book.jsonld'

# cat ${data_file}

bash ${runner} S1 "Export address book as JSON-LD" << END_SCRIPT
blazegraph drop
blazegraph import --file ../data/address-book.jsonld --format jsonld
blazegraph export --format jsonld
END_SCRIPT
