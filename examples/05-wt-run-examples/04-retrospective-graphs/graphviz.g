
{{ macro "gv_graph" '''
    digraph {{.}} { 
    rankdir=LR
''' }}

{{ macro "gv_title" '''
    fontname=Courier; fontsize=18; labelloc=t
    label="{{.}}"
''' }}

{{ macro "gv_end" '''
    }
''' }}

{{ macro "gv_cluster" '''
    subgraph {{ printf "cluster_%s" . }} { label=""; color=white; penwidth=0
    subgraph {{ printf "cluster_%s_inner" . }} { label=""; color=white
''' }}

{{ macro "gv_cluster_end" '''
    }}
''' }}

{{ macro "labeled_node" '''
    {{with $args := .}}
    {{with $node_id := index $args 0}}
    {{with $node_label := index $args 1}}
        "{{$node_id}}" [label="{{$node_label}}"]
    {{end}}{{end}}{{end}}
''' }}

{{ macro "gv_edge" '''
    {{with $args := .}}
        {{with $tail:= "08-branched-pipeline" }}
            {{with $head := index $args 0}}
                "{{$tail}}" -> "{{$head}}"
            {{end}}
        {{end}}
    {{end}}
''' }}

{{ macro "gv_input_edge" '''
    {{with $args := .}}
        {{with $head:= "08-branched-pipeline" }}
            {{with $tail := index $args 0}}
                "{{$tail}}" -> "{{$head}}"
            {{end}}
        {{end}}
    {{end}}
''' }}