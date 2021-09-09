#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'
GRAPHER='../../common/run_dot_examples.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT SDTL" << END_SCRIPT

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet 
geist import --format jsonld --file ../data/compute-sdth.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} E1 "EXPORT AS N-TRIPLES" << END_SCRIPT

geist export --format nt | sort

END_SCRIPT


bash ${RUNNER} Q1 "WHAT STEPS ARE EXECUTED BY THE PROGRAM?" << END_SCRIPT

geist query --format table << __END_QUERY__

    PREFIX sdth: <https://rdf-vocabulary.ddialliance.org/sdth#>
    PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
    PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>

    SELECT DISTINCT ?step ?step_source_text
    WHERE {
        ?program rdf:type sdth:Program .
        ?program sdth:hasStep ?step .
        ?step sdth:hasSourceCode ?step_source_text .
    }

__END_QUERY__

END_SCRIPT


bash ${RUNNER} Q2 "WHAT DATA FILES ARE LOADED BY THE PROGRAM?" << END_SCRIPT

geist query --format table << __END_QUERY__

    PREFIX sdth: <https://rdf-vocabulary.ddialliance.org/sdth#>

    SELECT DISTINCT ?file_name ?step ?step_source_text
    WHERE {
        ?program rdf:type sdth:Program .
        ?program sdth:hasStep ?step .
        ?step sdth:loadsFile ?file .
        ?file sdth:hasName ?file_name .
        ?step sdth:hasSourceCode ?step_source_text .
    }

__END_QUERY__

END_SCRIPT


bash ${RUNNER} Q3 "WHAT DATA FILES ARE SAVED BY THE PROGRAM?" << END_SCRIPT

geist query --format table << __END_QUERY__

    PREFIX sdth: <https://rdf-vocabulary.ddialliance.org/sdth#>

    SELECT DISTINCT ?file_name ?step ?step_source_text
    WHERE {
        ?program rdf:type sdth:Program .
        ?program sdth:hasStep ?step .
        ?step sdth:savesFile ?file .
        ?file sdth:hasName ?file_name .
        ?step sdth:hasSourceCode ?step_source_text .
    }

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q4 "WHAT VARIABLES ARE LOADED BY THE PROGRAM?" << END_SCRIPT

geist query --format table << __END_QUERY__

    PREFIX sdth: <https://rdf-vocabulary.ddialliance.org/sdth#>

    SELECT DISTINCT ?dataframe_name ?variable_name ?step_source_text
    WHERE {
        ?program rdf:type sdth:Program .
        ?program sdth:hasStep ?step .
        ?step sdth:loadsFile ?file .
        ?step sdth:producesDataframe ?dataframe .
        ?dataframe sdth:includesVariable ?variable .
        ?variable sdth:hasName ?variable_name .
        ?dataframe sdth:hasName ?dataframe_name .
        ?step sdth:hasSourceCode ?step_source_text .
    } ORDER BY ?variable_name ?source_line

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q5 "WHAT VARIABLES ARE SAVED BY THE SCRIPT?" << END_SCRIPT

geist query --format table << __END_QUERY__

    PREFIX sdth: <https://rdf-vocabulary.ddialliance.org/sdth#>

    SELECT DISTINCT ?dataframe_name ?variable_name ?step_source_text
    WHERE {
        ?program rdf:type sdth:Program .
        ?program sdth:hasStep ?step .
        ?step sdth:savesFile ?file .
        ?step sdth:consumesDataframe ?dataframe .
        ?dataframe sdth:includesVariable ?variable .
        ?variable sdth:hasName ?variable_name .
        ?dataframe sdth:hasName ?dataframe_name .
        ?step sdth:hasSourceCode ?step_source_text .

    } ORDER BY ?saved_variable ?source_line

__END_QUERY__

END_SCRIPT



# bash ${RUNNER} Q6 "WHAT COMMANDS UPDATE EACH DATAFRAME?" << END_SCRIPT

# geist query --format table << __END_QUERY__

#     PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

