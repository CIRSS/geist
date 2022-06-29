#!/usr/bin/env bash

# *****************************************************************************

bash_cell SETUP << END_CELL

# IMPORT SDTL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/compute-sdtl.jsonld

END_CELL

# *****************************************************************************

bash_cell E1 << END_CELL

# EXPORT AS N-TRIPLES

geist export --format nt | sort

END_CELL

# *****************************************************************************

bash_cell Q1 << END_CELL

# WHAT COMMANDS ARE EXECUTED BY THE SCRIPT?

geist query --format table << __END_QUERY__

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

END_CELL



bash_cell Q2 << END_CELL

# WHAT DATA FILES ARE LOADED BY THE SCRIPT?

geist query --format table << __END_QUERY__

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

END_CELL



bash_cell Q3 << END_CELL

# WHAT DATA FILES ARE SAVED BY THE SCRIPT?

geist query --format table << __END_QUERY__

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

END_CELL



bash_cell Q4 << END_CELL

# WHAT VARIABLES ARE LOADED BY THE SCRIPT?

geist query --format table << __END_QUERY__

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

END_CELL



bash_cell Q5 << END_CELL

# WHAT VARIABLES ARE SAVED BY THE SCRIPT?

geist query --format table << __END_QUERY__

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

END_CELL



bash_cell Q6 << END_CELL

# WHAT COMMANDS UPDATE EACH DATAFRAME?

geist query --format table << __END_QUERY__

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

END_CELL



bash_cell Q7 << END_CELL

# WHAT COMMANDS UPDATE EACH VARIABLE?

geist query --format table << __END_QUERY__

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

END_CELL


bash_cell Q8 << END_CELL

# WHAT COMMANDS USE EACH VARIABLE?

geist query --format table << __END_QUERY__

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

END_CELL



bash_cell Q9 << END_CELL

# WHAT VARIABLES WERE DIRECTLY AFFECTED BY OTHER VARIABLES?

geist query --format table << __END_QUERY__

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

END_CELL


bash_dot_cell GRAPH-1 << '__END_CELL__'

# DATAFRAME FLOW THROUGH COMMANDS

geist report << '__END_REPORT_TEMPLATE__'

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

__END_CELL__
