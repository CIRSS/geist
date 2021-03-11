package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/go/util"
)

func TestReportInfer_rdfs_subClassOf(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	load := func() {
		Main.InReader = strings.NewReader(`
			@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
			@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
			@prefix tm: <http://tmcphill.net/ns/data#> .

			tm:typeb rdfs:subClassOf tm:typea .
			tm:typec rdfs:subClassOf tm:typeb .

			tm:a1 rdf:type tm:typea .
			tm:a2 rdf:type tm:typea .
			tm:b1 rdf:type tm:typeb .
			tm:b2 rdf:type tm:typeb .
			tm:b3 rdf:type tm:typeb .
			tm:c1 rdf:type tm:typec .
		`)
		run("blazegraph import --format ttl")
	}

	report := func() {
		q :=
			`{{ prefix "rdf" "http://www.w3.org/1999/02/22-rdf-syntax-ns#" }}	\
 			{{ prefix "tm" "http://tmcphill.net/ns/data#" }}					\
																				\
			{{ select '''
				SELECT ?type
				WHERE
				{ ?type rdf:type rdfs:Class }
				ORDER BY ?type ''' | tabulate }}
																				\
			{{ select '''
				SELECT ?a
				WHERE
				{ ?a rdf:type tm:typea }
				ORDER BY ?a ''' | tabulate }}
																				\
			{{ select '''
				SELECT ?b
				WHERE
				{ ?b rdf:type tm:typeb }
				ORDER BY ?b ''' | tabulate }}
																				\
			{{ select '''
				SELECT ?c
				WHERE
				{ ?c rdf:type tm:typec }
				ORDER BY ?c ''' | tabulate }}
		`
		Main.InReader = strings.NewReader(q)
		run("blazegraph report")
	}

	t.Run("infer-none", func(t *testing.T) {
		outputBuffer.Reset()
		run("blazegraph destroy --dataset kb")
		run("blazegraph create --quiet --dataset kb --infer none")
		load()
		report()
		util.LineContentsEqual(t, outputBuffer.String(),
			`type
			 ===

			a
			=============================
            http://tmcphill.net/ns/data#a1
			http://tmcphill.net/ns/data#a2

            b
            =============================
            http://tmcphill.net/ns/data#b1
            http://tmcphill.net/ns/data#b2
            http://tmcphill.net/ns/data#b3

            c
            =============================
            http://tmcphill.net/ns/data#c1

			`)
	})

	t.Run("infer-rdfs", func(t *testing.T) {
		outputBuffer.Reset()
		run("blazegraph destroy --dataset kb")
		run("blazegraph create --quiet --dataset kb --infer rdfs")
		load()
		report()
		util.LineContentsEqual(t, outputBuffer.String(),
			`type
            ===============================================================
            http://tmcphill.net/ns/data#typea
            http://tmcphill.net/ns/data#typeb
            http://tmcphill.net/ns/data#typec
            http://www.w3.org/1999/02/22-rdf-syntax-ns#Alt
            http://www.w3.org/1999/02/22-rdf-syntax-ns#Bag
            http://www.w3.org/1999/02/22-rdf-syntax-ns#List
            http://www.w3.org/1999/02/22-rdf-syntax-ns#Property
            http://www.w3.org/1999/02/22-rdf-syntax-ns#Seq
            http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement
            http://www.w3.org/1999/02/22-rdf-syntax-ns#XMLLiteral
            http://www.w3.org/2000/01/rdf-schema#Class
            http://www.w3.org/2000/01/rdf-schema#Container
            http://www.w3.org/2000/01/rdf-schema#ContainerMembershipProperty
            http://www.w3.org/2000/01/rdf-schema#Datatype
            http://www.w3.org/2000/01/rdf-schema#Literal
            http://www.w3.org/2000/01/rdf-schema#Resource

			a
			=============================
			http://tmcphill.net/ns/data#a1
			http://tmcphill.net/ns/data#a2
			http://tmcphill.net/ns/data#b1
			http://tmcphill.net/ns/data#b2
			http://tmcphill.net/ns/data#b3
			http://tmcphill.net/ns/data#c1

			b
			=============================
			http://tmcphill.net/ns/data#b1
			http://tmcphill.net/ns/data#b2
			http://tmcphill.net/ns/data#b3
			http://tmcphill.net/ns/data#c1

			c
			=============================
			http://tmcphill.net/ns/data#c1

		`)
	})

	t.Run("infer-owl", func(t *testing.T) {
		outputBuffer.Reset()
		run("blazegraph destroy --dataset kb")
		run("blazegraph create --quiet --dataset kb --infer owl")
		load()
		report()
		util.LineContentsEqual(t, outputBuffer.String(),
			`type
			===============================================================
			http://tmcphill.net/ns/data#typea
			http://tmcphill.net/ns/data#typeb
			http://tmcphill.net/ns/data#typec
			http://www.w3.org/1999/02/22-rdf-syntax-ns#Alt
			http://www.w3.org/1999/02/22-rdf-syntax-ns#Bag
			http://www.w3.org/1999/02/22-rdf-syntax-ns#List
			http://www.w3.org/1999/02/22-rdf-syntax-ns#Property
			http://www.w3.org/1999/02/22-rdf-syntax-ns#Seq
			http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement
			http://www.w3.org/1999/02/22-rdf-syntax-ns#XMLLiteral
			http://www.w3.org/2000/01/rdf-schema#Class
			http://www.w3.org/2000/01/rdf-schema#Container
			http://www.w3.org/2000/01/rdf-schema#ContainerMembershipProperty
			http://www.w3.org/2000/01/rdf-schema#Datatype
			http://www.w3.org/2000/01/rdf-schema#Literal
			http://www.w3.org/2000/01/rdf-schema#Resource
			http://www.w3.org/2002/07/owl#Class
			http://www.w3.org/2002/07/owl#DatatypeProperty
			http://www.w3.org/2002/07/owl#ObjectProperty
			http://www.w3.org/2002/07/owl#Restriction
			http://www.w3.org/2002/07/owl#TransitiveProperty

			a
			=============================
			http://tmcphill.net/ns/data#a1
			http://tmcphill.net/ns/data#a2
			http://tmcphill.net/ns/data#b1
			http://tmcphill.net/ns/data#b2
			http://tmcphill.net/ns/data#b3
			http://tmcphill.net/ns/data#c1

			b
			=============================
			http://tmcphill.net/ns/data#b1
			http://tmcphill.net/ns/data#b2
			http://tmcphill.net/ns/data#b3
			http://tmcphill.net/ns/data#c1

			c
			=============================
			http://tmcphill.net/ns/data#c1

		`)
	})

}

