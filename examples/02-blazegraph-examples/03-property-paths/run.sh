#!/usr/bin/env bash

DOT_RUNNER='../../common/run_dot_examples.sh'
SCRIPT_RUNNER='../../common/run_script_example.sh'


bash ${SCRIPT_RUNNER} SETUP "INITIALIZE BLAZEGRAPH INSTANCE WITH CITATIONS" << END_SCRIPT

blazegraph drop
blazegraph import --file ../data/citations.ttl --format ttl

END_SCRIPT


bash ${SCRIPT_RUNNER} S1 "EXPORT CITATIONS AS N-TRIPLES" << END_SCRIPT

blazegraph export --format nt | sort

END_SCRIPT


bash ${SCRIPT_RUNNER} S2 "WHICH PAPERS DIRECTLY CITE WHICH PAPERS?" \
    << END_SCRIPT

blazegraph select --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>

    SELECT DISTINCT ?citing_paper ?cited_paper
    WHERE {
        ?citing_paper c:cites ?cited_paper .
    }
    ORDER BY ?citing_paper ?cited_paper

END_QUERY

END_SCRIPT


bash ${SCRIPT_RUNNER} S3 "WHICH PAPERS DEPEND ON WHICH PRIOR WORK?" \
    << END_SCRIPT

blazegraph select --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>

    SELECT DISTINCT ?paper ?prior_work
    WHERE {
        ?paper c:cites+ ?prior_work .
    }
    ORDER BY ?paper ?prior_work

END_QUERY

END_SCRIPT


bash ${SCRIPT_RUNNER} S4 "WHICH PAPERS DEPEND ON PAPER A?" \
    << END_SCRIPT

blazegraph select --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>
    prefix : <http://learningsparql.com/ns/papers#>

    SELECT DISTINCT ?paper
    WHERE {
        ?paper c:cites+ :paperA .
    }
    ORDER BY ?paper

END_QUERY

END_SCRIPT


bash ${SCRIPT_RUNNER} S5 "WHICH PAPERS CITE A PAPER THAT CITES PAPER A?" \
    << END_SCRIPT

blazegraph select --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>
    prefix : <http://learningsparql.com/ns/papers#>

    SELECT DISTINCT ?paper
    WHERE {
        ?paper c:cites/c:cites :paperA .
    }
    ORDER BY ?paper

END_QUERY

END_SCRIPT


bash ${SCRIPT_RUNNER} S6 "WHICH PAPERS CITE A PAPER CITED BY PAPER D?" \
    << END_SCRIPT

blazegraph select --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>
    prefix : <http://learningsparql.com/ns/papers#>

    SELECT DISTINCT ?paper
    WHERE {
        ?paper c:cites/^c:cites :paperD .
        FILTER(?paper != :paperD)
    }
    ORDER BY ?paper

END_QUERY

END_SCRIPT


bash ${SCRIPT_RUNNER} S7 "WHAT RESULTS DEPEND DIRECTLY ON RESULTS REPORTED BY PAPER A?" \
    << END_SCRIPT

blazegraph select --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>
    prefix : <http://learningsparql.com/ns/papers#>

    SELECT DISTINCT ?result
    WHERE {
        :paperA c:reports/^c:uses/c:reports ?result
    }
    ORDER BY ?result

END_QUERY

END_SCRIPT


bash ${SCRIPT_RUNNER} S7 "WHAT RESULTS DEPEND DIRECTLY OR INDIRECTLY ON RESULTS REPORTED BY PAPER A?" \
    << END_SCRIPT

blazegraph select --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>
    prefix : <http://learningsparql.com/ns/papers#>

    SELECT DISTINCT ?result
    WHERE {
        :paperA c:reports/(^c:uses/c:reports)+ ?result
    }
    ORDER BY ?result

END_QUERY

END_SCRIPT


bash ${DOT_RUNNER} S8 "Visualization of Citation Graph" \
    << '__END_SCRIPT__'

blazegraph report << '__END_REPORT_TEMPLATE__'

    {{{
        {{ include "graphviz.g" }}
    }}}

    {{ prefix "dc" "http://purl.org/dc/elements/1.1/" }}
    {{ prefix "c" "http://learningsparql.com/ns/citations#" }}

    {{ gv_graph "wt_run" }}

    {{ gv_title "Citation Graph" }}
    
    {{ gv_cluster "citations" }}

    # paper nodes
    {{ range $Paper := select '''
        SELECT ?paper ?title
        WHERE {
            ?paper rdf:type c:Paper .
            ?paper dc:title ?title .
        } ''' | rows }}                                             \\
        {{ gv_labeled_node (index $Paper 0) (index $Paper 1) }}
    {{ end }}
                                                                    \\
    # citation edges
    {{ range $Citation := select '''
            SELECT DISTINCT ?citing_paper ?cited_paper
            WHERE {
                ?citing_paper c:cites ?cited_paper .
            }
            ORDER BY ?citing_paper ?cited_paper        
        ''' | rows }}                                                \\
        {{ gv_edge (index $Citation 0) (index $Citation 1) }}
    {{ end }}
                                                                    \\
    {{ gv_cluster_end }}
    
    {{ gv_end }}

__END_REPORT_TEMPLATE__

__END_SCRIPT__
