#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'
GRAPHER='../../common/run_dot_examples.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT SDTL" << END_SCRIPT

blazegraph destroy --dataset kb --quiet
blazegraph create --dataset kb --quiet
blazegraph import --format jsonld --file ../data/compute-sdtl.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} E1 "EXPORT AS N-TRIPLES" << END_SCRIPT

blazegraph export --format nt | sort

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q1 "WHAT COMMANDS ARE EXECUTED BY THE SCRIPT?" << END_SCRIPT

blazegraph query --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>
    PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
    PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>

    SELECT DISTINCT ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory (<>|!<>) ?command .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?source_line

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q2 "WHAT DATA FILES ARE LOADED BY THE SCRIPT?" << END_SCRIPT

blazegraph query --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?file_name ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory (<>|!<>) ?command .
        ?command rdf:type sdtl:Load .
        ?command sdtl:FileName ?file_name .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    }

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q3 "WHAT DATA FILES ARE SAVED BY THE SCRIPT?" << END_SCRIPT

blazegraph query --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?file_name ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory (<>|!<>) ?command .
        ?command rdf:type sdtl:Save .
        ?command sdtl:FileName ?file_name .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    }

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q4 "WHAT VARIABLES ARE LOADED BY THE SCRIPT?" << END_SCRIPT

blazegraph query --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?loaded_variable ?dataframe ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory (<>|!<>) ?command .
        ?command rdf:type sdtl:Load .
        ?command sdtl:ProducesDataframe ?dataframe_description .
        ?dataframe_description sdtl:DataframeName ?dataframe .
        ?dataframe_description sdtl:VariableInventory ?variable_inventory .
        ?variable_inventory (<>|!<>)/sdtl:VariableName ?loaded_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?loaded_variable ?source_line

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q5 "WHAT VARIABLES ARE SAVED BY THE SCRIPT?" << END_SCRIPT

blazegraph query --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?saved_variable ?dataframe ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory (<>|!<>) ?command .
        ?command rdf:type sdtl:Save .
        ?command sdtl:ConsumesDataframe ?dataframe_description .
        ?dataframe_description sdtl:DataframeName ?dataframe .
        ?dataframe_description sdtl:VariableInventory ?variable_inventory .
        ?variable_inventory (<>|!<>)/sdtl:VariableName ?saved_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?saved_variable ?source_line

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q6 "WHAT COMMANDS UPDATE EACH DATAFRAME?" << END_SCRIPT

blazegraph query --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?dataframe ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory (<>|!<>) ?command .
        ?command sdtl:ProducesDataframe ?dataframe_description .
        ?dataframe_description sdtl:DataframeName ?dataframe .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?dataframe ?source_line

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q7 "WHAT COMMANDS UPDATE EACH VARIABLE?" << END_SCRIPT

blazegraph query --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?updated_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory (<>|!<>) ?command .
        ?command sdtl:Variable ?variable .
        ?variable sdtl:VariableName ?updated_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?variable ?source_line

__END_QUERY__

END_SCRIPT


bash ${RUNNER} Q8 "WHAT COMMANDS USE EACH VARIABLE?" << END_SCRIPT

blazegraph query --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?used_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory (<>|!<>) ?command .
        ?command sdtl:Expression ?expression .
        ?expression (sdtl:Arguments/sdtl:ArgumentValue)+/sdtl:VariableName ?used_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?used_variable ?source_line

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q9 "WHAT VARIABLES WERE DIRECTLY AFFECTED BY OTHER VARIABLES?" << END_SCRIPT

blazegraph query --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?affected_variable ?affecting_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory (<>|!<>) ?command .
        ?command sdtl:Variable/sdtl:VariableName ?affected_variable .
        ?command sdtl:Expression ?expression .
        ?expression (sdtl:Arguments/sdtl:ArgumentValue)+/sdtl:VariableName ?affecting_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?affected_variable ?affecting_variable ?source_line

__END_QUERY__

END_SCRIPT


bash ${GRAPHER} GRAPH-1 "DATAFRAME FLOW THROUGH COMMANDS" \
    << '__END_SCRIPT__'

blazegraph report << '__END_REPORT_TEMPLATE__'

    {{{
        {{ include "../../common/graphviz.g" }}
        {{ include "../../common/sdtl.g" }}
    }}}

    {{ gv_graph "sdtl_program" }}

    {{ gv_title "Dataframe-flow through commands" }}

    {{ gv_cluster "program_graph" }}

    # command nodes
    {{ sdtl_program_node_style }}
    node[width=8]
    {{ with $ProgramID := sdtl_select_program | value }}                                    \\

        {{ range $Command := (sdtl_select_commands $ProgramID | rows ) }}                   \\
            {{ gv_labeled_node (index $Command 0) (index $Command 1) }}
        {{ end }}                                                                           \\

        # dataframe edges
        {{ range $Edge := (sdtl_select_dataframe_edges $ProgramID | rows) }}                \\
            {{ gv_edge (index $Edge 0) (index $Edge 1) }}
        {{ end }}                                                                           \\
                                                                                            \\
    {{ end }}                                                                               \\
                                                                                            \\
    {{ gv_cluster_end }}

    {{ gv_end }}                                                                            \\

__END_REPORT_TEMPLATE__

__END_SCRIPT__
