#!/usr/bin/env bash

# *****************************************************************************

bash_cell SETUP << END_CELL

# INITIALIZE BLAZEGRAPH INSTANCE WITH ADDRESS BOOK
geist destroy --dataset kb --quiet
geist create --dataset kb --quiet
geist import --file ../data/address-book.jsonld --format jsonld

END_CELL

# *****************************************************************************

bash_cell S1 << END_CELL

# EXPORT ADDRESS BOOK AS JSON-LD

geist export --format jsonld

END_CELL

# *****************************************************************************

bash_cell S1 << END_CELL

# EXPORT ADDRESS BOOK AS TURTLE

geist export --format ttl

END_CELL

# *****************************************************************************

bash_cell S1 << END_CELL

# EXPORT ADDRESS BOOK AS N-TRIPLES

geist export --format nt | sort

END_CELL

# *****************************************************************************

bash_cell S1 << END_CELL

# EXPORT ADDRESS BOOK AS RDF-XML

geist export --format xml

END_CELL
