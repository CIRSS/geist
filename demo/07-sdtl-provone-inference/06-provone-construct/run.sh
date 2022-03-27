#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'
GRAPHER='../../common/run_dot_examples.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT SDTL AS JSON-LD" << END_SCRIPT

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet 
geist import --format jsonld --file ../data/compute-sdth.jsonld

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} E1 "EXPORT ORIGINAL SDTL AS N-TRIPLES" << END_SCRIPT

geist export --format nt | sort

END_SCRIPT


bash ${RUNNER} Q1 "CONSTRUCT PROVONE PROGRAMS VIA SPARQL CONSTRUCT QUERY" << END_SCRIPT

geist query --format table << __END_QUERY__

    PREFIX prov: <http://www.w3.org/ns/prov#>
    PREFIX provone: <http://purl.dataone.org/provone/2015/01/15/ontology#>
    PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
    PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
    PREFIX sdth: <https://rdf-vocabulary.ddialliance.org/sdth#>

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


__END_QUERY__

END_SCRIPT


bash ${RUNNER} R1 "CONSTRUCT PROVONE PROGRAMS VIA GEIST REPORT" << '__END_SCRIPT__'

geist report << '__END_REPORT_TEMPLATE__'

    {{{
        {{ include "../../common/sdth.g" }}
    }}}
                                                                                        \\
    {{ range $Program := (sdth_construct_provone_program_triples | rows) }}             \\
        {{ ntriple_print $Program }}
    {{ end }}                                                                           \
                                                                                        \\
    {{ range $Program := select_sdth_program | vector }}                                \\
        {{ range $ProgramStep := (select_sdth_program_steps $Program | rows) }}         \\
            <{{$Program}}> provone:hasSubProgram <{{ index $ProgramStep 0 }}> .
        {{ end }}
    {{ end }}
    
    {{ range $DataframeProducer := (select_dataframe_producers | rows) }}
        {{ with $StepName := (index $DataframeProducer 0) }}
            <{{ $StepName }}> provone:hasOutputPort <{{ $StepName }}/dataframeport/{{ index $DataframeProducer 2 }}_out> .
        {{ end }}
    {{ end }}

    {{ range $DataframeConsumer := (select_dataframe_consumers | rows) }}
        {{ with $StepName := (index $DataframeConsumer 0) }}
            <{{ $StepName }}> provone:hasInputPort <{{ $StepName }}/dataframeport/{{ index $DataframeConsumer 2 }}_in> .
        {{ end }}
    {{ end }}

__END_REPORT_TEMPLATE__

__END_SCRIPT__
