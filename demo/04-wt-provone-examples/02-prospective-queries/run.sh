#!/usr/bin/env bash

# *****************************************************************************

bash_cell SETUP "IMPORT PROVONE TRACE" << END_CELL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --format jsonld --file ../data/branched-pipeline.jsonld

END_CELL

# *****************************************************************************

bash_cell RETROSPECTIVE-1 "WHAT FILES WERE PROVIDED AS INPUT TO THE TALE?" \
    << END_CELL

geist query --format table << END_QUERY

    prefix prov: <http://www.w3.org/ns/prov#>
    prefix provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>
    prefix wt: <http://wholetale.org/ontology/wt#>

    SELECT ?tale_input_file_path
    WHERE {
        ?execution rdf:type provone:Execution .         # find the Tale execution
        ?execution prov:used ?tale_input .              # find all inputs to the Tale
        ?tale_input rdf:type wt:DataFile .              # select just the inputs that are data files
        ?tale_input wt:FilePath ?tale_input_file_path . # get the file path for each input data file
    }

END_QUERY

END_CELL

# *****************************************************************************

bash_cell RETROSPECTIVE-1 "WHAT FILES WERE PRODUCED AS OUTPUTS OF THE TALE?" \
    << END_CELL

geist query --format table << END_QUERY

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

END_CELL
