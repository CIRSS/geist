#!/usr/bin/env bash

# *****************************************************************************

run_cell SETUP "INITIALIZE BLAZEGRAPH INSTANCE" << END_CELL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet

END_CELL

# *****************************************************************************

run_cell S1 "IMPORT TWO TRIPLES AS N-TRIPLES" << END_CELL

geist import --format nt | sort << END_DATA

	<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .

END_DATA

geist export --format nt --sort=true

END_CELL

# *****************************************************************************

run_cell S2 "IMPORT TWO TRIPLES AS TURTLE" << END_CELL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet

geist import --format ttl << END_DATA

	@prefix data: <http://tmcphill.net/data#> .
	@prefix tags: <http://tmcphill.net/tags#> .

	data:y tags:tag "eight" .
	data:x tags:tag "seven" .

END_DATA

geist export --format nt --sort=true

END_CELL

# *****************************************************************************

run_cell S3 "IMPORT TWO TRIPLES AS JSON-LD" << END_CELL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet

geist import --format jsonld << END_DATA

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

geist export --format nt --sort=true

END_CELL

# *****************************************************************************

run_cell S4 "IMPORT TWO TRIPLES AS RDF-XML" << END_CELL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet

geist import --format xml << END_DATA

    <rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">

    <rdf:Description rdf:about="http://tmcphill.net/data#y">
        <tag xmlns="http://tmcphill.net/tags#">eight</tag>
    </rdf:Description>

    <rdf:Description rdf:about="http://tmcphill.net/data#x">
        <tag xmlns="http://tmcphill.net/tags#">seven</tag>
    </rdf:Description>

    </rdf:RDF>

END_DATA

geist export --format nt --sort=true

END_CELL
