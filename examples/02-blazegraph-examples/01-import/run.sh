#!/usr/bin/env bash

RUNNER='../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "INITIALIZE BLAZEGRAPH INSTANCE" << END_SCRIPT

blazegraph drop

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "IMPORT TWO TRIPLES AS N-TRIPLES" << END_SCRIPT

blazegraph import --format ttl << END_DATA

	<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .

END_DATA

blazegraph export --format nt

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S2 "IMPORT TWO TRIPLES AS TURTLE" << END_SCRIPT

blazegraph drop
blazegraph import --format ttl << END_DATA

	@prefix data: <http://tmcphill.net/data#> .
	@prefix tags: <http://tmcphill.net/tags#> .

	data:y tags:tag "eight" .
	data:x tags:tag "seven" .

END_DATA

blazegraph export --format nt

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S3 "IMPORT TWO TRIPLES AS JSON-LD" << END_SCRIPT

blazegraph drop
blazegraph import --format jsonld << END_DATA

    [
        {
            "@id": "http://tmcphill.net/data#x",
            "http://tmcphill.net/tags#tag": [ { "@value": "seven" } ]
        },
        {
            "@id": "http://tmcphill.net/data#y",
            "http://tmcphill.net/tags#tag": [ { "@value": "eight" } ]
        }
    ]

END_DATA

blazegraph export --format nt

END_SCRIPT
