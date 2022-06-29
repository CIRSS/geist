#!/usr/bin/env bash

bash_cell SETUP << END_CELL

# INITIALIZE BLAZEGRAPH INSTANCE WITH CITATIONS

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --file ../data/citations.ttl --format ttl

END_CELL


bash_cell S1 << END_CELL

# EXPORT CITATIONS AS N-TRIPLES

geist export --format nt | sort

END_CELL


bash_cell S2 << END_CELL

# WHICH PAPERS DIRECTLY CITE WHICH PAPERS?

geist query --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>

    SELECT DISTINCT ?citing_paper ?cited_paper
    WHERE {
        ?citing_paper c:cites ?cited_paper .
    }
    ORDER BY ?citing_paper ?cited_paper

END_QUERY

END_CELL


bash_cell S3 << END_CELL

# WHICH PAPERS DEPEND ON WHICH PRIOR WORK?

geist query --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>

    SELECT DISTINCT ?paper ?prior_work
    WHERE {
        ?paper c:cites+ ?prior_work .
    }
    ORDER BY ?paper ?prior_work

END_QUERY

END_CELL


bash_cell S4 << END_CELL

# WHICH PAPERS DEPEND ON PAPER A?

geist query --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>
    prefix : <http://learningsparql.com/ns/papers#>

    SELECT DISTINCT ?paper
    WHERE {
        ?paper c:cites+ :paperA .
    }
    ORDER BY ?paper

END_QUERY

END_CELL


bash_cell S5 << END_CELL

# WHICH PAPERS CITE A PAPER THAT CITES PAPER A?

geist query --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>
    prefix : <http://learningsparql.com/ns/papers#>

    SELECT DISTINCT ?paper
    WHERE {
        ?paper c:cites/c:cites :paperA .
    }
    ORDER BY ?paper

END_QUERY

END_CELL


bash_cell S6 << END_CELL

# WHICH PAPERS CITE A PAPER CITED BY PAPER D?

geist query --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>
    prefix : <http://learningsparql.com/ns/papers#>

    SELECT DISTINCT ?paper
    WHERE {
        ?paper c:cites/^c:cites :paperD .
        FILTER(?paper != :paperD)
    }
    ORDER BY ?paper

END_QUERY

END_CELL


bash_cell S7 << END_CELL

# WHAT RESULTS DEPEND DIRECTLY ON RESULTS REPORTED BY PAPER A?

geist query --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>
    prefix : <http://learningsparql.com/ns/papers#>

    SELECT DISTINCT ?result
    WHERE {
        :paperA c:reports/^c:uses/c:reports ?result
    }
    ORDER BY ?result

END_QUERY

END_CELL


bash_cell S7 << END_CELL

# WHAT RESULTS DEPEND DIRECTLY OR INDIRECTLY ON RESULTS REPORTED BY PAPER A?

geist query --format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>
    prefix : <http://learningsparql.com/ns/papers#>

    SELECT DISTINCT ?result
    WHERE {
        :paperA c:reports/(^c:uses/c:reports)+ ?result
    }
    ORDER BY ?result

END_QUERY

END_CELL


bash_dot_cell S8 << '__END_CELL__'

# Visualization of Paper-Citation Graph

geist report << '__END_REPORT_TEMPLATE__'
    {{{
        {{ include "graphviz.g" }}
    }}}                                                             \\
                                                                    \\
    {{ prefix "dc" "http://purl.org/dc/elements/1.1/" }}            \\
    {{ prefix "c" "http://learningsparql.com/ns/citations#" }}      \\
                                                                    \\
    {{ gv_graph "wt_run" }}

    {{ gv_title "Paper-Citation Graph" }}

    {{ gv_cluster "citations" }}

    # paper nodes
    node[shape=box style="filled" fillcolor="#CCFFCC" peripheries=1 fontname=Courier]
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

__END_CELL__


bash_dot_cell S9 << '__END_CELL__'

# Visualization of Paper-Citation Graph