func TestReportInfer_rdf_inverseOf(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	load := func() {
		Main.InReader = strings.NewReader(`
			@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
			@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
			@prefix owl: <http://www.w3.org/2002/07/owl#> .
			@prefix person: <http://tmcphill.net/ns/person#> .
			@prefix verb: <http://tmcphill.net/ns/verb#> .
			@prefix tool: <http://tmcphill.net/ns/tool#> .

			verb:usedBy owl:inverseOf verb:uses .

			person:tim verb:uses tool:mouse .
			tool:keyboard verb:usedBy person:tim .
		`)
		run("blazegraph import --format ttl")
	}

	report := func() {
		q :=
			`{{ prefix "rdf" "http://www.w3.org/1999/02/22-rdf-syntax-ns#" }} 	\
			{{ prefix "verb" "http://tmcphill.net/ns/verb#" }}					\
																				\
			*** Object of verb:uses ***

			{{ select '''
				SELECT ?usedTool
				WHERE
				{ ?person verb:uses ?usedTool }
				ORDER BY ?usedTool ''' | tabulate }}

			*** Subject of verb:usedBy ***

			{{ select '''
				SELECT ?usedTool
				WHERE
				{ ?usedTool verb:usedBy ?person }
				ORDER BY ?usedTool ''' | tabulate }}
		`
		Main.InReader = strings.NewReader(q)
		run("blazegraph report")
	}

	t.Run("infer-none", func(t *testing.T) {
		outputBuffer.Reset()
		run("blazegraph destroy --dataset kb")
		run("blazegraph create --quiet --dataset kb --infer none")
		load()
		report()
		util.LineContentsEqual(t, outputBuffer.String(),
			`*** Object of verb:uses ***

			usedTool
			================================
			http://tmcphill.net/ns/tool#mouse


			*** Subject of verb:usedBy ***

			usedTool
			===================================
			http://tmcphill.net/ns/tool#keyboard

			`)
	})

	t.Run("infer-rdfs", func(t *testing.T) {
		outputBuffer.Reset()
		run("blazegraph destroy --dataset kb")
		run("blazegraph create --quiet --dataset kb --infer rdfs")
		load()
		report()
		util.LineContentsEqual(t, outputBuffer.String(),
			`*** Object of verb:uses ***

			usedTool
			================================
			http://tmcphill.net/ns/tool#mouse


			*** Subject of verb:usedBy ***

			usedTool
			===================================
			http://tmcphill.net/ns/tool#keyboard

			`)
	})

	t.Run("infer-owl", func(t *testing.T) {
		outputBuffer.Reset()
		run("blazegraph destroy --dataset kb")
		run("blazegraph create --quiet --dataset kb --infer owl")
		load()
		report()
		util.LineContentsEqual(t, outputBuffer.String(),
			`*** Object of verb:uses ***

			usedTool
			===================================
			http://tmcphill.net/ns/tool#keyboard
			http://tmcphill.net/ns/tool#mouse


			*** Subject of verb:usedBy ***

			usedTool
			===================================
			http://tmcphill.net/ns/tool#keyboard
			http://tmcphill.net/ns/tool#mouse

			`)
	})
}

