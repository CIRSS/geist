#!/usr/bin/env bash

# *****************************************************************************

bash_cell SETUP << END_CELL

# IMPORT PROVONE TRACE

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/branched-pipeline.jsonld

END_CELL


bash_dot_cell GRAPH-1 << '__END_CELL__'

# EMPTY DOT FILE

geist report << '__END_REPORT_TEMPLATE__'
                                                                    \\
{{{
    {{ include "graphviz.g" }}
}}}
                                                                    \\
    # A graphviz file
    {{ gv_graph "wt_run" }}
    {{ gv_end }}
                                                                    \\
__END_REPORT_TEMPLATE__

__END_CELL__


bash_dot_cell GRAPH-2 << '__END_CELL__'

# TITLED EMPTY DOT FILE

geist report << '__END_REPORT_TEMPLATE__'
                                                                    \\
{{{
    {{ include "graphviz.g" }}
    {{ include "wt.g" }}
}}}
                                                                    \\
    {{ with $RunID := wt_select_run | value}}

        # Run ID: {{ $RunID }}
        {{ gv_graph "wt_run" }}
        {{ gv_title (wt_select_tale_name $RunID | value) }}
        {{ gv_end }}

    {{ end }}
                                                                    \\
__END_REPORT_TEMPLATE__

__END_CELL__


bash_dot_cell GRAPH-3 << '__END_CELL__'

# Node for Tale Run

geist report << '__END_REPORT_TEMPLATE__'
                                                                    \\
    {{{
        {{ include "graphviz.g" }}
        {{ include "wt.g" }}
    }}}
                                                                    \\
    {{ with $RunID := wt_select_run | value }}

        # Run ID: {{ $RunID }}
        {{ gv_graph "wt_run" }}
        {{ gv_title "Tale Run" }}
        {{ wt_run_node $RunID }}
        {{ gv_end }}

    {{ end }}
                                                                    \\
__END_REPORT_TEMPLATE__

__END_CELL__


bash_dot_cell GRAPH-4 << '__END_CELL__'

# Tale Run with Inputs and Outputs

geist report << '__END_REPORT_TEMPLATE__'
                                                                                \\
    {{{
        {{ include "graphviz.g" }}
        {{ include "wt.g" }}
    }}}                                                                         \\
                                                                                \\
    {{ with $RunID := wt_select_run | value }}                                  \\

        # Run ID: {{ $RunID }}
        {{ gv_graph "wt_run" }}

        # graph title
        {{ gv_title "Tale Inputs and Outputs" }}

        # the tale run
        {{ wt_run_node $RunID }}

        # output files
        {{ with $OutputFiles := (wt_select_tale_output_files $RunID | rows) }}  \\
            {{ wt_file_nodes_cluster "outputs" $OutputFiles }}
            {{ wt_out_file_edges $RunID $OutputFiles }}                         \\
        {{ end }}

        # input files
        {{ with $InputFiles := (wt_select_tale_input_files $RunID | rows) }}    \\
            {{ wt_file_nodes_cluster "inputs" $InputFiles }}
            {{ wt_in_file_edges $RunID $InputFiles }}
        {{ end }}                                                               \\
                                                                                \\
        {{ gv_end }}
                                                                                \\
    {{ end }}                                                                   \\
                                                                                \\
__END_REPORT_TEMPLATE__

__END_CELL__


bash_dot_cell GRAPH-5 << '__END_CELL__'

# Tale Processes and Data Files

geist report << '__END_REPORT_TEMPLATE__'
                                                                                \\
    {{{
        {{ include "graphviz.g" }}
        {{ include "wt.g" }}
    }}}
                                                                                \\
    {{ with $RunID := wt_select_run | value }}                                  \\

        # Run ID: {{ $RunID }}
        {{ gv_graph "wt_run" }}

        # graph title
        {{ gv_title "Tale Processes and Data Files" }}

        {{ gv_cluster "Processes" }}

        # data files
        {{ wt_data_file_nodes $RunID }}

        # processes
        {{ wt_process_nodes $RunID }}

        # process input data file edges
        {{ wt_process_input_data_file_edges $RunID }}

        # process output data file edges
        {{ wt_process_output_data_file_edges $RunID }}

        {{ gv_cluster_end }}

        {{ gv_end }}
                                                                                \\
    {{ end }}                                                                   \\
                                                                                \\
__END_REPORT_TEMPLATE__

__END_CELL__
