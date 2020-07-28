#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT SDTL-PROVONE TRACE" << END_SCRIPT

blazegraph drop
blazegraph import --format jsonld --file ../data/single-command.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q1 "WHAT IS THE TOP-LEVEL PROGRAM IN THE TRACE?" << END_SCRIPT

blazegraph query --format csv << END_QUERY

    prefix provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>

    SELECT ?program
    WHERE {
        ?program rdf:type provone:Program .
        FILTER NOT EXISTS { ?parentProgram provone:hasSubProgram ?program . }
    }

END_QUERY

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q2 "WHAT ARE THE SUB-PROGRAMS IN THE TRACE?" << END_SCRIPT

blazegraph query --format csv << END_QUERY

    prefix provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>

    SELECT ?subProgram
    WHERE {
        ?parentProgram provone:hasSubProgram ?subProgram .
    }

END_QUERY

END_SCRIPT