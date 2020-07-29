#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT PROVONE TRACE" << END_SCRIPT

blazegraph drop
blazegraph import --format jsonld --file ../data/branched-pipeline.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} RETROSPECTIVE-1 "WHAT FILES WERE PROVIDED AS INPUT TO THE TALE?" \
    << END_SCRIPT

blazegraph select --format csv << END_QUERY

    prefix prov: <http://www.w3.org/ns/prov#>
    prefix provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>

    SELECT ?tale_input_file
    WHERE {
        ?execution rdf:type provone:Execution .
        ?execution prov:used ?tale_input_file .
    }

END_QUERY

END_SCRIPT
