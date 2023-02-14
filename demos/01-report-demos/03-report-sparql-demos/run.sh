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

bash_cell select_direct_citations_from_report << END_CELL

geist report << END_TEMPLATE

    {{ select '''

        prefix c: <http://learningsparql.com/ns/citations#>

        SELECT DISTINCT ?citing_paper ?cited_paper
        WHERE {
            ?citing_paper c:cites ?cited_paper .
        }
        ORDER BY ?citing_paper ?cited_paper

    ''' | tabulate }}

END_TEMPLATE

END_CELL