#     SELECT DISTINCT ?dataframe ?command ?source_line ?source_text
#     WHERE {
#         ?program rdf:type sdtl:Program .
#         ?program sdtl:Commands ?commandinventory .
#         ?commandinventory (<>|!<>) ?command .
#         ?command sdtl:ProducesDataframe ?dataframe_description .
#         ?dataframe_description sdtl:DataframeName ?dataframe .
#         ?command sdtl:SourceInformation ?source_info .
#         ?source_info sdtl:LineNumberStart ?source_line .
#         ?source_info sdtl:OriginalSourceText ?source_text .
#     } ORDER BY ?dataframe ?source_line

# __END_QUERY__

# END_SCRIPT



# bash ${RUNNER} Q7 "WHAT COMMANDS UPDATE EACH VARIABLE?" << END_SCRIPT

# geist query --format table << __END_QUERY__

#     PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

#     SELECT DISTINCT ?updated_variable ?command ?source_line ?source_text
#     WHERE {
#         ?program rdf:type sdtl:Program .
#         ?program sdtl:Commands ?commandinventory .
#         ?commandinventory (<>|!<>) ?command .
#         ?command sdtl:Variable ?variable .
#         ?variable sdtl:VariableName ?updated_variable .
#         ?command sdtl:SourceInformation ?source_info .
#         ?source_info sdtl:LineNumberStart ?source_line .
#         ?source_info sdtl:OriginalSourceText ?source_text .
#     } ORDER BY ?variable ?source_line

# __END_QUERY__

# END_SCRIPT


# bash ${RUNNER} Q8 "WHAT COMMANDS USE EACH VARIABLE?" << END_SCRIPT

# geist query --format table << __END_QUERY__

#     PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

#     SELECT DISTINCT ?used_variable ?command ?source_line ?source_text
#     WHERE {
#         ?program rdf:type sdtl:Program .
#         ?program sdtl:Commands ?commandinventory .
#         ?commandinventory (<>|!<>) ?command .
#         ?command sdtl:Expression ?expression .
#         ?expression (sdtl:Arguments/sdtl:ArgumentValue)+/sdtl:VariableName ?used_variable .
#         ?command sdtl:SourceInformation ?source_info .
#         ?source_info sdtl:LineNumberStart ?source_line .
#         ?source_info sdtl:OriginalSourceText ?source_text .
#     } ORDER BY ?used_variable ?source_line

# __END_QUERY__

# END_SCRIPT



# bash ${RUNNER} Q9 "WHAT VARIABLES WERE DIRECTLY AFFECTED BY OTHER VARIABLES?" << END_SCRIPT

# geist query --format table << __END_QUERY__

#     PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

#     SELECT DISTINCT ?affected_variable ?affecting_variable ?command ?source_line ?source_text
#     WHERE {
#         ?program rdf:type sdtl:Program .
#         ?program sdtl:Commands ?commandinventory .
#         ?commandinventory (<>|!<>) ?command .
#         ?command sdtl:Variable/sdtl:VariableName ?affected_variable .
#         ?command sdtl:Expression ?expression .
#         ?expression (sdtl:Arguments/sdtl:ArgumentValue)+/sdtl:VariableName ?affecting_variable .
#         ?command sdtl:SourceInformation ?source_info .
#         ?source_info sdtl:LineNumberStart ?source_line .
#         ?source_info sdtl:OriginalSourceText ?source_text .
#     } ORDER BY ?affected_variable ?affecting_variable ?source_line

# __END_QUERY__

# END_SCRIPT


# # bash ${GRAPHER} GRAPH-1 "DATAFRAME FLOW THROUGH COMMANDS" \
# #     << '__END_SCRIPT__'

# # geist report << '__END_REPORT_TEMPLATE__'

# #     {{{
# #         {{ include "../../common/graphviz.g" }}
# #         {{ include "../../common/sdtl.g" }}
# #     }}}

# #     {{ gv_graph "sdtl_program" }}

# #     {{ gv_title "Dataframe-flow through commands" }}

# #     {{ gv_cluster "program_graph" }}

