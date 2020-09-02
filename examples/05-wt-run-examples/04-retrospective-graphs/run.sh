#!/usr/bin/env bash

DOT_RUNNER='../../common/run_dot_examples.sh'
SCRIPT_RUNNER='../../common/run_script_example.sh'


# *****************************************************************************

bash ${SCRIPT_RUNNER} SETUP "IMPORT PROVONE TRACE" << END_SCRIPT

blazegraph drop
blazegraph import --format jsonld --file ../data/branched-pipeline.jsonld

END_SCRIPT


bash ${DOT_RUNNER} GRAPH-1 "EMPTY DOT FILE" \
    << '__END_SCRIPT__'

blazegraph report << '__END_REPORT_TEMPLATE__'

{{{
    {{ include "graphviz.g" }}
}}}

    # A graphviz file
    {{ gv_graph "wt_run" }}
    {{ gv_end }}

__END_REPORT_TEMPLATE__

__END_SCRIPT__


bash ${DOT_RUNNER} GRAPH-2 "TITLED EMPTY DOT FILE" \
    << '__END_SCRIPT__'

blazegraph report << '__END_REPORT_TEMPLATE__'

{{{
    {{ include "graphviz.g" }}
    {{ include "wt.g" }}
}}}

    {{ with $RunID := wt_select_run_id | value}}

        # Run ID: {{ $RunID }}
        {{ gv_graph "wt_run" }} 
        {{ gv_title (wt_select_tale_name $RunID | value) }}
        {{ gv_end }}
    
    {{ end }}

__END_REPORT_TEMPLATE__

__END_SCRIPT__


bash ${DOT_RUNNER} GRAPH-3 "Node for Tale Run" \
    << '__END_SCRIPT__'

blazegraph report << '__END_REPORT_TEMPLATE__'

    {{{
        {{ include "graphviz.g" }}
        {{ include "wt.g" }}
    }}}

    {{ with $RunID := wt_select_run_id | value }}

        # Run ID: {{ $RunID }}
        {{ gv_graph "wt_run" }} 
        {{ gv_title (wt_select_tale_name $RunID | value) }}
        {{ wt_run_node_style }}
        {{ wt_node_run $RunID }}
        {{ gv_end }}
    
    {{ end }}

__END_REPORT_TEMPLATE__

__END_SCRIPT__


bash ${DOT_RUNNER} GRAPH-4 "Node for Tale Run" \
    << '__END_SCRIPT__'

blazegraph report << '__END_REPORT_TEMPLATE__'

    {{{
        {{ include "graphviz.g" }}
        {{ include "wt.g" }}
    }}}

    {{ with $Run := wt_select_run | vector }}                                                   \\
        {{ with $RunID := index $Run 0 }}                                                       \\
            {{ with $TaleName := index $Run 1 }}                                                \\
                {{ with $RunScript := index $Run 2 }}                                           \\
 
                    # Run ID: {{ $RunID }}
                    {{ gv_graph "wt_run" }}

                    # graph title
                    {{ gv_title "Tale Inputs and Outputs" }}
                    
                    # the tale run
                    {{ wt_run_node_style }}
                    {{ gv_labeled_node $RunID $TaleName }}
                    
                    # output files
                    {{ with $OutputFiles := (wt_select_tale_output_files $RunScript | rows) }}  \\
                        {{ wt_file_nodes_cluster "outputs" $OutputFiles }}
                        {{ wt_out_file_edges $RunID $OutputFiles }}                             \\
                    {{ end }}                                                                   

                    # input files
                    {{ with $InputFiles := (wt_select_tale_input_files $RunScript | rows) }}    \\
                        {{ wt_file_nodes_cluster "inputs" $InputFiles }}
                        {{ wt_in_file_edges $RunID $InputFiles }}                                      
                    {{ end }}                                                                   \\
                                                                                                \\
                    {{ gv_end }}
    
                {{ end }}
            {{ end }}
        {{ end }}
    {{ end }}

__END_REPORT_TEMPLATE__

__END_SCRIPT__