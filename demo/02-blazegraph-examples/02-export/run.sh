#!/usr/bin/env bash

# *****************************************************************************

bash_cell SETUP "INITIALIZE BLAZEGRAPH INSTANCE WITH ADDRESS BOOK" << END_CELL

geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --file ../data/address-book.jsonld --format jsonld

END_CELL

# *****************************************************************************

bash_cell S1 "EXPORT ADDRESS BOOK AS JSON-LD" << END_CELL

geist export --format jsonld

END_CELL

# *****************************************************************************

bash_cell S1 "EXPORT ADDRESS BOOK AS TURTLE" << END_CELL

geist export --format ttl

END_CELL

# *****************************************************************************

bash_cell S1 "EXPORT ADDRESS BOOK AS N-TRIPLES" << END_CELL

geist export --format nt | sort

END_CELL

# *****************************************************************************

bash_cell S1 "EXPORT ADDRESS BOOK AS RDF-XML" << END_CELL

geist export --format xml

END_CELL
