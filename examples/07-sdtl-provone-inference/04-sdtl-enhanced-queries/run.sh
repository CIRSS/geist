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
        ?commandinventory ?index ?command .
        ?command sdtl:OperatesOn ?operand .
        ?operand sdtl:VariableName ?used_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?used_variable ?source_line

__END_QUERY__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q2 "WHAT VARIABLES WERE DIRECTLY AFFECTED BY OTHER VARIABLES?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?affected_variable ?affecting_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory ?index ?command .
        ?command sdtl:Variable/sdtl:VariableName ?affected_variable .
        ?command sdtl:OperatesOn/sdtl:VariableName ?affecting_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?affected_variable ?affecting_variable ?source_line

__END_QUERY__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q3 "WHAT VARIABLES DIRECTLY AFFECTED THE KELVIN VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?affecting_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory ?index ?command .
        ?command sdtl:Variable/sdtl:VariableName "Kelvin"^^<http://www.w3.org/2001/XMLSchema#string> .
        ?command sdtl:OperatesOn/sdtl:VariableName ?affecting_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?affecting_variable ?source_line

__END_QUERY__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q4 "WHAT VARIABLES DIRECTLY AFFECTED VARIABLES THAT DIRECTLY AFFECTED THE KELVIN VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?indirectly_affecting_variable ?indirectly_affecting_command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .

        ?program sdtl:Commands ?commandinventory .
        ?commandinventory ?index1 ?directly_affecting_command .
        ?commandinventory ?index2 ?indirectly_affecting_command .
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

bash ${RUNNER} Q5 "WHAT VARIABLES DIRECTLY OR INDIRECTLY AFFECTED THE KELVIN VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?variable
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?commandinventory .
        ?commandinventory ?index ?command .
        ?command sdtl:Variable/sdtl:VariableName "Kelvin"^^<http://www.w3.org/2001/XMLSchema#string> .
        ?command sdtl:OperatesOn/sdtl:VariableName/(^sdtl:VariableName/^sdtl:Variable/sdtl:OperatesOn/sdtl:VariableName)* ?variable .
    } ORDER BY ?variable

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
