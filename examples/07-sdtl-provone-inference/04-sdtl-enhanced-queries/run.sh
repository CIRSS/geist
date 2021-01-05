#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'
GRAPHER='../../common/run_dot_examples.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT SDTL" << END_SCRIPT

blazegraph destroy --dataset kb
blazegraph create --dataset kb --infer owl
blazegraph import --file ../data/sdtl-enhanced-rules.ttl
blazegraph import --format jsonld --file ../data/compute-sdtl.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} E1 "EXPORT AS N-TRIPLES" << END_SCRIPT

blazegraph export --format nt | sort

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q1 "WHAT COMMANDS USE EACH VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?used_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory rdfs:member ?command .
        ?command sdtl:OperatesOn ?operand .
        ?operand sdtl:VariableName ?used_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?used_variable ?source_line

__END_QUERY__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q2 "WHAT VARIABLES DIRECTLY AFFECT OTHER VARIABLES?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?affecting_variable  ?affected_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory rdfs:member ?command .
        ?command sdtl:Variable/sdtl:VariableName ?affected_variable .
        ?command sdtl:OperatesOn/sdtl:VariableName ?affecting_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?affecting_variable ?affected_variable ?source_line

__END_QUERY__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q3 "WHAT VARIABLES DIRECTLY AFFECT THE KELVIN VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?affecting_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory rdfs:member ?command .
        ?command sdtl:Variable/sdtl:VariableName "Kelvin"^^<http://www.w3.org/2001/XMLSchema#string> .
        ?command sdtl:OperatesOn/sdtl:VariableName ?affecting_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?affecting_variable ?source_line

__END_QUERY__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q4 "WHAT VARIABLES DIRECTLY AFFECT VARIABLES THAT DIRECTLY AFFECT THE KELVIN VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?indirectly_affecting_variable ?indirectly_affecting_command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .

        ?program sdtl:Commands ?commandinventory .
        ?commandinventory rdfs:member ?directly_affecting_command .
        ?commandinventory rdfs:member ?indirectly_affecting_command .
        ?directly_affecting_command sdtl:Variable/sdtl:VariableName "Kelvin"^^<http://www.w3.org/2001/XMLSchema#string> .
        ?directly_affecting_command sdtl:OperatesOn/sdtl:VariableName/^sdtl:VariableName/^sdtl:Variable ?indirectly_affecting_command  .
        ?indirectly_affecting_command sdtl:OperatesOn/sdtl:VariableName ?indirectly_affecting_variable .
        ?indirectly_affecting_command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?affected_variable ?affecting_variable ?source_line

__END_QUERY__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q5 "WHAT VARIABLES DIRECTLY OR INDIRECTLY AFFECT THE KELVIN VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?variable
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory rdfs:member ?command .
        ?command sdtl:Variable/sdtl:VariableName "Kelvin"^^<http://www.w3.org/2001/XMLSchema#string> .
        ?command sdtl:OperatesOn/sdtl:VariableName/(^sdtl:VariableName/^sdtl:Variable/sdtl:OperatesOn/sdtl:VariableName)* ?variable .
    } ORDER BY ?variable

__END_QUERY__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q6 "WHAT COMMANDS AFFECT EACH VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?affected_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory rdfs:member ?command .
        {
            ?command sdtl:Variable ?variable .
        }
        UNION
        {
            ?command rdf:type sdtl:Load .
            ?command sdtl:ProducesDataframe ?dataframe_description .
            ?dataframe_description sdtl:VariableInventory ?variable_inventory .
            ?variable_inventory rdfs:member ?variable .
        }
        ?variable sdtl:VariableName ?affected_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?affected_variable ?source_line

__END_QUERY__

END_SCRIPT


# *****************************************************************************

