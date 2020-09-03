{{ prefix "prov" "http://www.w3.org/ns/prov#" }}
{{ prefix "provone" "http://purl.dataone.org/provone/2015/01/15/ontology#" }}
{{ prefix "wt" "http://wholetale.org/ontology/wt#" }}

{{ query "wt_select_run" '''
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

{{ macro "wt_run_node_style" '''
    node[shape=box style="filled" fillcolor="#FFFFFF" peripheries=2 fontname=Courier]
''' }}

{{ macro "wt_run_node" "RunID" '''
    {{wt_run_node_style}}
    {{with $TaleName := (wt_select_tale_name $RunID | value) }} \\
        {{ gv_labeled_node $RunID $TaleName }}
    {{end}}
''' }}

{{ macro "wt_node_style_file" '''
    
    node[shape=box style="rounded,filled" fillcolor="#FFFFCC" peripheries=1 fontname=Helvetica]

''' }}

{{ query "wt_select_tale_output_files" "RunID" '''
    SELECT DISTINCT ?fileID ?filePath
    WHERE {
        $RunID wt:TaleRunScript ?runScript .
        ?e wt:ExecutionOf ?$runScript .            
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

{{ query "wt_select_tale_input_files" "RunID" '''
    SELECT DISTINCT ?f ?fp
    WHERE {
        $RunID wt:TaleRunScript ?runScript .
        ?e wt:ExecutionOf ?$runScript .            
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

{{ macro "wt_file_nodes_cluster" "ClusterName" "Files" '''
    {{ gv_cluster $ClusterName }}
        {{ wt_node_style_file }}
        {{ range $File := $Files }}                                 \\              
            {{ gv_labeled_node (index $File 0) (index $File 1) }} 
        {{ end }}                                                   \\                       
    {{ gv_cluster_end }}
''' }}

{{ macro "wt_out_file_edges" "SourceNode" "Files" '''
    {{ range $File := $Files }}                                     \\
        {{ gv_edge $SourceNode (index $File 0) }} 
    {{ end }}  
''' }}

{{ macro "wt_in_file_edges" "SinkNode" "Files" '''
    {{ range $File := $Files }}                                     \\
        {{ gv_edge (index $File 0) $SinkNode }} 
    {{ end }}  
''' }}