func TestReportInfer_rdfs_domain_range(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	load := func() {
		Main.InReader = strings.NewReader(`
			@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
			@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
			@prefix owl: <http://www.w3.org/2002/07/owl#> .
			@prefix tm: <http://tmcphill.net/ns#> .
			@prefix person: <http://tmcphill.net/ns/person#> .
			@prefix verb: <http://tmcphill.net/ns/verb#> .
			@prefix tool: <http://tmcphill.net/ns/tool#> .

			verb:uses rdfs:domain tm:Person .
			verb:uses rdfs:range tm:Tool .
			verb:usedBy owl:inverseOf verb:uses .

			person:tim rdf:type tm:Person .
			tool:keyboard rdf:type tm:Tool .

			person:tim verb:uses tool:mouse .
			tool:keyboard verb:usedBy person:tim .

			tool:mouse verb:uses person:tim .
		`)
		run("blazegraph import --format ttl")
	}

	report := func() {
		q :=
			`{{ prefix "rdf"     "http://www.w3.org/1999/02/22-rdf-syntax-ns#" }}	\
			{{ prefix "tm"      "http://tmcphill.net/ns/data#" }}					\
			{{ prefix "verb"    "http://tmcphill.net/ns/verb#" }}					\
			{{ prefix "tool"    "http://tmcphill.net/ns/tool#" }}					\
			{{ prefix "person"  "http://tmcphill.net/ns/person#" }}					\
																					\
			*** Statements about verb:uses ***
			{{ select '''
				SELECT ?p ?o
				WHERE
				{ verb:uses ?p ?o }
				ORDER BY ?p ?o ''' | tabulate }}

			*** Statements about tool:keyboard ***
			{{ select '''
				SELECT ?p ?o
				WHERE
				{ tool:keyboard ?p ?o }
				ORDER BY ?p ?o ''' | tabulate }}

			*** Statements about tool:mouse ***
			{{ select '''
				SELECT ?p ?o
				WHERE
				{ tool:mouse ?p ?o }
				ORDER BY ?p ?o ''' | tabulate }}

			*** Statements about person:tim ***
			{{ select '''
				SELECT ?p ?o
				WHERE
				{ person:tim ?p ?o }
				ORDER BY ?p ?o ''' | tabulate }}

			*** What verb:uses what ***
			{{ select '''
				SELECT ?s ?o
				WHERE
				{ ?s verb:uses ?o }
				ORDER BY ?s ?o ''' | tabulate }}

			*** What is verb:usedBy what ***
			{{ select '''
				SELECT ?s ?o
				WHERE
				{ ?s verb:usedBy ?o }
				ORDER BY ?s ?o ''' | tabulate }}
			`
		Main.InReader = strings.NewReader(q)
		run("blazegraph report")
	}

	t.Run("infer-none", func(t *testing.T) {
		outputBuffer.Reset()
		run("blazegraph destroy --dataset kb")
		run("blazegraph create --quiet --dataset kb --infer none")
		load()
		report()
		util.LineContentsEqual(t, outputBuffer.String(),
			`*** Statements about verb:uses ***
			p                                           | o
			===========================================================================
			http://www.w3.org/2000/01/rdf-schema#domain | http://tmcphill.net/ns#Person
			http://www.w3.org/2000/01/rdf-schema#range  | http://tmcphill.net/ns#Tool


			*** Statements about tool:keyboard ***
			p                                               | o
			=============================================================================
			http://tmcphill.net/ns/verb#usedBy              | http://tmcphill.net/ns/person#tim
			http://www.w3.org/1999/02/22-rdf-syntax-ns#type | http://tmcphill.net/ns#Tool


			*** Statements about tool:mouse ***
			p                                | o
			====================================================================
			http://tmcphill.net/ns/verb#uses | http://tmcphill.net/ns/person#tim


			*** Statements about person:tim ***
            p                                               | o
            ===============================================================================
            http://tmcphill.net/ns/verb#uses                | http://tmcphill.net/ns/tool#mouse
            http://www.w3.org/1999/02/22-rdf-syntax-ns#type | http://tmcphill.net/ns#Person


			*** What verb:uses what ***
			s                                 | o
			=====================================================================
			http://tmcphill.net/ns/person#tim | http://tmcphill.net/ns/tool#mouse
			http://tmcphill.net/ns/tool#mouse | http://tmcphill.net/ns/person#tim


			*** What is verb:usedBy what ***
			s                                    | o
			========================================================================
			http://tmcphill.net/ns/tool#keyboard | http://tmcphill.net/ns/person#tim

			`)
	})

	t.Run("infer-rdfs", func(t *testing.T) {
		outputBuffer.Reset()
		run("blazegraph destroy --dataset kb")
		run("blazegraph create --quiet --dataset kb --infer rdfs")
		load()
		report()
		util.LineContentsEqual(t, outputBuffer.String(),
			`*** Statements about verb:uses ***
			p                                                  | o
            =====================================================================================================
            http://www.w3.org/1999/02/22-rdf-syntax-ns#type    | http://www.w3.org/1999/02/22-rdf-syntax-ns#Property
            http://www.w3.org/2000/01/rdf-schema#domain        | http://tmcphill.net/ns#Person
            http://www.w3.org/2000/01/rdf-schema#range         | http://tmcphill.net/ns#Tool
            http://www.w3.org/2000/01/rdf-schema#subPropertyOf | http://tmcphill.net/ns/verb#uses


            *** Statements about tool:keyboard ***
            p                                               | o
            =============================================================================
            http://tmcphill.net/ns/verb#usedBy              | http://tmcphill.net/ns/person#tim
            http://www.w3.org/1999/02/22-rdf-syntax-ns#type | http://tmcphill.net/ns#Tool


            *** Statements about tool:mouse ***
            p                                               | o
            ===============================================================================
            http://tmcphill.net/ns/verb#uses                | http://tmcphill.net/ns/person#tim
            http://www.w3.org/1999/02/22-rdf-syntax-ns#type | http://tmcphill.net/ns#Person
            http://www.w3.org/1999/02/22-rdf-syntax-ns#type | http://tmcphill.net/ns#Tool


			*** Statements about person:tim ***
            p                                               | o
            ===============================================================================
            http://tmcphill.net/ns/verb#uses                | http://tmcphill.net/ns/tool#mouse
            http://www.w3.org/1999/02/22-rdf-syntax-ns#type | http://tmcphill.net/ns#Person
			http://www.w3.org/1999/02/22-rdf-syntax-ns#type | http://tmcphill.net/ns#Tool


            *** What verb:uses what ***
            s                                 | o
            =====================================================================
            http://tmcphill.net/ns/person#tim | http://tmcphill.net/ns/tool#mouse
            http://tmcphill.net/ns/tool#mouse | http://tmcphill.net/ns/person#tim


            *** What is verb:usedBy what ***
            s                                    | o
            ========================================================================
            http://tmcphill.net/ns/tool#keyboard | http://tmcphill.net/ns/person#tim

			`)
	})

	t.Run("infer-owl", func(t *testing.T) {
		outputBuffer.Reset()
		run("blazegraph destroy --dataset kb")
		run("blazegraph create --quiet --dataset kb --infer owl")
		load()
		report()
		util.LineContentsEqual(t, outputBuffer.String(),
			`*** Statements about verb:uses ***
			p                                                  | o
			=====================================================================================================
			http://www.w3.org/1999/02/22-rdf-syntax-ns#type    | http://www.w3.org/1999/02/22-rdf-syntax-ns#Property
			http://www.w3.org/2000/01/rdf-schema#domain        | http://tmcphill.net/ns#Person
			http://www.w3.org/2000/01/rdf-schema#range         | http://tmcphill.net/ns#Tool
			http://www.w3.org/2000/01/rdf-schema#subPropertyOf | http://tmcphill.net/ns/verb#uses
			http://www.w3.org/2002/07/owl#inverseOf            | http://tmcphill.net/ns/verb#usedBy


			*** Statements about tool:keyboard ***
			p                                               | o
			=============================================================================
			http://tmcphill.net/ns/verb#usedBy              | http://tmcphill.net/ns/person#tim
			http://www.w3.org/1999/02/22-rdf-syntax-ns#type | http://tmcphill.net/ns#Tool


			*** Statements about tool:mouse ***
			p                                               | o
			===============================================================================
			http://tmcphill.net/ns/verb#usedBy              | http://tmcphill.net/ns/person#tim
			http://tmcphill.net/ns/verb#uses                | http://tmcphill.net/ns/person#tim
			http://www.w3.org/1999/02/22-rdf-syntax-ns#type | http://tmcphill.net/ns#Person
			http://www.w3.org/1999/02/22-rdf-syntax-ns#type | http://tmcphill.net/ns#Tool


			*** Statements about person:tim ***
            p                                               | o
            ===============================================================================
            http://tmcphill.net/ns/verb#usedBy              | http://tmcphill.net/ns/tool#mouse
            http://tmcphill.net/ns/verb#uses                | http://tmcphill.net/ns/tool#keyboard
            http://tmcphill.net/ns/verb#uses                | http://tmcphill.net/ns/tool#mouse
            http://www.w3.org/1999/02/22-rdf-syntax-ns#type | http://tmcphill.net/ns#Person
            http://www.w3.org/1999/02/22-rdf-syntax-ns#type | http://tmcphill.net/ns#Tool


			*** What verb:uses what ***
			s                                 | o
			========================================================================
			http://tmcphill.net/ns/person#tim | http://tmcphill.net/ns/tool#keyboard
			http://tmcphill.net/ns/person#tim | http://tmcphill.net/ns/tool#mouse
			http://tmcphill.net/ns/tool#mouse | http://tmcphill.net/ns/person#tim


			*** What is verb:usedBy what ***
			s                                    | o
			========================================================================
			http://tmcphill.net/ns/person#tim    | http://tmcphill.net/ns/tool#mouse
			http://tmcphill.net/ns/tool#keyboard | http://tmcphill.net/ns/person#tim
			http://tmcphill.net/ns/tool#mouse    | http://tmcphill.net/ns/person#tim

			`)
	})

}

