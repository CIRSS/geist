#!/usr/bin/env bash

data_file='../data/address-book.jsonld'

arq_cell ${data_file} Q1 << END_QUERY

# List all triples in blazegraph.

PREFIX ab: <http://learningsparql.com/ns/addressbook#>
CONSTRUCT
{ ?s ?p ?o }
FROM <http://127.0.0.1:9999/blazegraph/sparql>
WHERE
{ ?s ?p ?o }
END_QUERY

arq_cell ${data_file} Q2 << END_QUERY

# Select all triples in blazegraph.

PREFIX ab: <http://learningsparql.com/ns/addressbook#>
SELECT ?s ?p ?o
WHERE
{
    SERVICE <http://127.0.0.1:9999/blazegraph/sparql>
    {
        SELECT ?s ?p ?o
        WHERE {}
    }
}
END_QUERY
