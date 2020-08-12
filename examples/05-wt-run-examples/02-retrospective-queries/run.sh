#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT PROVONE TRACE" << END_SCRIPT

blazegraph drop
blazegraph import --format jsonld --file ../data/branched-pipeline.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} RETROSPECTIVE-1 "WHAT DATA FILES WERE USED AS INPUT BY THE TALE?" \
    << END_SCRIPT

blazegraph select --format table << END_QUERY

    prefix prov: <http://www.w3.org/ns/prov#>
    prefix provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>
    prefix wt: <http://wholetale.org/ontology/wt#>

    SELECT DISTINCT ?read_file ?tale_input_file_path
    WHERE {
        ?run rdf:type wt:TaleRun .                     
        ?run wt:TaleRunScript ?run_script .         
        ?run_script_process wt:ExecutionOf ?run_script .             
        ?run_sub_process wt:ParentProcess ?run_script_process .             
        ?run_sub_process wt:ReadFile ?read_file .
        FILTER NOT EXISTS { ?tale_process wt:WroteFile ?read_file . }
        ?read_file wt:FilePath ?tale_input_file_path .
    }
    ORDER BY ?tale_input_file_path

END_QUERY

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} RETROSPECTIVE-1 "WHAT FILES WERE PRODUCED AS OUTPUTS OF THE TALE?" \
    << END_SCRIPT

blazegraph select --format table << END_QUERY

    prefix prov: <http://www.w3.org/ns/prov#>
    prefix provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>
    prefix wt: <http://wholetale.org/ontology/wt#>

    SELECT ?tale_out_file_path
    WHERE {
        ?execution rdf:type provone:Execution .         # find the Tale execution
        ?execution prov:generated ?tale_output .        # find all outputs of the Tale
        ?tale_output rdf:type wt:DataFile .             # select just the outputs that are data files
        ?tale_output wt:FilePath ?tale_out_file_path .  # get the file path for each input data file
    }

END_QUERY

END_SCRIPT
