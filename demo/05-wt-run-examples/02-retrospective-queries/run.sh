#!/usr/bin/env bash

# *****************************************************************************

bash_cell SETUP "IMPORT PROVONE TRACE" << END_CELL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/branched-pipeline.jsonld

END_CELL

# *****************************************************************************

bash_cell RETROSPECTIVE-1 "WHAT DATA FILES WERE USED AS INPUT BY THE TALE?" \
    << __END_CELL__

geist query --format table << __END_QUERY__

    PREFIX prov: <http://www.w3.org/ns/prov#>
    PREFIX provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>
    PREFIX wt: <http://wholetale.org/ontology/wt#>

    SELECT DISTINCT ?tale_input_file_path ?read_file
    WHERE {
        ?run rdf:type wt:TaleRun .                              # Identify the Tale run described by this JSON-LD document.
        ?run wt:TaleRunScript ?run_script .                     # Identify the script used to run the Tale as a whole.
        ?run_process wt:ExecutionOf ?run_script .               # Identify the process that is the execution of that run script.
        ?run_sub_process (wt:ChildProcessOf)+ ?run_process .    # Find all child processes of the run script execution.
        ?run_sub_process wt:ReadFile ?read_file .               # Identify files read by those run subprocesses.
        FILTER NOT EXISTS {                                     # Filter out any files written by the Tale run, leaving
            ?tale_process wt:WroteFile ?read_file . }           #   just the input files.
        ?read_file wt:FilePath ?tale_input_file_path .          # Get the file path for each of the input files.
    }
    ORDER BY ?tale_input_file_path

__END_QUERY__

__END_CELL__


# *****************************************************************************

bash_cell RETROSPECTIVE-2 "WHAT FILES WERE PRODUCED AS OUTPUTS OF THE TALE?" \
    << __END_CELL__

geist query --format table << __END_QUERY__

    PREFIX prov: <http://www.w3.org/ns/prov#>
    PREFIX provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>
    PREFIX wt: <http://wholetale.org/ontology/wt#>

    SELECT DISTINCT ?tale_output_file_path ?written_file
    WHERE {
        ?run rdf:type wt:TaleRun .                          # Identify the Tale run described by this JSON-LD document.
        ?run wt:TaleRunScript ?run_script .                 # Identify the script used to run the Tale as a whole.
        ?run_process wt:ExecutionOf ?run_script .           # Identify the process that is the execution of that run script.
        ?run_sub_process (wt:ChildProcessOf)+  ?run_process .    # Find all child processes of the run script execution.
        ?run_sub_process wt:WroteFile ?written_file .       # Identify files written by those run subprocesses.
        FILTER NOT EXISTS { ?written_file                   # Filter out intermediate products of the Tale run, leaving
            wt:FileRole wt:TaleIntermediateData . }         #   just the finial output input files.
        ?written_file wt:FilePath ?tale_output_file_path .  # Get the file path for each of the output files.
    }

__END_QUERY__

__END_CELL__
