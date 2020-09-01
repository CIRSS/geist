{{ prefix "prov" "http://www.w3.org/ns/prov#" }}
{{ prefix "provone" "http://purl.dataone.org/provone/2015/01/15/ontology#" }}
{{ prefix "wt" "http://wholetale.org/ontology/wt#" }}

{{ query "SelectRunID" '''
    SELECT ?r 
    WHERE {
        ?r a wt:TaleRun
    }
''' }}

{{ query "SelectTaleName" '''
    SELECT ?n 
    WHERE {
        <{{.}}> wt:TaleName ?n
    }
''' }}

{{ macro "wt_run_node_style" '''
    node[shape=box style="filled" fillcolor="#FFFFFF" peripheries=2 fontname=Courier]
''' }}

{{ macro "wt_node_run" '''
    "{{.}}"
''' }}

{{ macro "wt_node_style_file" '''
    
    node[shape=box style="rounded,filled" fillcolor="#FFFFCC" peripheries=1 fontname=Helvetica]

''' }}

{{ query "SelectTaleOutputFilePaths" '''
    SELECT DISTINCT ?f ?fp
    WHERE {
        ?e wt:ExecutionOf <{{.}}> .               
        ?p wt:ChildProcessOf ?e .   
        ?p wt:WroteFile ?f .          
        FILTER NOT EXISTS {
            ?_ wt:ReadFile ?f . 
        }
        ?f wt:FilePath ?fp .
    }
    ORDER BY ?fp
''' }}
}}}

