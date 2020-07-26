#!/usr/bin/env bash

RUNNER='../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "INITIALIZE BLAZEGRAPH INSTANCE" << END_SCRIPT

blazegraph drop

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} S1 "IMPORT TWO TRIPLES AS N-TRIPLES" << END_SCRIPT

blazegraph import --format ttl << END_DATA

	<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .

END_DATA

blazegraph export --format nt

END_SCRIPT

# bash ${RUNNER} S1 "EXPORT ADDRESS BOOK AS TURTLE" << END_SCRIPT
# blazegraph export --format ttl
# END_SCRIPT

# bash ${RUNNER} S1 "EXPORT ADDRESS BOOK AS N-TRIPLES" << END_SCRIPT
# blazegraph export --format nt
# END_SCRIPT

# bash ${RUNNER} S1 "EXPORT ADDRESS BOOK AS RDF-XML" << END_SCRIPT
# blazegraph export --format xml
# END_SCRIPT
