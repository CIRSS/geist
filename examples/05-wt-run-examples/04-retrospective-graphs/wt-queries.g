{{ prefix "prov" "http://www.w3.org/ns/prov#" }}
{{ prefix "provone" "http://purl.dataone.org/provone/2015/01/15/ontology#" }}
{{ prefix "wt" "http://wholetale.org/ontology/wt#" }}

{{ query "GetRunID" '''

    SELECT ?r 
    WHERE {
        ?r a wt:TaleRun
    }

''' }}

{{ query "GetTaleName" '''

    SELECT ?n 
    WHERE {
        <{{.}}> wt:TaleName ?n
    }

''' }}
