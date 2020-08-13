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
        ?run_sub_process wt:ParentProcess ?run_process .   
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

{{ range (select '''

    SELECT DISTINCT ?tale_input_file_path
    WHERE {
        ?run rdf:type wt:TaleRun .                          
        ?run wt:TaleRunScript ?run_script .                 
        ?run_process wt:ExecutionOf ?run_script .               
        ?run_sub_process wt:ParentProcess ?run_process .   
        ?run_sub_process wt:ReadFile ?read_file .          
        FILTER NOT EXISTS {                               
            ?_ wt:WroteFile ?read_file . }     
        ?read_file wt:FilePath ?tale_input_file_path .
    }
    ORDER BY ?tale_input_file_path

''' | vector) }}
    {{println .}}
{{end}}

__END_REPORT_TEMPLATE__

__END_SCRIPT__

