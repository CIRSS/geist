package tests

import (
	"testing"

	"github.com/cirss/geist/reporter"
	"github.com/cirss/geist/util"
)

func TestReportTemplate_StaticText_OneLinef(t *testing.T) {

	rt := reporter.NewReportTemplate(
		"main",
		"42 items are made of cotton", nil)

	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual,
		"42 items are made of cotton")
}

func TestReportTemplate_StaticReport_MultipleLines(t *testing.T) {

	rt := reporter.NewReportTemplate(
		"main",
		`
		42 items
		are made of
		cotton
		`, nil)

	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual,
		`
		42 items
		are made of
		cotton
	`)
}

func TestReportTemplate_StaticReport_MultipleLines_EscapeOneLineEnding(t *testing.T) {

	rt := reporter.NewReportTemplate(
		"main",
		`
		42 items{{sp}}    \
		are made of
		cotton
		`, nil)

	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual,
		`
		42 items are made of
		cotton
	`)
}

func TestReportTemplate_StaticReport_MultipleLines_EscapeTwoLineEnding(t *testing.T) {

	rt := reporter.NewReportTemplate(
		"main",
		`
		42 items{{sp}}    \
		are made of{{sp}} \
		cotton
		`, nil)

	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual,
		`
		42 items are made of cotton
	`)
}

func TestReportTemplate_StaticReport_PercentCharacter(t *testing.T) {

	rt := reporter.NewReportTemplate(
		"main",
		`
		42% of items
		are made of
		cotton
		`, nil)

	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual,
		`
		42% of items
		are made of
		cotton
	`)
}

func TestReportTemplate_StaticReport_Printf(t *testing.T) {

	rt := reporter.NewReportTemplate(
		"main",
		`
		{{printf "%d" 42}}{{println "% of items"}} \
		are made of
		cotton
		`, nil)

	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual,
		`
		42% of items
		are made of
		cotton
	`)
}
