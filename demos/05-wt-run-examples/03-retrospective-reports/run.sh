#!/usr/bin/env bash

# *****************************************************************************

bash_cell SETUP << END_CELL

# IMPORT PROVONE TRACE

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/branched-pipeline.jsonld

END_CELL

# *****************************************************************************

bash_cell REPORT-1 << '__END_CELL__'

# WHAT DATA FILES WERE USED AS INPUT BY THE TALE?

geist report << '__END_REPORT_TEMPLATE__'
                                                                                \\
{{ prefix "prov" "http://www.w3.org/ns/prov#" }}                                \\
{{ prefix "provone" "http://purl.dataone.org/provone/2015/01/15/ontology#" }}   \\
{{ prefix "wt" "http://wholetale.org/ontology/wt#" }}                           \\
                                                                                \\
{{ select '''

    SELECT DISTINCT ?tale_input_file_path ?read_file
    WHERE {
        ?run rdf:type wt:TaleRun .
        ?run wt:TaleRunScript ?run_script .
        ?run_process wt:ExecutionOf ?run_script .
        ?run_sub_process (wt:ChildProcessOf)+ ?run_process .
        ?run_sub_process wt:ReadFile ?read_file .
        FILTER NOT EXISTS {
            ?_ wt:WroteFile ?read_file . }
        ?read_file wt:FilePath ?tale_input_file_path .
    }
    ORDER BY ?tale_input_file_path                                              
                                                                                
''' | tabulate }}                                                               \\
                                                                                \\
__END_REPORT_TEMPLATE__

__END_CELL__


bash_cell REPORT-2 << '__END_CELL__'

# WHAT DATA FILES WERE USED AS INPUT BY THE TALE?

geist report << '__END_REPORT_TEMPLATE__'
                                                                                                        \\
{{ prefix "prov" "http://www.w3.org/ns/prov#" }}                                                        \\
{{ prefix "provone" "http://purl.dataone.org/provone/2015/01/15/ontology#" }}                           \\
{{ prefix "wt" "http://wholetale.org/ontology/wt#" }}                                                   \\
                                                                                                        \\
{{ with $Run := (select "SELECT ?r WHERE {?r a wt:TaleRun}") | value }}                                 \\
                                                                                                        \\
    Tale Run:   {{ $Run }}
    Tale Name:  {{ (select "SELECT ?n WHERE {<{{.}}> wt:TaleName ?n}" $Run | value) }}
                                                                                                        \\
    {{ with $RunScript := (select "SELECT ?s WHERE {<{{.}}> wt:TaleRunScript ?s}" $Run | value) }}      \\
        Tale Script: {{ (select "SELECT ?n WHERE {<{{.}}> wt:FilePath ?n}" $RunScript | value) }}

    Tale Inputs:
        {{ range $InputFile := (select '''
            SELECT DISTINCT ?fp WHERE {
                ?e wt:ExecutionOf <{{.}}> .
                ?p (wt:ChildProcessOf)+ ?e .
                ?p wt:ReadFile ?f .
                FILTER NOT EXISTS {
                    ?_ wt:WroteFile ?f . }
                ?f wt:FilePath ?fp .
            } ORDER BY ?fp''' $RunScript | vector) }}                                                   \\
                {{ $InputFile }}
        {{end}}                                                                                         \\
    {{end}}                                                                                             \\
{{end}}                                                                                                 \\
                                                                                                        \\
__END_REPORT_TEMPLATE__

__END_CELL__

bash_cell REPORT-3 << '__END_CELL__'

# WHAT DATA FILES WERE USED AS INPUT BY THE TALE?

geist report << '__END_REPORT_TEMPLATE__'
                                                                                                        \\
{{{
{{ prefix "prov" "http://www.w3.org/ns/prov#" }}
{{ prefix "provone" "http://purl.dataone.org/provone/2015/01/15/ontology#" }}
{{ prefix "wt" "http://wholetale.org/ontology/wt#" }}
                                                                                                        
{{ query "GetRunID" '''
    SELECT ?r
    WHERE {
        ?r a wt:TaleRun
    }
''' }}

{{ query "GetTaleName" "RunID" '''
    SELECT ?n
    WHERE {
        <{{$RunID}}> wt:TaleName ?n
    }
''' }}

{{ query "GetRunScriptID" "RunID" '''
    SELECT ?s
    WHERE {
        <{{$RunID}}> wt:TaleRunScript ?s
    }
''' }}

{{ query "GetFilePath" "FileID" '''
    SELECT ?n
    WHERE {
        <{{$FileID}}> wt:FilePath ?n
    }
''' }}

{{ query "GetInputFilePaths" "RunScriptID" '''
    SELECT DISTINCT ?fp
    WHERE {
        ?e wt:ExecutionOf <{{$RunScriptID}}> .
        ?p (wt:ChildProcessOf)+ ?e .
        ?p wt:ReadFile ?f .
        FILTER NOT EXISTS {
            ?_ wt:WroteFile ?f .
        }
        ?f wt:FilePath ?fp .
    }
    ORDER BY ?fp
''' }}
}}}                                                             \\
                                                                \\
{{ with $RunID := GetRunID | value }}                           \\
                                                                \\
    Tale Run:    {{ $RunID }}
    Tale Name:   {{ (GetTaleName $RunID | value) }}
                                                                \\
    {{ with $RunScriptID := (GetRunScriptID $RunID | value) }}  \\
                                                                \\
    Tale Script: {{ (GetFilePath $RunScriptID | value) }}

    Tale Inputs:
        {{ range (GetInputFilePaths $RunScriptID | vector) }}   \\
            {{ . }}
        {{ end }}                                               \\
    {{ end }}                                                   \\
{{ end }}                                                       \\
                                                                \\
__END_REPORT_TEMPLATE__

__END_CELL__
