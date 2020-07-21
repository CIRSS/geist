package main

import (
	"strings"
	"testing"

	"github.com/tmcphillips/blazegraph-util/assert"
)

func TestBlazegraphCmd_query_json(t *testing.T) {

	var resultsBuffer strings.Builder
	Main.OutWriter = &resultsBuffer
	Main.ErrWriter = &resultsBuffer

	runWithArgs("blazegraph drop")
	runWithArgs("blazegraph load --file testdata/in.nt --format ttl")

	t.Run("json", func(t *testing.T) {
		resultsBuffer.Reset()
		runWithArgs("blazegraph query --file testdata/q1.sparql --format json")
		assert.JSONEquals(t, resultsBuffer.String(), `{
			"head": { "vars": ["s", "o"] },
			"results": { "bindings": [
				{
				"o": { "type": "literal", "value": "seven" },
				"s": { "type": "uri", "value": "http://tmcphill.net/data#x" }
				},
				{
				"o": { "type": "literal", "value": "eight" },
				"s": { "type": "uri", "value": "http://tmcphill.net/data#y" }
				}
			] } }
		`)
	})

	t.Run("xml", func(t *testing.T) {
		resultsBuffer.Reset()
		runWithArgs("blazegraph query --file testdata/q1.sparql --format xml")
		assert.LineContentsEqual(t, resultsBuffer.String(), `
			<?xml version='1.0' encoding='UTF-8'?>
            <sparql xmlns='http://www.w3.org/2005/sparql-results#'>
            	<head>
            		<variable name='s'/>
            		<variable name='o'/>
            	</head>
            	<results>
            		<result>
            			<binding name='s'>
            				<uri>http://tmcphill.net/data#x</uri>
            			</binding>
            			<binding name='o'>
            				<literal>seven</literal>
            			</binding>
            		</result>
            		<result>
            			<binding name='s'>
            				<uri>http://tmcphill.net/data#y</uri>
            			</binding>
            			<binding name='o'>
            				<literal>eight</literal>
            			</binding>
            		</result>
            	</results>
            </sparql>
		`)
	})

	t.Run("csv", func(t *testing.T) {
		resultsBuffer.Reset()
		runWithArgs("blazegraph query --file testdata/q1.sparql --format csv")
		results := resultsBuffer.String()
		assert.LineContentsEqual(t, results, `
			s,o
			http://tmcphill.net/data#x,seven
			http://tmcphill.net/data#y,eight`)
	})

}
