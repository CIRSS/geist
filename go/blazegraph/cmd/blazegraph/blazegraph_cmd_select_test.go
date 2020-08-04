package main

import (
	"strings"
	"testing"

	"github.com/tmcphillips/blazegraph-util/util"
)

func TestBlazegraphCmd_query_json(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph drop")

	Main.InReader = strings.NewReader(`
		<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
	`)
	run("blazegraph import --format ttl")

	query := `
		prefix ab: <http://tmcphill.net/tags#>
		SELECT ?s ?o
		WHERE
		{ ?s ab:tag ?o }
	`

	t.Run("json", func(t *testing.T) {
		outputBuffer.Reset()
		Main.InReader = strings.NewReader(query)
		run("blazegraph select --format json")
		util.JSONEquals(t, outputBuffer.String(), `{
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

	t.Run("table-with-separators", func(t *testing.T) {
		outputBuffer.Reset()
		Main.InReader = strings.NewReader(query)
		run("blazegraph select --format table")
		util.LineContentsEqual(t, outputBuffer.String(), `
			s                          | o
            ----------------------------------
            http://tmcphill.net/data#x | seven
            http://tmcphill.net/data#y | eight

		`)
	})

	t.Run("table-without-separators", func(t *testing.T) {
		outputBuffer.Reset()
		Main.InReader = strings.NewReader(query)
		run("blazegraph select --format table --columnseparators=false")
		util.LineContentsEqual(t, outputBuffer.String(), `
			s                            o
            ----------------------------------
            http://tmcphill.net/data#x   seven
            http://tmcphill.net/data#y   eight

		`)
	})

	t.Run("xml", func(t *testing.T) {
		outputBuffer.Reset()
		Main.InReader = strings.NewReader(query)
		run("blazegraph select --format xml")
		util.LineContentsEqual(t, outputBuffer.String(), `
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
		outputBuffer.Reset()
		Main.InReader = strings.NewReader(query)
		run("blazegraph select --format csv")
		util.LineContentsEqual(t, outputBuffer.String(), `
			s,o
			http://tmcphill.net/data#x,seven
			http://tmcphill.net/data#y,eight`)
	})

}