geist report << '__END_REPORT_TEMPLATE__'

    {{{
        {{ include "graphviz.g" }}
    }}}
                                                                \\
    {{ prefix "dc" "http://purl.org/dc/elements/1.1/" }}        \\
    {{ prefix "c" "http://learningsparql.com/ns/citations#" }}  \\
                                                                \\
    {{ gv_graph "wt_run" }}

    {{ gv_title "Result-Dependency Graph" }}

    {{ gv_cluster "citations" }}

    # result nodes
    node[shape=box style="rounded,filled" fillcolor="#FFFFCC" peripheries=1 fontname=Helvetica]
    {{ range $Result := select '''
        SELECT DISTINCT ?result ?label
        WHERE {
            ?paper rdf:type c:Paper .
            ?paper c:reports ?result .
            ?result rdfs:label ?label
        }
        ORDER BY ?result
        ''' | rows }}                                             \\
        {{ gv_labeled_node (index $Result 0) (index $Result 1) }}
    {{ end }}
                                                                    \\
    # result dependency edges
    {{ range $Dependency := select '''
            SELECT DISTINCT ?result1 ?result2
            WHERE {
                 ?result2 ^c:uses/c:reports ?result1
           }
            ORDER BY ?result1 ?result2
        ''' | rows }}                                               \\
        {{ gv_edge (index $Dependency 0) (index $Dependency 1) }}
    {{ end }}
                                                                    \\
    {{ gv_cluster_end }}

    {{ gv_end }}
                                                                    \\
__END_REPORT_TEMPLATE__

__END_CELL__


bash_dot_cell S10 << '__END_CELL__'

# Visualization of Paper-Result Graph

geist report << '__END_REPORT_TEMPLATE__'
                                                                    \\
    {{{                                                             \\
        {{ include "graphviz.g" }}                                  \\
    }}}                                                             \\
                                                                    \\
    {{ prefix "dc" "http://purl.org/dc/elements/1.1/" }}            \\
    {{ prefix "c" "http://learningsparql.com/ns/citations#" }}      \\
                                                                    \\
    {{ gv_graph "wt_run" }}

    {{ gv_title "Paper-Result Graph" }}

    {{ gv_cluster "citations" }}

    # paper nodes
    node[shape=box style="filled" fillcolor="#CCFFCC" peripheries=1 fontname=Courier]
    {{ range $Paper := select '''
        SELECT ?paper ?title
        WHERE {
            ?paper rdf:type c:Paper .
            ?paper dc:title ?title .
        }
        ORDER BY ?paper
        ''' | rows }}                                             \\
        {{ gv_labeled_node (index $Paper 0) (index $Paper 1) }}
    {{ end }}
                                                                    \\
    # result nodes
    node[shape=box style="rounded,filled" fillcolor="#FFFFCC" peripheries=1 fontname=Helvetica]
    {{ range $Result := select '''
        SELECT DISTINCT ?result ?label
        WHERE {
            ?paper rdf:type c:Paper .
            ?paper c:reports ?result .
            ?result rdfs:label ?label
        }
        ORDER BY ?result
        ''' | rows }}                                             \\
        {{ gv_labeled_node (index $Result 0) (index $Result 1) }}
    {{ end }}
                                                                    \\
    # reports edges
    {{ range $Report := select '''
            SELECT DISTINCT ?paper ?result
            WHERE {
                ?paper c:reports ?result .
            }
            ORDER BY ?paper ?result
        ''' | rows }}                                                \\
        {{ gv_edge (index $Report 0) (index $Report 1) }}
    {{ end }}

    # uses edges
    {{ range $Use := select '''
            SELECT DISTINCT ?result ?paper
            WHERE {
                ?paper c:uses ?result .
            }
            ORDER BY ?paper ?result
        ''' | rows }}                                                \\
        {{ gv_edge (index $Use 0) (index $Use 1) }}
    {{ end }}
                                                                    \\
    {{ gv_cluster_end }}

    {{ gv_end }}                                                    \\
                                                                    \\
__END_REPORT_TEMPLATE__

__END_CELL__
