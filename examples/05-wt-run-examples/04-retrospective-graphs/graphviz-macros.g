
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


