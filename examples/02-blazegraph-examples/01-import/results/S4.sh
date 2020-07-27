# IMPORT TWO TRIPLES AS RDF-XML

blazegraph drop
blazegraph import --format xml << END_DATA

    <rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">

    <rdf:Description rdf:about="http://tmcphill.net/data#y">
        <tag xmlns="http://tmcphill.net/tags#">eight</tag>
    </rdf:Description>

    <rdf:Description rdf:about="http://tmcphill.net/data#x">
        <tag xmlns="http://tmcphill.net/tags#">seven</tag>
    </rdf:Description>

    </rdf:RDF>

END_DATA

blazegraph export --format nt | sort

