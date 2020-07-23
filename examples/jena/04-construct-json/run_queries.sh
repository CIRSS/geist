#!/usr/bin/env bash

run_query='../common/run_query.sh'
data_file='../data/address-book.jsonld'

cat ${data_file}

bash ${run_query} ${data_file} Q1 "Construct new triples for everyone's name and email addresses as JSON" << END_QUERY
PREFIX afn: <http://jena.apache.org/ARQ/function#>
PREFIX ab: <http://learningsparql.com/ns/addressbook#>
JSON {
    "firstname": ?firstname, 
    "email": ?email
}
WHERE {
    {
        SELECT ?firstname
        WHERE { 
            ?person ab:firstname ?firstname
        }
    }
    {
        SELECT ?email
        WHERE { 
            ?person ab:email ?email
        }
    }
}

END_QUERY


