#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT SDTL" << END_SCRIPT

blazegraph destroy --dataset kb
blazegraph create --dataset kb
blazegraph import --format jsonld --file ../data/compute-sdtl.jsonld

END_SCRIPT



# *****************************************************************************

bash ${RUNNER} Q1 "WHAT COMMANDS ARE EXECUTED BY THE SCRIPT?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?command ?sourceline ?sourcetext
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command sdtl:SourceInformation ?sourceinfo .
        ?sourceinfo sdtl:LineNumberStart ?sourceline .
        ?sourceinfo sdtl:OriginalSourceText ?sourcetext .
    } ORDER BY ?line

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q2 "WHAT DATA FILES ARE LOADED BY THE SCRIPT?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?filename ?command ?sourceline ?sourcetext
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command rdf:type sdtl:Load .
        ?command sdtl:FileName ?filename .
        ?command sdtl:SourceInformation ?sourceinfo .
        ?sourceinfo sdtl:LineNumberStart ?sourceline .
        ?sourceinfo sdtl:OriginalSourceText ?sourcetext .
    }

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q3 "WHAT DATA FILES ARE SAVED BY THE SCRIPT?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?filename ?command ?sourceline ?sourcetext
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command rdf:type sdtl:Save .
        ?command sdtl:FileName ?filename .
        ?command sdtl:SourceInformation ?sourceinfo .
        ?sourceinfo sdtl:LineNumberStart ?sourceline .
        ?sourceinfo sdtl:OriginalSourceText ?sourcetext .
    }

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q4 "WHAT VARIABLES ARE LOADED BY THE SCRIPT?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?variablename ?dataframe ?sourceline ?sourcetext
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command rdf:type sdtl:Load .
        ?command sdtl:ProducesDataframe ?dataframedesc .
        ?dataframedesc sdtl:DataframeName ?dataframe .
        ?dataframedesc sdtl:VariableInventory ?variablename .
        ?command sdtl:SourceInformation ?sourceinfo .
        ?sourceinfo sdtl:LineNumberStart ?sourceline .
        ?sourceinfo sdtl:OriginalSourceText ?sourcetext .
    }

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q5 "WHAT VARIABLES ARE SAVED BY THE SCRIPT?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?variablename ?dataframe ?sourceline ?sourcetext
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command rdf:type sdtl:Save .
        ?command sdtl:ProducesDataframe ?dataframedesc .
        ?dataframedesc sdtl:DataframeName ?dataframe .
        ?dataframedesc sdtl:VariableInventory ?variablename .
        ?command sdtl:SourceInformation ?sourceinfo .
        ?sourceinfo sdtl:LineNumberStart ?sourceline .
        ?sourceinfo sdtl:OriginalSourceText ?sourcetext .
    }

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q6 "WHAT COMMANDS UPDATE EACH DATAFRAME?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?dataframe ?command ?sourceline ?sourcetext
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command sdtl:ProducesDataframe ?dataframedesc .
        ?dataframedesc sdtl:DataframeName ?dataframe .
        ?dataframedesc sdtl:VariableInventory ?variable .
        ?command sdtl:SourceInformation ?sourceinfo .
        ?sourceinfo sdtl:LineNumberStart ?sourceline .
        ?sourceinfo sdtl:OriginalSourceText ?sourcetext .
    } ORDER BY ?dataframe ?sourceline

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q7 "WHAT COMMANDS UPDATE EACH VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?variablename ?command ?sourceline ?sourcetext
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command sdtl:Variable ?variable .
        ?variable sdtl:VariableName ?variablename .
        ?command sdtl:SourceInformation ?sourceinfo .
        ?sourceinfo sdtl:LineNumberStart ?sourceline .
        ?sourceinfo sdtl:OriginalSourceText ?sourcetext .
    } ORDER BY ?variable ?sourceline

__END_QUERY__

END_SCRIPT


bash ${RUNNER} Q8 "WHAT COMMANDS USE EACH VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?usedvariable ?command ?sourceline ?sourcetext
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command sdtl:Expression ?expression .
        ?expression (sdtl:Arguments/sdtl:ArgumentValue)+/sdtl:VariableName ?usedvariable .
        ?command sdtl:SourceInformation ?sourceinfo .
        ?sourceinfo sdtl:LineNumberStart ?sourceline .
        ?sourceinfo sdtl:OriginalSourceText ?sourcetext .
    } ORDER BY ?usedvariable ?sourceline

__END_QUERY__

END_SCRIPT

