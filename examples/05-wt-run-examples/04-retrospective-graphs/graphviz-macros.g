
{{ macro "gv_graph" '''

    digraph {{.}} {  \\n
    rankdir=LR       \\n

''' }}

{{ macro "gv_title" '''

    fontname=Courier; fontsize=18; labelloc=t   \\n
    label="{{.}}"                               \\n

''' }}

{{ macro "gv_end" '''

    }   \\n

''' }}


