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


