#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "INITIALIZE BLAZEGRAPH INSTANCE" << END_SCRIPT

blazegraph destroy --dataset kb
blazegraph create --dataset kb

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "IMPORT TWO TRIPLES AS N-TRIPLES" << END_SCRIPT

blazegraph import --format nt | sort << END_DATA

	<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .

END_DATA

blazegraph export --format nt --sort=true

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S2 "IMPORT TWO TRIPLES AS TURTLE" << END_SCRIPT

blazegraph destroy --dataset kb
blazegraph create --dataset kb

blazegraph import --format ttl << END_DATA

	@prefix data: <http://tmcphill.net/data#> .
	@prefix tags: <http://tmcphill.net/tags#> .

	data:y tags:tag "eight" .
	data:x tags:tag "seven" .

END_DATA

blazegraph export --format nt --sort=true

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S3 "IMPORT TWO TRIPLES AS JSON-LD" << END_SCRIPT

blazegraph destroy --dataset kb
blazegraph create --dataset kb

blazegraph import --format jsonld << END_DATA

    [
        {
            "@id": "http://tmcphill.net/data#x",
            "http://tmcphill.net/tags#tag": "seven"
        },
        {
            "@id": "http://tmcphill.net/data#y",
            "http://tmcphill.net/tags#tag": "eight"
        }
    ]

END_DATA

blazegraph export --format nt --sort=true

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S4 "IMPORT TWO TRIPLES AS RDF-XML" << END_SCRIPT

blazegraph destroy --dataset kb
blazegraph create --dataset kb

blazegraph import --format xml << END_DATA

    <rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">

    <rdf:Description rdf:about="http://tmcphill.net/data#y">
        <tag xmlns="http://tmcphill.net/tags#">eight</tag>
    </rdf:Description>

    <rdf:Description rdf:about="http://tmcphill.net/data#x">
        <tag xmlns="http://tmcphill.net/tags#">seven</tag>
    </rdf:Description>

    </rdf:RDF>

END_DATA

blazegraph export --format nt --sort=true

END_SCRIPT