# #     # command nodes
# #     {{ sdtl_program_node_style }}
# #     node[width=8]
# #     {{ with $ProgramID := sdtl_select_program | value }}                                    \\

# #         {{ range $Command := (sdtl_select_commands $ProgramID | rows ) }}                   \\
# #             {{ gv_labeled_node (index $Command 0) (index $Command 1) }}
# #         {{ end }}                                                                           \\

# #         # dataframe edges
# #         {{ range $Edge := (sdtl_select_dataframe_edges $ProgramID | rows) }}                \\
# #             {{ gv_edge (index $Edge 0) (index $Edge 1) }}
# #         {{ end }}                                                                           \\
# #                                                                                             \\
# #     {{ end }}                                                                               \\
# #                                                                                             \\
# #     {{ gv_cluster_end }}

# #     {{ gv_end }}                                                                            \\

# # __END_REPORT_TEMPLATE__

# # __END_SCRIPT__


# # *****************************************************************************

# bash ${RUNNER} Q1 "WHAT COMMANDS USE EACH VARIABLE?" << END_SCRIPT

# geist query --format table << __END_QUERY__

#     {{{

#     {{ include "../../common/sdtl.g" }}

#     }}}

#     PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

#     SELECT DISTINCT ?used_variable ?command ?source_line ?source_text
#     WHERE {
#         ?program rdf:type sdtl:Program .
#         ?program sdtl:Commands/rdfs:member ?command .
#         ?command sdtl:OperatesOn/sdtl:VariableName ?used_variable .
#         {{ command_source "?command" "?source_line" "?source_text" }}
#     } ORDER BY ?used_variable ?source_line

# __END_QUERY__


# END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} Q2 "WHAT VARIABLES DIRECTLY AFFECT OTHER VARIABLES?" << END_SCRIPT

# geist query --format table << __END_QUERY__

#     {{{

#     {{ include "../../common/sdtl.g" }}

#     }}}

#     PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

#     SELECT DISTINCT ?affecting_variable  ?affected_variable ?command ?source_line ?source_text
#     WHERE {
#         ?program rdf:type sdtl:Program .
#         ?program sdtl:Commands/rdfs:member ?command .
#         ?command sdtl:Variable/sdtl:VariableName ?affected_variable .
#         ?command sdtl:OperatesOn/sdtl:VariableName ?affecting_variable .
#         {{ command_source "?command" "?source_line" "?source_text" }}
#     } ORDER BY ?affecting_variable ?affected_variable ?source_line

# __END_QUERY__

# END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} Q3 "WHAT VARIABLES DIRECTLY AFFECT THE KELVIN VARIABLE?" << END_SCRIPT

# geist query --format table << __END_QUERY__

#     {{{

#     {{ include "../../common/sdtl.g" }}

#     }}}

#     PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

#     SELECT DISTINCT ?affecting_variable ?command ?source_line ?source_text
#     WHERE {
#         ?program rdf:type sdtl:Program .
#         ?program sdtl:Commands/rdfs:member ?command .
#         ?command sdtl:Variable/sdtl:VariableName "Kelvin"^^<http://www.w3.org/2001/XMLSchema#string> .
#         ?command sdtl:OperatesOn/sdtl:VariableName ?affecting_variable
#         {{ command_source "?command" "?source_line" "?source_text" }}
#     } ORDER BY ?affecting_variable ?source_line

# __END_QUERY__

# END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} Q4 "WHAT VARIABLES DIRECTLY AFFECT VARIABLES THAT DIRECTLY AFFECT THE KELVIN VARIABLE?" << END_SCRIPT

# geist query --format table << __END_QUERY__

#     {{{

#     {{ include "../../common/sdtl.g" }}

#     }}}

#     PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

#     SELECT DISTINCT ?indirectly_affecting_variable ?indirectly_affecting_command ?source_line ?source_text
#     WHERE {
#         ?program rdf:type sdtl:Program .

