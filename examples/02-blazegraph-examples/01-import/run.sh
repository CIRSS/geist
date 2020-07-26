#!/usr/bin/env bash

RUNNER='../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "INITIALIZE BLAZEGRAPH INSTANCE" << END_SCRIPT

blazegraph drop

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "IMPORT TWO TRIPLES AS N-TRIPLES" << END_SCRIPT

blazegraph import --format nt << END_DATA

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
            "http://tmcphill.net/tags#tag": "seven"
        },
        {
            "@id": "http://tmcphill.net/data#y",
            "http://tmcphill.net/tags#tag": "eight"
        }
    ]

END_DATA

blazegraph export --format nt

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S4 "IMPORT TWO TRIPLES AS RDF-XML" << END_SCRIPT

blazegraph drop
blazegraph import --format xml << END_DATA

    <rdf:RDF
        xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
        xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#"
        xmlns:sesame="http://www.openrdf.org/schema/sesame#"
        xmlns:owl="http://www.w3.org/2002/07/owl#"
        xmlns:xsd="http://www.w3.org/2001/XMLSchema#"
        xmlns:fn="http://www.w3.org/2005/xpath-functions#"
        xmlns:foaf="http://xmlns.com/foaf/0.1/"
        xmlns:dc="http://purl.org/dc/elements/1.1/"
        xmlns:hint="http://www.bigdata.com/queryHints#"
        xmlns:bd="http://www.bigdata.com/rdf#"
        xmlns:bds="http://www.bigdata.com/rdf/search#">

    <rdf:Description rdf:about="http://tmcphill.net/data#y">
        <tag xmlns="http://tmcphill.net/tags#">eight</tag>
    </rdf:Description>

    <rdf:Description rdf:about="http://tmcphill.net/tags#tag">
        <rdf:type rdf:resource="http://www.w3.org/1999/02/22-rdf-syntax-ns#Property"/>
        <rdfs:subPropertyOf rdf:resource="http://tmcphill.net/tags#tag"/>
    </rdf:Description>

    <rdf:Description rdf:about="http://tmcphill.net/data#x">
        <tag xmlns="http://tmcphill.net/tags#">seven</tag>
    </rdf:Description>

    </rdf:RDF>

END_DATA

blazegraph export --format nt

END_SCRIPT
