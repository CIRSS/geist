package main

import (
	"strings"
	"testing"

	"github.com/tmcphillips/blazegraph-util/util"
)

func TestBlazegraphCmd_report(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph drop")

	Main.InReader = strings.NewReader(`
		<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
	`)
	run("blazegraph import --format ttl")

	t.Run("constant-template", func(t *testing.T) {
		outputBuffer.Reset()
		template := `A constant template.`
		Main.InReader = strings.NewReader(template)
		run("blazegraph report")
		util.LineContentsEqual(t, outputBuffer.String(), `
			A constant template.
		`)
	})

	t.Run("table-with-separators", func(t *testing.T) {
		outputBuffer.Reset()
		template := `{{select <% A constant template %>}}`
		Main.InReader = strings.NewReader(template)
		run("blazegraph report")
		util.LineContentsEqual(t, outputBuffer.String(), `
			A CONSTANT TEMPLATE
		`)
	})

}
