# IMPORT TWO TRIPLES AS N-TRIPLES

blazegraph import --format nt << END_DATA

	<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .

END_DATA

blazegraph export --format nt

