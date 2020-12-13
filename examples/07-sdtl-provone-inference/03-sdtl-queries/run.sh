#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT PROVONE TRACE" << END_SCRIPT

blazegraph destroy --dataset kb
blazegraph create --dataset kb
blazegraph import --format jsonld --file ../data/compute-sdtl.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q1 "WHAT COMMANDS ARE EXECUTED BY THE SCRIPT?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?command ?sourcetext
    WHERE {
        ?program rdf:type sdtl:Program .            # Identify the SDTL program.
        ?program sdtl:Commands ?command .
        ?command sdtl:SourceInformation ?sourceinfo .
        ?sourceinfo sdtl:originalSourceText ?sourcetext .
    }

__END_QUERY__

END_SCRIPT
