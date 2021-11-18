

{{ prefix "sdtl"       "https://rdf-vocabulary.ddialliance.org/sdth#" }}
{{ prefix "prov"       "http://www.w3.org/ns/prov#" }}
{{ prefix "provone"    "http://purl.dataone.org/provone/2015/01/15/ontology#" }}
{{ prefix "rdf"        "http://www.w3.org/1999/02/22-rdf-syntax-ns#" }}
{{ prefix "rdfs"       "http://www.w3.org/2000/01/rdf-schema#" }}
{{ prefix "sdth"       "https://rdf-vocabulary.ddialliance.org/sdth#" }}

{{ query "sdth_construct_provone_programs" '''
    CONSTRUCT {
        ?program rdf:type provone:Program .
    }
    WHERE {
        {
            ?program rdf:type sdth:Program .
        }
        UNION
        {
            ?program rdf:type sdth:ProgramStep .
        }
    } 
'''}}
