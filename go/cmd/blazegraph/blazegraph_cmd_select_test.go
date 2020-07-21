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

	t.Run("default", func(t *testing.T) {
		resultsBuffer.Reset()
		runWithArgs("blazegraph query --file testdata/q1.sparql")
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

	t.Run("csv", func(t *testing.T) {
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
