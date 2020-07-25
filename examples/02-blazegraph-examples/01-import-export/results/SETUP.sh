# INITIALIZE BLAZEGRAPH INSTANCE WITH ADDRESS BOOK

blazegraph drop
blazegraph import --file ../data/address-book.jsonld --format jsonld
