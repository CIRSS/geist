
{{ macro "gv_graph" '''

    digraph {{.}} {  {{nl}}
    rankdir=LR       {{nl}}

''' }}

{{ macro "gv_title" '''

    fontname=Courier; fontsize=18; labelloc=t   {{nl}}
    label="{{.}}"                               {{nl}}

''' }}

{{ macro "gv_end" '''

    }   {{nl}}

''' }}


