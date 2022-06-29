#!/usr/bin/env bash

# *****************************************************************************

bash_cell SETUP << END_CELL

# IMPORT SDTL-PROVONE TRACE

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/single-command.jsonld

END_CELL

# *****************************************************************************

bash_cell PROSPECTIVE-1 << END_CELL

# WHAT IS THE TOP-LEVEL PROGRAM IN THE TRACE?

geist query --format table << END_QUERY

    prefix provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>

    SELECT ?program
    WHERE {
        ?program rdf:type provone:Program .
        FILTER NOT EXISTS { ?parentProgram provone:hasSubProgram ?program . }
    }

END_QUERY

END_CELL

# *****************************************************************************

bash_cell PROSPECTIVE-2 << END_CELL

# WHAT ARE THE SUB-PROGRAMS IN THE TRACE?

geist query --format table << END_QUERY

    prefix provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>

    SELECT ?subProgram
    WHERE {
        ?parentProgram provone:hasSubProgram ?subProgram .
    }

END_QUERY

END_CELL

# *****************************************************************************

bash_cell PROSPECTIVE-3 << END_CELL

# WHAT ARE THE OUTPUT PORTS AND ASSOCIATED VARIABLES IN THE TRACE?

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

END_CELL
