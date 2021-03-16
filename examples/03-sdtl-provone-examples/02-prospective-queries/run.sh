#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT SDTL-PROVONE TRACE" << END_SCRIPT

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/single-command.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} PROSPECTIVE-1 "WHAT IS THE TOP-LEVEL PROGRAM IN THE TRACE?" \
    << END_SCRIPT

geist query --format table << END_QUERY

    prefix provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>

    SELECT ?program
    WHERE {
        ?program rdf:type provone:Program .
        FILTER NOT EXISTS { ?parentProgram provone:hasSubProgram ?program . }
    }

END_QUERY

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} PROSPECTIVE-2 "WHAT ARE THE SUB-PROGRAMS IN THE TRACE?" \
    << END_SCRIPT

geist query --format table << END_QUERY

    prefix provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>

    SELECT ?subProgram
    WHERE {
        ?parentProgram provone:hasSubProgram ?subProgram .
    }

END_QUERY

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} PROSPECTIVE-3 "WHAT ARE THE OUTPUT PORTS AND ASSOCIATED VARIABLES IN THE TRACE?" \
    << END_SCRIPT

geist query --format table << END_QUERY

    prefix provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>
    prefix sdtl: <http://SDTLnamespaceURL#>

    SELECT ?program ?port ?variableName
    WHERE {
        ?port rdf:type provone:Port .
        ?program provone:hasOutPort ?port .
        ?port sdtl:variable ?variable .
        ?variable sdtl:variableName  ?variableName .
    }

END_QUERY

END_SCRIPT
