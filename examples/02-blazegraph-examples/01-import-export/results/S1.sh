# Export address book as JSON-LD

blazegraph drop
blazegraph import --file ../data/address-book.jsonld --format jsonld
blazegraph export --format jsonld
