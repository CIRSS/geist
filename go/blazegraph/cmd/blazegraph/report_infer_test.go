package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/util"
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
		q := `
			{{ prefix "rdf" "http://www.w3.org/1999/02/22-rdf-syntax-ns#" }}
			{{ prefix "tm" "http://tmcphill.net/ns/data#" }}

			{{ select '''
				SELECT ?type
				WHERE
				{ ?type rdf:type rdfs:Class }
				ORDER BY ?type ''' | tabulate }}

			{{ select '''
				SELECT ?a
				WHERE
				{ ?a rdf:type tm:typea }
				ORDER BY ?a ''' | tabulate }}

			{{ select '''
				SELECT ?b
				WHERE
				{ ?b rdf:type tm:typeb }
				ORDER BY ?b ''' | tabulate }}

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
		run("blazegraph create --dataset kb --infer none")
		load()
		report()
		util.LineContentsEqual(t, outputBuffer.String(), `
			type
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
            http://tmcphill.net/ns/data#c1		`)
	})

	t.Run("infer-rdfs", func(t *testing.T) {
		outputBuffer.Reset()
		run("blazegraph destroy --dataset kb")
		run("blazegraph create --dataset kb --infer rdfs")
		load()
		report()
		util.LineContentsEqual(t, outputBuffer.String(), `
			type
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
		run("blazegraph create --dataset kb --infer owl")
		load()
		report()
		util.LineContentsEqual(t, outputBuffer.String(), `
			type
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