#         ?program sdtl:Commands ?commandinventory .
#         ?commandinventory rdfs:member ?directly_affecting_command .
#         ?commandinventory rdfs:member ?indirectly_affecting_command .
#         ?directly_affecting_command sdtl:Variable/sdtl:VariableName "Kelvin"^^<http://www.w3.org/2001/XMLSchema#string> .
#         ?directly_affecting_command sdtl:OperatesOn/sdtl:VariableName/^sdtl:VariableName/^sdtl:Variable ?indirectly_affecting_command  .
#         ?indirectly_affecting_command sdtl:OperatesOn/sdtl:VariableName ?indirectly_affecting_variable .
#         {{ command_source "?indirectly_affecting_command" "?source_line" "?source_text" }}
#     } ORDER BY ?affected_variable ?affecting_variable ?source_line

# __END_QUERY__

# END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} Q5 "WHAT VARIABLES DIRECTLY OR INDIRECTLY AFFECT THE KELVIN VARIABLE?" << END_SCRIPT

# geist query --format table << __END_QUERY__

#     PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

#     SELECT DISTINCT ?variable
#     WHERE {
#         ?program rdf:type sdtl:Program .
#         ?program sdtl:Commands/rdfs:member ?command .
#         ?command sdtl:Variable/sdtl:VariableName "Kelvin"^^<http://www.w3.org/2001/XMLSchema#string> .
#         ?command sdtl:OperatesOn/sdtl:VariableName/(^sdtl:VariableName/^sdtl:Variable/sdtl:OperatesOn/sdtl:VariableName)* ?variable .
#     } ORDER BY ?variable

# __END_QUERY__

# END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} Q6 "WHAT COMMANDS AFFECT EACH VARIABLE?" << END_SCRIPT

# geist query --format table << __END_QUERY__

#     {{{

#     {{ include "../../common/sdtl.g" }}

#     }}}

#     PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

#     SELECT DISTINCT ?affected_variable ?command ?source_line ?source_text
#     WHERE {
#         ?program rdf:type sdtl:Program .
#         ?program sdtl:Commands/rdfs:member ?command .
#         {
#             ?command sdtl:Variable ?variable .
#         }
#         UNION
#         {
#             ?command rdf:type sdtl:Load .
#             ?command sdtl:ProducesDataframe/sdtl:VariableInventory/rdfs:member ?variable .
#         }
#         ?variable sdtl:VariableName ?affected_variable .
#         {{ command_source "?command" "?source_line" "?source_text" }}
#     } ORDER BY ?affected_variable ?source_line

# __END_QUERY__

# END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} Q7 "WHAT COMMANDS READ VARIABLE VALUES ASSIGNED BY OTHER COMMANDS?" << END_SCRIPT

# geist query --format table << __END_QUERY__

#     {{{

#     {{ include "../../common/sdtl.g" }}

#     }}}

#     PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

#     SELECT DISTINCT ?variable ?writer_line ?writer_text ?reader_line ?reader_text
#     WHERE {
#         {{ program_command "?program" "?reader" }} .
#         {{ variable_write_read_edge "?variable" "?writer" "?reader" }} .
#         {{ command_source "?writer" "?writer_line" "?writer_text" }} .
#         {{ command_source "?reader" "?reader_line" "?reader_text" }}
#         FILTER ( ?variable = "A" )
#     } ORDER BY ?variable ?writer_line ?reader_line

# __END_QUERY__

# END_SCRIPT

# bash ${RUNNER} R1 "REPORT HISTORY OF EACH VARIABLE" << 'END_SCRIPT'

# geist report << '__END_REPORT_TEMPLATE__'

# {{{
#     {{ include "../../common/sdtl.g" }}
# }}}

