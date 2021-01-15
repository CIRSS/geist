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

    {{{

    {{ include "../../common/sdtl.g" }}

    }}}

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?used_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands/rdfs:member ?command .
        ?command sdtl:OperatesOn/sdtl:VariableName ?used_variable .
        {{ command_source "?command" "?source_line" "?source_text" }}
    } ORDER BY ?used_variable ?source_line

__END_QUERY__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q2 "WHAT VARIABLES DIRECTLY AFFECT OTHER VARIABLES?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    {{{

    {{ include "../../common/sdtl.g" }}

    }}}

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?affecting_variable  ?affected_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands/rdfs:member ?command .
        ?command sdtl:Variable/sdtl:VariableName ?affected_variable .
        ?command sdtl:OperatesOn/sdtl:VariableName ?affecting_variable .
        {{ command_source "?command" "?source_line" "?source_text" }}
    } ORDER BY ?affecting_variable ?affected_variable ?source_line

__END_QUERY__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q3 "WHAT VARIABLES DIRECTLY AFFECT THE KELVIN VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    {{{

    {{ include "../../common/sdtl.g" }}

    }}}

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?affecting_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands/rdfs:member ?command .
        ?command sdtl:Variable/sdtl:VariableName "Kelvin"^^<http://www.w3.org/2001/XMLSchema#string> .
        ?command sdtl:OperatesOn/sdtl:VariableName ?affecting_variable
        {{ command_source "?command" "?source_line" "?source_text" }}
    } ORDER BY ?affecting_variable ?source_line

__END_QUERY__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q4 "WHAT VARIABLES DIRECTLY AFFECT VARIABLES THAT DIRECTLY AFFECT THE KELVIN VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    {{{

    {{ include "../../common/sdtl.g" }}

    }}}

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
        {{ command_source "?indirectly_affecting_command" "?source_line" "?source_text" }}
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
        ?program sdtl:Commands/rdfs:member ?command .
        ?command sdtl:Variable/sdtl:VariableName "Kelvin"^^<http://www.w3.org/2001/XMLSchema#string> .
        ?command sdtl:OperatesOn/sdtl:VariableName/(^sdtl:VariableName/^sdtl:Variable/sdtl:OperatesOn/sdtl:VariableName)* ?variable .
    } ORDER BY ?variable

__END_QUERY__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} Q6 "WHAT COMMANDS AFFECT EACH VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    {{{

    {{ include "../../common/sdtl.g" }}

    }}}

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?affected_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands/rdfs:member ?command .
        {
            ?command sdtl:Variable ?variable .
        }
        UNION
        {
            ?command rdf:type sdtl:Load .
            ?command sdtl:ProducesDataframe/sdtl:VariableInventory/rdfs:member ?variable .
        }
        ?variable sdtl:VariableName ?affected_variable .
        {{ command_source "?command" "?source_line" "?source_text" }}
    } ORDER BY ?affected_variable ?source_line

__END_QUERY__

END_SCRIPT

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

bash ${RUNNER} R2 "WHAT COMMANDS WRITE TO EACH VARIABLE?" << END_SCRIPT

blazegraph report << '__END_REPORT_TEMPLATE__'

{{{
    {{ include "../../common/sdtl.g" }}
}}}

---------- COMMANDS THAT WRITE TO EACH VARIABLE ----------

{{ select '''
    SELECT DISTINCT ?written_variable ?source_line ?source_text
    WHERE {
        {{ program_command "?program" "?command" }} .
        {{ variable_writer  "?command" "?written_variable" }} .
        {{ command_source "?command" "?source_line" "?source_text" }}
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
        {{ program_command "?program" "?command" }} .
        {{ variable_reader  "?command" "?read_variable" }} .
        {{ command_source "?command" "?source_line" "?source_text" }}
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
    SELECT DISTINCT ?variable ?writer_line ?writer_text ?reader_line ?reader_text
    WHERE {
        {{ program_command "?program" "?reader" }} .
        {{ upstream_variable_writer "?variable" "?reader" "?writer" }} .
        {{ command_source "?writer" "?writer_line" "?writer_text" }} .
        {{ command_source "?reader" "?reader_line" "?reader_text" }}
    }
    ORDER BY ?variable ?writer_line
''' | tabulate }}

__END_REPORT_TEMPLATE__

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} R5 "WHAT COMMANDS READ VARIABLES WRITTEN BY MULTIPLE UPSTREAM COMMANDS?" << END_SCRIPT

blazegraph report << '__END_REPORT_TEMPLATE__'

{{{
    {{ include "../../common/sdtl.g" }}
}}}

---------- COMMANDS THAT READ VARIABLES WRITTEN BY MULTIPLE UPSTREAM COMMANDS ----------

{{ select '''
    SELECT DISTINCT ?variable ?writer_line ?writer_text ?intermediate_writer_line
        ?intermediate_writer_text  ?reader_line ?reader_text
    WHERE {
        {{ program_command "?program" "?reader" }} .
        {{ upstream_variable_writer "?variable" "?reader" "?intermediate_writer" }} .
        {{ upstream_variable_writer "?variable" "?intermediate_writer" "?writer" }} .
        {{ command_source "?writer" "?writer_line" "?writer_text" }} .
        {{ command_source "?intermediate_writer" "?intermediate_writer_line" "?intermediate_writer_text" }} .
        {{ command_source "?reader" "?reader_line" "?reader_text" }}
    }
    ORDER BY ?variable ?writer_line
''' | tabulate }}

__END_REPORT_TEMPLATE__

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


# *****************************************************************************

bash ${GRAPHER} GRAPH-2 "VARIABLE FLOW THROUGH COMMANDS" \
    << '__END_SCRIPT__'

blazegraph report << '__END_REPORT_TEMPLATE__'

    {{{
        {{ include "../../common/graphviz.g" }}
        {{ include "../../common/sdtl.g" }}
    }}}

    {{ gv_graph "sdtl_program" }}

    {{ gv_title "Variable-flow through commands" }}

    {{ gv_cluster "program_graph" }}

    # command nodes
    {{ sdtl_program_node_style }}
    node[width=8]
    {{ with $ProgramID := sdtl_select_program | value }}                                    \\

        {{ range $Command := (sdtl_select_commands $ProgramID | rows ) }}                   \\
            {{ gv_labeled_node (index $Command 0) (index $Command 1) }}
        {{ end }}                                                                           \\

        # variable write->read edges
        {{ range $Edge := (sdtl_select_variable_write_read_edges $ProgramID | rows) }} \\
            {{ gv_labeled_edge (index $Edge 0) (index $Edge 1) (index $Edge 2) }}
        {{ end }}                                                                           \\

    {{ end }}                                                                               \\
                                                                                            \\
    {{ gv_cluster_end }}

    {{ gv_end }}                                                                            \\

__END_REPORT_TEMPLATE__

__END_SCRIPT__

