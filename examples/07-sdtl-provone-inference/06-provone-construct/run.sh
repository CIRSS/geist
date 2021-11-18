#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'
GRAPHER='../../common/run_dot_examples.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT SDTL AS JSON-LD" << END_SCRIPT

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet 
geist import --format jsonld --file ../data/compute-sdth.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} E1 "EXPORT ORIGINAL SDTL AS N-TRIPLES" << END_SCRIPT

geist export --format nt | sort

END_SCRIPT


bash ${RUNNER} Q1 "CONSTRUCT PROVONE PROGRAMS VIA CONSTRUCT QUERY" << END_SCRIPT

geist query --format table << __END_QUERY__

    PREFIX prov: <http://www.w3.org/ns/prov#>
    PREFIX provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>
    PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
    PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
    PREFIX sdth: <https://rdf-vocabulary.ddialliance.org/sdth#>

    CONSTRUCT {
        ?program rdf:type provone:Program . 
    }
    WHERE {
        {
            ?program rdf:type sdth:Program . 
        }
        UNION
        {
            ?program rdf:type sdth:ProgramStep .
        }
    } 


__END_QUERY__

END_SCRIPT

