#!/usr/bin/env bash

# *****************************************************************************

bash_cell load_citation_data << END_CELL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --file citations.ttl --format ttl

END_CELL

# *****************************************************************************

bash_cell export_citation_data << END_CELL

geist export --format nt | sort

END_CELL

# *****************************************************************************

bash_cell query_direct_citations << END_CELL

geist query -format table << END_QUERY

    prefix c: <http://learningsparql.com/ns/citations#>

    SELECT DISTINCT ?citing_paper ?cited_paper
    WHERE {
        ?citing_paper c:cites ?cited_paper .
    }
    ORDER BY ?citing_paper ?cited_paper

END_QUERY

END_CELL

# *****************************************************************************

bash_cell tabulate_direct_citations_from_report << END_CELL

geist report << END_TEMPLATE

    {{ prefix "c" "http://learningsparql.com/ns/citations#" }}

    {{ select '''

        SELECT DISTINCT ?citing_paper ?cited_paper
        WHERE {
            ?citing_paper c:cites ?cited_paper .
        }
        ORDER BY ?citing_paper ?cited_paper

    ''' | tabulate }}

END_TEMPLATE

END_CELL


# *****************************************************************************

bash_cell format_direct_citations_from_report << 'END_CELL'

geist report << 'END_TEMPLATE'

    {{ prefix "c" "http://learningsparql.com/ns/citations#" }}              \\
                                                                            \\
    {{ range $Citation := select '''

            SELECT DISTINCT ?citing_paper ?cited_paper
            WHERE {
                ?citing_paper c:cites ?cited_paper .
            }
            ORDER BY ?citing_paper ?cited_paper

        ''' | rows }}                                                       \\
                                                                            \\
        {{ index $Citation 0 }} ---cites--> {{ index $Citation 1 }}
                                                                            \\
    {{ end }}

END_TEMPLATE

END_CELL


# *****************************************************************************

bash_cell format_direct_citations_using_query << 'END_CELL'

geist report << 'END_TEMPLATE'

    {{{
        {{ query "select_direct_citations" '''
            SELECT DISTINCT ?citing_paper ?cited_paper
            WHERE {
                ?citing_paper c:cites ?cited_paper .
            }
            ORDER BY ?citing_paper ?cited_paper
        '''}}
    }}}

    {{ prefix "c" "http://learningsparql.com/ns/citations#" }}              \\
                                                                            \\
    {{ range $Citation := select_direct_citations | rows }}                 \\
        {{ index $Citation 0 }} ---cites--> {{ index $Citation 1 }}
    {{ end }}

END_TEMPLATE

END_CELL

# *****************************************************************************

bash_dot_cell visualize_direct_citations_using_graphviz << 'END_CELL'

geist report << 'END_TEMPLATE'

    {{{
        {{ include "graphviz.g" }}

        {{ query "select_papers" '''
            SELECT ?paper ?title
            WHERE {
                ?paper rdf:type c:Paper .
                ?paper dc:title ?title .
            }
            ORDER BY ?paper ?title
        ''' }}

        {{ query "select_direct_citations" '''
            SELECT DISTINCT ?citing_paper ?cited_paper
            WHERE {
                ?citing_paper c:cites ?cited_paper .
            }
            ORDER BY ?citing_paper ?cited_paper
        '''}}
    }}}                                                             \\
                                                                    \\
    {{ prefix "dc" "http://purl.org/dc/elements/1.1/" }}            \\
    {{ prefix "c" "http://learningsparql.com/ns/citations#" }}      \\
                                                                    \\
    {{ gv_graph "citation_graph" "BT" }}
    {{ gv_title "Paper-Citation Graph" }}
    {{ gv_cluster "citations" }}

    node[shape=box style="filled" fillcolor="#CCFFCC" peripheries=1 fontname=Courier]
    {{ range $Paper := select_papers | rows }}                      \\
        {{ gv_labeled_node (index $Paper 0) (index $Paper 1) }}
    {{ end }}
                                                                    \\
    {{ range $Citation := select_direct_citations | rows }}         \\
        {{ gv_edge (index $Citation 0) (index $Citation 1) }}
    {{ end }}
                                                                    \\
    {{ gv_cluster_end }}
    {{ gv_end }}

END_TEMPLATE

END_CELL