// func TestReportInfer(t *testing.T) {

// 	var outputBuffer strings.Builder
// 	Main.OutWriter = &outputBuffer
// 	Main.ErrWriter = &outputBuffer

// 	load := func() {
// 		Main.InReader = strings.NewReader(`
// 			@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
// 			@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
// 			@prefix tm: <http://tmcphill.net/ns/data#> .
// 		`)
// 		run("blazegraph import --format ttl")
// 	}

// 	report := func() {
// 		q := `
// 			{{ prefix "rdf" "http://www.w3.org/1999/02/22-rdf-syntax-ns#" }}
// 			{{ prefix "tm" "http://tmcphill.net/ns/data#" }}

// 		`
// 		Main.InReader = strings.NewReader(q)
// 		run("blazegraph report")
// 	}

// 	t.Run("infer-none", func(t *testing.T) {
// 		outputBuffer.Reset()
// 		run("blazegraph destroy --dataset kb")
// 		run("blazegraph create --quiet --dataset kb --infer none")
// 		load()
// 		report()
// 		util.LineContentsEqual(t, outputBuffer.String(), `
// 		`)
// 	})

// 	t.Run("infer-rdfs", func(t *testing.T) {
// 		outputBuffer.Reset()
// 		run("blazegraph destroy --dataset kb")
// 		run("blazegraph create --quiet --dataset kb --infer rdfs")
// 		load()
// 		report()
// 		util.LineContentsEqual(t, outputBuffer.String(), `
// 		`)
// 	})

// 	t.Run("infer-owl", func(t *testing.T) {
// 		outputBuffer.Reset()
// 		run("blazegraph destroy --dataset kb")
// 		run("blazegraph create --quiet --dataset kb --infer owl")
// 		load()
// 		report()
// 		util.LineContentsEqual(t, outputBuffer.String(), `
// 		`)
// 	})

// }
