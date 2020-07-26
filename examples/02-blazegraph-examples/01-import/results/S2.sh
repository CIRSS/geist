# IMPORT TWO TRIPLES AS TURTLE

blazegraph drop
blazegraph import --format ttl << END_DATA

	@prefix data: <http://tmcphill.net/data#> .
	@prefix tags: <http://tmcphill.net/tags#> .

	data:y tags:tag "eight" .
	data:x tags:tag "seven" .

END_DATA

blazegraph export --format nt

