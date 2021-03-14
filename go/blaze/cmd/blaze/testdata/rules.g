
{{ macro "foo" '''
    <:foo>
''' }}

{{ macro "bar" "Sub" "Obj" '''
    {{_subject $Sub}} <:bar> {{_object $Obj}}
''' }}

{{ rule "foo_bar_baz" "s" "o" '''
    {{_subject $s}} {{foo}} ?y .
    {{bar "?y" "?z"}} .
    ?z <:baz> {{_object $o}} .
''' }}

