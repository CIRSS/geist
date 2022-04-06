#!/usr/bin/env bash

data_file='../data/address-book.jsonld'

arq_cell ${data_file} Q1 "List all triples in blazegraph." << END_QUERY
PREFIX ab: <http://learningsparql.com/ns/addressbook#>
CONSTRUCT
{ ?s ?p ?o }
FROM <http://127.0.0.1:9999/blazegraph/sparql>
WHERE
{ ?s ?p ?o }
END_QUERY

arq_cell ${data_file} Q2 "Select all triples in blazegraph." << END_QUERY
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
