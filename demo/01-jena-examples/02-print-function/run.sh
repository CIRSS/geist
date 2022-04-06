#!/usr/bin/env bash

data_file='../data/address-book.jsonld'

cat ${data_file}

arq_cell ${data_file} Q1 "What is everyone's email addresses (and print them too)?" << END_QUERY
PREFIX afn: <http://jena.apache.org/ARQ/function#>
PREFIX ab: <http://learningsparql.com/ns/addressbook#>
SELECT ?firstname ?email
WHERE
{
    ?person ab:firstname ?firstname .
    ?person ab:email     ?email     .
    FILTER(CONTAINS(?firstname, "i"))
    FILTER(afn:print(?email))
}
END_QUERY

arq_cell ${data_file} Q2 "What is everyone's email addresses (and print them in subqueries)?" << END_QUERY
PREFIX afn: <http://jena.apache.org/ARQ/function#>
PREFIX ab: <http://learningsparql.com/ns/addressbook#>
SELECT *
WHERE {
    {
        SELECT ?firstname
        WHERE { 
            ?person ab:firstname ?firstname
            FILTER(afn:print(?firstname))
        }
    }
    {
        SELECT ?email
        WHERE { 
            ?person ab:email ?email
            FILTER(afn:print(?email))
        }
    }
}
END_QUERY

arq_cell ${data_file} Q3 "Construct new triples for everyone's name and email addresses (and print them in subqueries)?" << END_QUERY
PREFIX afn: <http://jena.apache.org/ARQ/function#>
PREFIX ab: <http://learningsparql.com/ns/addressbook#>
PREFIX tm: <http://learningsparql.com/tm#>
CONSTRUCT {
    ?person tm:firstname ?firstname . 
    ?person tm:email ?email .
}
WHERE {
    ?person ab:email ?email.
    ?person ab:firstname ?firstname .
    FILTER(afn:print(?firstname))
    FILTER(afn:print(?email))
}
END_QUERY