bash ${RUNNER} Q7 "WHAT IS THE LAST COMMAND THAT UPDATES EACH VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?variable ?command ?upstream_command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory rdfs:member ?command .
        {
            ?command rdf:type sdtl:Save .
            ?save_command sdtl:ConsumesDataframe ?saved_dataframe .
            ?saved_dataframe sdtl:VariableInventory/rdfs:member/sdtl:VariableName ?variable .
        }
        UNION
        {
            ?command rdf:type sdtl:Compute .
            ?command sdtl:OperatesOn/sdtl:VariableName ?variable .
        }

        ?save_command (sdtl:ConsumesDataframe/^sdtl:ProducesDataframe)+ ?upstream_command .

        {
            ?upstream_command sdtl:Variable/sdtl:VariableName  ?variable .
        }
        UNION
        {
            ?upstream_command rdf:type sdtl:Load .
            ?upstream_command sdtl:ProducesDataframe ?dataframe_description .
            ?dataframe_description sdtl:VariableInventory ?variable_inventory .
            ?variable_inventory rdfs:member/sdtl:VariableName ?variable .
        }

        FILTER NOT EXISTS
        {
            ?intermediate_command (sdtl:ConsumesDataframe/^sdtl:ProducesDataframe)+ ?upstream_command .

            {
                ?intermediate_command sdtl:Variable/sdtl:VariableName  ?variable .
            }
            UNION
            {
                ?intermediate_command rdf:type sdtl:Load .
                ?intermediate_command sdtl:ProducesDataframe ?dataframe_description .
                ?dataframe_description sdtl:VariableInventory ?variable_inventory .
                ?variable_inventory rdfs:member/sdtl:VariableName ?variable .
            }
        }

        ?upstream_command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .

    } ORDER BY ?source_line ?command

__END_QUERY__

END_SCRIPT

# *****************************************************************************

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
            {{ gv_labeled_edge (index $Edge 0) (index $Edge 1) (index $Edge 3) }}
        {{ end }}                                                                           \\
                                                                                            \\
    {{ end }}                                                                               \\
                                                                                            \\
    {{ gv_cluster_end }}

    {{ gv_end }}                                                                            \\

__END_REPORT_TEMPLATE__

__END_SCRIPT__


# # *****************************************************************************

# bash ${GRAPHER} GRAPH-2 "VARIABLE FLOW THROUGH COMMANDS" \
#     << '__END_SCRIPT__'

# blazegraph report << '__END_REPORT_TEMPLATE__'

#     {{{
#         {{ include "../../common/graphviz.g" }}
#         {{ include "../../common/sdtl.g" }}
#     }}}

#     {{ gv_graph "sdtl_program" }}

#     {{ gv_title "Dataframe-flow through commands" }}

#     {{ gv_cluster "program_graph" }}

#     # command nodes
#     {{ sdtl_program_node_style }}
#     node[width=8]
#     {{ with $ProgramID := sdtl_select_program | value }}                                    \\

#         {{ range $Command := (sdtl_select_commands $ProgramID | rows ) }}                   \\
#             {{ gv_labeled_node (index $Command 0) (index $Command 1) }}
#         {{ end }}                                                                           \\

#         # dataframe edges
#         {{ range $Edge := (sdtl_select_compute_variable_compute_edges $ProgramID | rows) }} \\
#             {{ gv_labeled_edge (index $Edge 0) (index $Edge 1) (index $Edge 2) }}
#         {{ end }}                                                                           \\
#                                                                                             \\
#         {{ range $Edge := (sdtl_select_load_variable_compute_edges $ProgramID | rows) }} \\
#             {{ gv_labeled_edge (index $Edge 0) (index $Edge 1) (index $Edge 2) }}
#         {{ end }}                                                                           \\

#     {{ end }}                                                                               \\
#                                                                                             \\
#     {{ gv_cluster_end }}

#     {{ gv_end }}                                                                            \\

# __END_REPORT_TEMPLATE__

# __END_SCRIPT__

# *****************************************************************************

bash ${RUNNER} R1 "REPORT HISTORY OF EACH VARIABLE" << 'END_SCRIPT'

blazegraph report << '__END_REPORT_TEMPLATE__'

{{{
    {{ include "../../common/sdtl.g" }}
}}}

---------- VARIABLE HISTORY REPORT ----------
                                                                                                    \\
{{ with $Program := sdtl_select_program | value }}                                                  \\
    {{ range $SaveCommand := sdtl_select_save_commands $Program | rows }}
        ===================
        Dataframe {{ index $SaveCommand 2 }}
        ===================
        {{ range $VariableName := sdtl_select_dataframe_variables (index $SaveCommand 1) | vector }}
            Variable {{ $VariableName }}
            -------------------
            {{ range $LoadCommand := sdtl_select_load_commands $Program $VariableName | rows }}     \\
                Load    | Line {{ index $LoadCommand 3 }} | {{ index $LoadCommand 4 }}
            {{ end }}                                                                               \\
            {{ range $UpdateCommand := sdtl_select_update_commands $Program $VariableName | rows }} \\
                Compute | Line {{ index $UpdateCommand 1 }} | {{ index $UpdateCommand 2 }}
            {{ end }}                                                                               \\
                Save    | Line {{ index $SaveCommand 3 }} | {{ index $SaveCommand 4 }}
        {{ end }}
    {{ end }}
{{ end }}

