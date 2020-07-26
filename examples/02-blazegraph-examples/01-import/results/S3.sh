# IMPORT TWO TRIPLES AS JSON-LD

blazegraph drop
blazegraph import --format jsonld << END_DATA

    [
        {
            "@id": "http://tmcphill.net/data#x",
            "http://tmcphill.net/tags#tag": [ { "@value": "seven" } ]
        },
        {
            "@id": "http://tmcphill.net/data#y",
            "http://tmcphill.net/tags#tag": [ { "@value": "eight" } ]
        }
    ]

END_DATA

blazegraph export --format nt

