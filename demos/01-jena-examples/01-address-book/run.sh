#!/usr/bin/env bash

data_file='../data/address-book.jsonld'

cat ${data_file}

arq_cell ${data_file} Q1 "What are Craig's email addresses?" << END_QUERY
PREFIX ab: <http://learningsparql.com/ns/addressbook#>
SELECT ?email
WHERE
{
    ?person ab:firstname "Craig"    .
    ?person ab:email     ?email     .
}
END_QUERY

arq_cell ${data_file} Q2 "Whose telephone number is (245) 646-5488?" << END_QUERY
PREFIX ab: <http://learningsparql.com/ns/addressbook#> 
SELECT ?name
WHERE
{
    ?person ab:homeTel   "(245) 646-5488"   ; 
            ab:firstname ?name              .
}
END_QUERY

arq_cell ${data_file} Q3 "List phone numbers by nickname or else first name." << END_QUERY
PREFIX ab: <http://learningsparql.com/ns/addressbook#> 
SELECT ?firstname ?phone
WHERE
{
    ?person ab:homeTel   ?phone .
    OPTIONAL { ?person ab:nickname ?firstname . }
    OPTIONAL { ?person ab:firstname ?firstname . }
}
END_QUERY

arq_cell ${data_file} Q4 "List everything known about Cindy." << END_QUERY
PREFIX ab: <http://learningsparql.com/ns/addressbook#> 
SELECT ?propertyName ?propertyValue
WHERE
{
    ?person ab:firstname  "Cindy"           ;
            ?propertyName ?propertyValue    . 
}
END_QUERY

arq_cell ${data_file} Q5 "List everyone who has a yahoo email address." << END_QUERY
PREFIX ab: <http://learningsparql.com/ns/addressbook#> 
SELECT ?firstname ?lastname ?email
WHERE
{
    ?person ab:email     ?email             ;
            ab:firstname ?firstname         ;
            ab:lastname  ?lastname          .
    FILTER (regex(?email, "yahoo", "i"))    .
}
END_QUERY

arq_cell ${data_file} Q6 "List everyone's home and mobile phone number." << END_QUERY
PREFIX ab: <http://learningsparql.com/ns/addressbook#> 
SELECT ?firstname ?lastname ?home ?mobile
WHERE
{
    ?person ab:lastname  ?lastname      ;
            ab:firstname ?firstname     ;
            ab:homeTel   ?home          . 
    OPTIONAL {
        ?person ab:mobileTel  ?mobile   . 
    }
}
END_QUERY

arq_cell ${data_file} Q6 "List everyone who does not have a mobile number." << END_QUERY
PREFIX ab: <http://learningsparql.com/ns/addressbook#> 
SELECT ?firstname ?lastname
WHERE
{
    ?person ab:lastname  ?lastname      ;
            ab:firstname ?firstname     .
    FILTER NOT EXISTS {
        ?person ab:mobileTel ?mobile   . 
    }
}
END_QUERY

arq_cell ${data_file} Q7 "List everyone who either has a nickname or a mobile number." << END_QUERY
PREFIX ab: <http://learningsparql.com/ns/addressbook#> 
SELECT ?firstname ?lastname ?nickname ?mobile
WHERE
{
    ?person ab:lastname  ?lastname      ;
            ab:firstname ?firstname     .
    OPTIONAL{ ?person ab:nickname ?nickname . }
    OPTIONAL{ ?person ab:mobileTel ?mobile . }
    FILTER( bound(?nickname) || bound(?mobile))

}
END_QUERY
