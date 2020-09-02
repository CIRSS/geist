{{ prefix "prov" "http://www.w3.org/ns/prov#" }}
{{ prefix "provone" "http://purl.dataone.org/provone/2015/01/15/ontology#" }}
{{ prefix "wt" "http://wholetale.org/ontology/wt#" }}

{{ query "wt_select_run_id" '''
    SELECT ?r 
    WHERE {
        ?r a wt:TaleRun
    }
''' }}

{{ query "wt_select_tale_name" "RunID" '''
    SELECT ?n 
    WHERE {
        <{{$RunID}}> wt:TaleName ?n
    }
''' }}

{{ query "wt_select_run" '''
    SELECT ?runID ?taleName ?runScript
    WHERE {
        ?runID a wt:TaleRun .
        ?runID wt:TaleName ?taleName .
        ?runID wt:TaleRunScript ?runScript
    }
''' }}

{{ macro "wt_node_run" "Label" '''
    "{{$Label}}"
''' }}


{{ macro "wt_run_node_style" '''
    node[shape=box style="filled" fillcolor="#FFFFFF" peripheries=2 fontname=Courier]
''' }}

{{ macro "wt_node_style_file" '''
    
    node[shape=box style="rounded,filled" fillcolor="#FFFFCC" peripheries=1 fontname=Helvetica]

''' }}

{{ query "wt_select_tale_output_files" "RunScript" '''
    SELECT DISTINCT ?fileID ?filePath
    WHERE {
        ?e wt:ExecutionOf <{{$RunScript}}> .            
        ?p wt:ChildProcessOf ?e .   
        ?p wt:WroteFile ?fileID .          
        FILTER NOT EXISTS {
            ?_ wt:ReadFile ?fileID . 
        }
        ?fileID wt:FilePath ?filePath .
    }
    ORDER BY ?filePath
''' }}
}}}

{{ query "wt_select_tale_input_files" "RunScript" '''
    SELECT DISTINCT ?f ?fp
    WHERE {
        ?e wt:ExecutionOf <{{$RunScript}}> .               
        ?p wt:ChildProcessOf ?e .   
        ?p wt:ReadFile ?f .          
        FILTER NOT EXISTS {
            ?_ wt:WroteFile ?f . 
        }
        ?f wt:FilePath ?fp .
    }
    ORDER BY ?fp
''' }}
}}}

{{ query "SelectTaleOutputEdges" '''
    SELECT DISTINCT ?taleID ?filePath
    WHERE {
        ?e wt:ExecutionOf <{{.}}> .            
        ?p wt:ChildProcessOf ?e .   
        ?p wt:WroteFile ?fileID .          
        FILTER NOT EXISTS {
            ?_ wt:ReadFile ?fileID . 
        }
        ?fileID wt:FilePath ?filePath .
    }
    ORDER BY ?filePath
''' }}
}}}