# ---------- VARIABLE HISTORY REPORT ----------
#                                                                                                     \\
# {{ with $Program := sdtl_select_program | value }}                                                  \\
#     {{ range $SaveCommand := sdtl_select_save_commands $Program | rows }}
#         ===================
#         Dataframe {{ index $SaveCommand 2 }}
#         ===================
#         {{ range $VariableName := sdtl_select_dataframe_variables (index $SaveCommand 1) | vector }}
#             Variable {{ $VariableName }}
#             -------------------
#             {{ range $LoadCommand := sdtl_select_load_commands $Program $VariableName | rows }}     \\
#                 Load    | Line {{ index $LoadCommand 3 }} | {{ index $LoadCommand 4 }}
#             {{ end }}                                                                               \\
#             {{ range $UpdateCommand := sdtl_select_update_commands $Program $VariableName | rows }} \\
#                 Compute | Line {{ index $UpdateCommand 1 }} | {{ index $UpdateCommand 2 }}
#             {{ end }}                                                                               \\
#                 Save    | Line {{ index $SaveCommand 3 }} | {{ index $SaveCommand 4 }}
#         {{ end }}
#     {{ end }}
# {{ end }}

# __END_REPORT_TEMPLATE__

# END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} R2 "WHAT COMMANDS WRITE TO EACH VARIABLE?" << END_SCRIPT

# geist report << '__END_REPORT_TEMPLATE__'

# {{{
#     {{ include "../../common/sdtl.g" }}
# }}}

# ---------- COMMANDS THAT WRITE TO EACH VARIABLE ----------

# {{ select '''
#     SELECT DISTINCT ?written_variable ?source_line ?source_text
#     WHERE {
#         {{ program_command "?program" "?command" }} .
#         {{ variable_writer  "?command" "?written_variable" }} .
#         {{ command_source "?command" "?source_line" "?source_text" }}
#     }
#     ORDER BY ?written_variable ?source_line
# ''' | tabulate }}

# __END_REPORT_TEMPLATE__

# END_SCRIPT


# # *****************************************************************************

# bash ${RUNNER} R3 "WHAT COMMANDS READ FROM EACH VARIABLE?" << END_SCRIPT

# geist report << '__END_REPORT_TEMPLATE__'

# {{{
#     {{ include "../../common/sdtl.g" }}
# }}}

# ---------- COMMANDS THAT READ FROM EACH VARIABLE ----------

# {{ select '''
#     SELECT DISTINCT ?read_variable ?source_line ?source_text
#     WHERE {
#         {{ program_command "?program" "?command" }} .
#         {{ variable_reader  "?command" "?read_variable" }} .
#         {{ command_source "?command" "?source_line" "?source_text" }}
#     }
#     ORDER BY ?read_variable ?source_line
# ''' | tabulate }}

# __END_REPORT_TEMPLATE__

# END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} R4 "WHAT COMMANDS WRITE VARIABLES READ BY DOWNSTREAM COMMANDS?" << END_SCRIPT

# geist report << '__END_REPORT_TEMPLATE__'

# {{{
#     {{ include "../../common/sdtl.g" }}
# }}}

# ---------- COMMANDS THAT WRITE VARIABLES READ BY DOWNSTREAM COMMANDS ----------

# {{ select '''
#     SELECT DISTINCT ?variable ?writer_line ?writer_text ?reader_line ?reader_text
#     WHERE {
#         {{ program_command "?program" "?reader" }} .
#         {{ variable_reader  "?reader" "?variable" }} .
#         {{ upstream_variable_writer "?variable" "?reader" "?writer" }} .
#         {{ command_source "?writer" "?writer_line" "?writer_text" }} .
#         {{ command_source "?reader" "?reader_line" "?reader_text" }}
#     }
#     ORDER BY ?variable ?writer_line
# ''' | tabulate }}

# __END_REPORT_TEMPLATE__

# END_SCRIPT

# # *****************************************************************************

# bash ${RUNNER} R5 "WHAT COMMANDS READ VARIABLES WRITTEN BY MULTIPLE UPSTREAM COMMANDS?" << END_SCRIPT

# geist report << '__END_REPORT_TEMPLATE__'

# {{{
#     {{ include "../../common/sdtl.g" }}
# }}}

# ---------- COMMANDS THAT READ VARIABLES WRITTEN BY MULTIPLE UPSTREAM COMMANDS ----------

