package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/util"
)

func TestBlazegraphCmd_query_json(t *testing.T) {

	var outputBuffer strings.Builder
	Program.OutWriter = &outputBuffer
	Program.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --dataset kb")

	Program.InReader = strings.NewReader(`
		<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
	`)

	run("blazegraph import --format ttl")

	query := `
		prefix ab: <http://tmcphill.net/tags#>
		SELECT ?s ?o
		WHERE
		{ ?s ab:tag ?o }
		ORDER BY ?s
		`

	t.Run("json", func(t *testing.T) {
		outputBuffer.Reset()
		Program.InReader = strings.NewReader(query)
		assertExitCode(t, "blazegraph query --format json", 0)
		util.JSONEquals(t, outputBuffer.String(),
			`{
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
		Program.InReader = strings.NewReader(query)
		assertExitCode(t, "blazegraph query --format table", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`s                          | o
 			 ==================================
             http://tmcphill.net/data#x | seven
             http://tmcphill.net/data#y | eight
			`)
	})

	t.Run("table-without-separators", func(t *testing.T) {
		outputBuffer.Reset()
		Program.InReader = strings.NewReader(query)
		assertExitCode(t, "blazegraph query --format table --columnseparators=false", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`s                            o
			 ==================================
             http://tmcphill.net/data#x   seven
             http://tmcphill.net/data#y   eight
		`)
	})

	t.Run("xml", func(t *testing.T) {
		outputBuffer.Reset()
		Program.InReader = strings.NewReader(query)
		assertExitCode(t, "blazegraph query --format xml", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`<?xml version='1.0' encoding='UTF-8'?>
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
		Program.InReader = strings.NewReader(query)
		assertExitCode(t, "blazegraph query --format csv", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`s,o
			 http://tmcphill.net/data#x,seven
			 http://tmcphill.net/data#y,eight
			`)
	})
}

func TestBlazegraphCmd_query_help(t *testing.T) {

	var outputBuffer strings.Builder
	Program.OutWriter = &outputBuffer
	Program.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph query help", 0)
	util.LineContentsEqual(t, outputBuffer.String(),
		`
		Performs a SPARQL query on the identified RDF dataset.

		Usage: blazegraph query [<flags>]

		Flags:

		-columnseparators
				Display column separators in table format (default true)
		-dataset name
    	    	name of RDF dataset to query (default "kb")
		-dryrun
				Output query but do not execute it
		-file string
				File containing the SPARQL query to execute (default "-")
		-format string
			Format of result set to produce [csv, json, table, or xml] (default "json")
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output

	`)
}

func TestBlazegraphCmd_help_select(t *testing.T) {

	var outputBuffer strings.Builder
	Program.OutWriter = &outputBuffer
	Program.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph help query", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`
		Performs a SPARQL query on the identified RDF dataset.

		Usage: blazegraph query [<flags>]

		Flags:

		-columnseparators
				Display column separators in table format (default true)
		-dataset name
    	    	name of RDF dataset to query (default "kb")
		-dryrun
				Output query but do not execute it
		-file string
				File containing the SPARQL query to execute (default "-")
		-format string
			Format of result set to produce [csv, json, table, or xml] (default "json")
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output

	`)
}

func TestBlazegraphCmd_query_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Program.OutWriter = &outputBuffer
	Program.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph query --not-a-flag", 1)

	util.LineContentsEqual(t, outputBuffer.String(),
		`
		flag provided but not defined: -not-a-flag

		Usage: blazegraph query [<flags>]

		Flags:

		-columnseparators
				Display column separators in table format (default true)
		-dataset name
    	    	name of RDF dataset to query (default "kb")
		-dryrun
				Output query but do not execute it
		-file string
				File containing the SPARQL query to execute (default "-")
		-format string
			Format of result set to produce [csv, json, table, or xml] (default "json")
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output

	`)
}
