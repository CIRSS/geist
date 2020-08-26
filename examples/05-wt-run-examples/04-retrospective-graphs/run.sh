#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT PROVONE TRACE" << END_SCRIPT

blazegraph drop
blazegraph import --format jsonld --file ../data/branched-pipeline.jsonld

END_SCRIPT


bash ${RUNNER} GRAPH-1 "EMPTY DOT FILE" \
    << '__END_SCRIPT__'

blazegraph report << '__END_REPORT_TEMPLATE__'

{{{
    {{ include "graphviz-macros.g" }}
}}}

    % A graphviz file
    {{ gv_graph "wt_run" }}
    {{ gv_end }}

__END_REPORT_TEMPLATE__

__END_SCRIPT__



bash ${RUNNER} GRAPH-2 "TITLED EMPTY DOT FILE" \
    << '__END_SCRIPT__'

blazegraph report << '__END_REPORT_TEMPLATE__'

    {{{
        {{ include "graphviz-macros.g" }}
        {{ include "wt-queries.g" }}
    }}}

    {{ with $RunID := GetRunID | value}}
        % Run ID: {{ $RunID }}
        {{ gv_graph "wt_run" }} 
        {{ gv_title (GetTaleName $RunID | value) }} \
    {{ end }}
    {{ gv_end }}

__END_REPORT_TEMPLATE__

__END_SCRIPT__