# {{ select '''
#     SELECT DISTINCT ?variable ?writer_line ?writer_text ?intermediate_writer_line
#         ?intermediate_writer_text  ?reader_line ?reader_text
#     WHERE {
#         {{ program_command "?program" "?reader" }} .
#         {{ variable_reader  "?reader" "?variable" }} .
#         {{ upstream_variable_writer "?variable" "?reader" "?intermediate_writer" }} .
#         {{ variable_reader  "?intermediate_writer" "?variable" }} .
#         {{ upstream_variable_writer "?variable" "?intermediate_writer" "?writer" }} .
#         {{ command_source "?writer" "?writer_line" "?writer_text" }} .
#         {{ command_source "?intermediate_writer" "?intermediate_writer_line" "?intermediate_writer_text" }} .
#         {{ command_source "?reader" "?reader_line" "?reader_text" }}
#     }
#     ORDER BY ?variable ?writer_line
# ''' | tabulate }}

# __END_REPORT_TEMPLATE__

# END_SCRIPT

# # # *****************************************************************************

# # bash ${GRAPHER} GRAPH-1 "DATAFRAME FLOW THROUGH COMMANDS" \
# #     << '__END_SCRIPT__'

# # geist report << '__END_REPORT_TEMPLATE__'

# #     {{{
# #         {{ include "../../common/graphviz.g" }}
# #         {{ include "../../common/sdtl.g" }}
# #     }}}

# #     {{ gv_graph "sdtl_program" }}

# #     {{ gv_title "Dataframe-flow through commands" }}

# #     {{ gv_cluster "program_graph" }}

# #     # command nodes
# #     {{ sdtl_program_node_style }}
# #     node[width=8]
# #     {{ with $ProgramID := sdtl_select_program | value }}                                    \\

# #         {{ range $Command := (sdtl_select_commands $ProgramID | rows ) }}                   \\
# #             {{ gv_labeled_node (index $Command 0) (index $Command 1) }}
# #         {{ end }}                                                                           \\

# #         # dataframe edges
# #         {{ range $Edge := (sdtl_select_dataframe_edges $ProgramID | rows) }}                \\
# #             {{ gv_labeled_edge (index $Edge 0) (index $Edge 1) (index $Edge 3) }}
# #         {{ end }}                                                                           \\
# #                                                                                             \\
# #     {{ end }}                                                                               \\
# #                                                                                             \\
# #     {{ gv_cluster_end }}

# #     {{ gv_end }}                                                                            \\

# # __END_REPORT_TEMPLATE__

# # __END_SCRIPT__


# # # *****************************************************************************

# # bash ${GRAPHER} GRAPH-2 "VARIABLE FLOW THROUGH COMMANDS" \
# #     << '__END_SCRIPT__'

# # geist report << '__END_REPORT_TEMPLATE__'

# #     {{{
# #         {{ include "../../common/graphviz.g" }}
# #         {{ include "../../common/sdtl.g" }}
# #     }}}

# #     {{ gv_graph "sdtl_program" }}

# #     {{ gv_title "Variable-flow through commands" }}

# #     {{ gv_cluster "program_graph" }}

# #     # command nodes
# #     {{ sdtl_program_node_style }}
# #     node[width=8]
# #     {{ with $ProgramID := sdtl_select_program | value }}                                    \\

# #         {{ range $Command := (sdtl_select_commands $ProgramID | rows ) }}                   \\
# #             {{ gv_labeled_node (index $Command 0) (index $Command 1) }}
# #         {{ end }}                                                                           \\

# #         # variable write->read edges
# #         {{ range $Edge := (sdtl_select_variable_write_read_edges $ProgramID | rows) }} \\
# #             {{ gv_labeled_edge (index $Edge 1) (index $Edge 2) (index $Edge 0) }}
# #         {{ end }}                                                                           \\

# #     {{ end }}                                                                               \\
# #                                                                                             \\
# #     {{ gv_cluster_end }}

# #     {{ gv_end }}                                                                            \\

# # __END_REPORT_TEMPLATE__

# # __END_SCRIPT__