__END_REPORT_TEMPLATE__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} R2 "REPORT LAST UPDATE CORRESPONDING TO EACH VARIABLE USE" << 'END_SCRIPT'

blazegraph report << '__END_REPORT_TEMPLATE__'

{{{
    {{ include "../../common/sdtl.g" }}
}}}

---------- VARIABLE HISTORY REPORT ----------
                                                                                                    \\
{{ with $Program := sdtl_select_program | value }}                                                  \\
    {{ range $SaveCommand := sdtl_select_save_commands $Program | rows }}
        ===================
        Dataframe {{ index $SaveCommand 2 }}
        ===================
        {{ range $VariableName := sdtl_select_dataframe_variables (index $SaveCommand 1) | vector }}
            Variable {{ $VariableName }}
            -------------------
            {{ range $VariableFlow := sdtl_select_variable_flows $Program $VariableName | rows }}      \\
            [Line {{ index $VariableFlow 0 }}] {{ index $VariableFlow 1 }} ---> [Line {{ index $VariableFlow 2 }}] {{ index $VariableFlow 3 }}
            {{ end }}                                                                               \\
        {{ end }}
    {{ end }}
{{ end }}

__END_REPORT_TEMPLATE__

END_SCRIPT


# *****************************************************************************

bash ${RUNNER} R3 "WHAT COMMANDS WRITE TO EACH VARIABLE?" << END_SCRIPT

blazegraph report << '__END_REPORT_TEMPLATE__'

{{{
    {{ include "../../common/sdtl.g" }}
}}}

---------- COMMANDS THAT WRITE TO EACH VARIABLE ----------

{{ select '''
    SELECT DISTINCT ?written_variable ?source_line ?source_text
    WHERE {
        {{ program_has_commands "?program" "?command" }} .
        {{ command_writes_variable  "?command" "?written_variable" }} .
        {{ command_has_source "?command" "?source_line" "?source_text" }}
    }
    ORDER BY ?written_variable ?source_line
''' | tabulate }}

__END_REPORT_TEMPLATE__

END_SCRIPT


# *****************************************************************************

bash ${RUNNER} R3 "WHAT COMMANDS READ FROM EACH VARIABLE?" << END_SCRIPT

blazegraph report << '__END_REPORT_TEMPLATE__'

{{{
    {{ include "../../common/sdtl.g" }}
}}}

---------- COMMANDS THAT READ FROM EACH VARIABLE ----------

{{ select '''
    SELECT DISTINCT ?read_variable ?source_line ?source_text
    WHERE {
        {{ program_has_commands "?program" "?command" }} .
        {{ command_reads_variable  "?command" "?read_variable" }} .
        {{ command_has_source "?command" "?source_line" "?source_text" }}
    }
    ORDER BY ?read_variable ?source_line
''' | tabulate }}

__END_REPORT_TEMPLATE__

END_SCRIPT


# *****************************************************************************

bash ${RUNNER} R4 "WHAT COMMANDS WRITE VARIABLES READ BY DOWNSTREAM COMMANDS?" << END_SCRIPT

blazegraph report << '__END_REPORT_TEMPLATE__'

{{{
    {{ include "../../common/sdtl.g" }}
}}}

---------- COMMANDS THAT WRITE VARIABLES READ BY DOWNSTREAM COMMANDS ----------

{{ select '''
    SELECT DISTINCT ?top_variable ?writer_line ?writer_text ?reader_line ?reader_text
    WHERE {
        {{ program_has_commands "?top_program" "?top_reader" }} .
        {{ upstream_writer "?top_variable" "?top_reader" "?top_writer" }}
        {{ command_has_source "?top_writer" "?writer_line" "?writer_text" }}
        {{ command_has_source_2 "?top_reader" "?reader_line" "?reader_text" }}
    }
    ORDER BY ?top_variable ?writer_line
''' | tabulate }}

__END_REPORT_TEMPLATE__

END_SCRIPT
