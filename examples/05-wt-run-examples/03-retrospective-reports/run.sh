#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT PROVONE TRACE" << END_SCRIPT

blazegraph drop
blazegraph import --format jsonld --file ../data/branched-pipeline.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} REPORT-1 "WHAT DATA FILES WERE USED AS INPUT BY THE TALE?" \
    << __END_SCRIPT__

blazegraph report << __END_REPORT_TEMPLATE__

{{ prefix "prov" "http://www.w3.org/ns/prov#" }}
{{ prefix "provone" "http://purl.dataone.org/provone/2015/01/15/ontology#" }}
{{ prefix "wt" "http://wholetale.org/ontology/wt#" }}

{{ select '''

    SELECT DISTINCT ?tale_input_file_path ?read_file
    WHERE {
        ?run rdf:type wt:TaleRun .                          
        ?run wt:TaleRunScript ?run_script .                 
        ?run_process wt:ExecutionOf ?run_script .               
        ?run_sub_process wt:ChildProcessOf ?run_process .   
        ?run_sub_process wt:ReadFile ?read_file .          
        FILTER NOT EXISTS {                               
            ?_ wt:WroteFile ?read_file . }     
        ?read_file wt:FilePath ?tale_input_file_path .
    }
    ORDER BY ?tale_input_file_path

''' | tabulate }}

__END_REPORT_TEMPLATE__

__END_SCRIPT__


bash ${RUNNER} REPORT-2 "WHAT DATA FILES WERE USED AS INPUT BY THE TALE?" \
    << '__END_SCRIPT__'

blazegraph report << '__END_REPORT_TEMPLATE__'

{{ prefix "prov" "http://www.w3.org/ns/prov#" }}
{{ prefix "provone" "http://purl.dataone.org/provone/2015/01/15/ontology#" }}
{{ prefix "wt" "http://wholetale.org/ontology/wt#" }}

{{ with $Run := (select "SELECT ?r WHERE {?r a wt:TaleRun}") | value }}

    {{ println "Tale Run:   " $Run }}

    {{ println "Tale Name:  " (select '''
            SELECT ?n WHERE {<{{.}}> wt:TaleName ?n}''' $Run | value) }}

    {{ with $RunScript := (select '''
            SELECT ?s WHERE {<{{.}}> wt:TaleRunScript ?s}''' $Run | value) }}

        {{ println "Tale Script:" (select '''
            SELECT ?n WHERE {<{{.}}> wt:FilePath ?n}''' $RunScript | value) }}

        {{ println }}
        {{ println "Tale Inputs:" }} 
        {{ range $InputFile := (select '''
            SELECT DISTINCT ?fp WHERE {
                ?e wt:ExecutionOf <{{.}}> .               
                ?p wt:ChildProcessOf ?e .   
                ?p wt:ReadFile ?f .          
                FILTER NOT EXISTS {                               
                    ?_ wt:WroteFile ?f . }     
                ?f wt:FilePath ?fp .
            } ORDER BY ?fp''' $RunScript | vector) }}

            {{println $InputFile}}

        {{end}}
    {{end}}
{{end}}

__END_REPORT_TEMPLATE__

__END_SCRIPT__

bash ${RUNNER} REPORT-3 "WHAT DATA FILES WERE USED AS INPUT BY THE TALE?" \
    << '__END_SCRIPT__'

blazegraph report << '__END_REPORT_TEMPLATE__'

{{ prefix "prov" "http://www.w3.org/ns/prov#" }}
{{ prefix "provone" "http://purl.dataone.org/provone/2015/01/15/ontology#" }}
{{ prefix "wt" "http://wholetale.org/ontology/wt#" }}

{{ query "GetRunID" '''
    SELECT ?r 
    WHERE {
        ?r a wt:TaleRun
    }
''' }}

{{ query "GetTaleName" '''
    SELECT ?n 
    WHERE {
        <{{.}}> wt:TaleName ?n
    }
''' }}

{{ query "GetRunScriptID" '''
    SELECT ?s 
    WHERE {
        <{{.}}> wt:TaleRunScript ?s
    }
''' }}

{{ query "GetFilePath" '''
    SELECT ?n
    WHERE {
        <{{.}}> wt:FilePath ?n
    }
''' }}

{{ query "GetInputFilePaths" '''
    SELECT DISTINCT ?fp
    WHERE {
        ?e wt:ExecutionOf <{{.}}> .               
        ?p wt:ChildProcessOf ?e .   
        ?p wt:ReadFile ?f .          
        FILTER NOT EXISTS {
            ?_ wt:WroteFile ?f . 
        }
        ?f wt:FilePath ?fp .
    }
    ORDER BY ?fp
''' }}

{{ with $RunID := runquery "GetRunID" | value }}
    {{ println "Tale Run:   " $RunID }}
    {{ println "Tale Name:  " (runquery "GetTaleName" $RunID | value) }}
    {{ with $RunScriptID := (runquery "GetRunScriptID" $RunID | value) }}
        {{ println "Tale Script:" (runquery "GetFilePath" $RunScriptID | value) }}
        {{ println }}
        {{ println "Tale Inputs:" }} 
        {{ range (runquery "GetInputFilePaths" $RunScriptID | vector) }}
            {{ println . }}
        {{ end }}
    {{ end }}
{{ end }}

__END_REPORT_TEMPLATE__

__END_SCRIPT__