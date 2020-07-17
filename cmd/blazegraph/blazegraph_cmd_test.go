package main

import (
	"os"
	"strings"
)

func runWithArgs(commandLine string) {
	Main.ErrWriter = os.Stdout
	os.Args = strings.Fields(commandLine)
	Main.Run()
}

func ExampleBlazegraphCmd_drop_then_dump() {
	runWithArgs("blazegraph drop")
	runWithArgs("blazegraph dump")
	// Output:
	//
}

func ExampleBlazegraphCmd_drop_load_turtle_then_dump() {
	runWithArgs("blazegraph drop")
	runWithArgs("blazegraph -f testdata/in.nt load")
	runWithArgs("blazegraph dump")
	// Output:
	// <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://tmcphill.net/tags#tag> .
	// <http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
}

func ExampleBlazegraphCmd_drop_load_jsonld_then_dump() {
	runWithArgs("blazegraph drop")
	runWithArgs("blazegraph -f testdata/address-book.jsonld load-jsonld")
	runWithArgs("blazegraph dump")
	// Output:
	// <http://learningsparql.com/ns/addressbook#email> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://learningsparql.com/ns/addressbook#email> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://learningsparql.com/ns/addressbook#email> .
	// <http://learningsparql.com/ns/addressbook#firstname> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://learningsparql.com/ns/addressbook#firstname> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://learningsparql.com/ns/addressbook#firstname> .
	// <http://learningsparql.com/ns/addressbook#homeTel> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://learningsparql.com/ns/addressbook#homeTel> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://learningsparql.com/ns/addressbook#homeTel> .
	// <http://learningsparql.com/ns/addressbook#lastname> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://learningsparql.com/ns/addressbook#lastname> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://learningsparql.com/ns/addressbook#lastname> .
	// <http://learningsparql.com/ns/addressbook#mobileTel> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://learningsparql.com/ns/addressbook#mobileTel> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://learningsparql.com/ns/addressbook#mobileTel> .
	// <http://learningsparql.com/ns/addressbook#nickname> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://learningsparql.com/ns/addressbook#nickname> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://learningsparql.com/ns/addressbook#nickname> .
	// <http://learningsparql.com/ns/data#i0432> <http://learningsparql.com/ns/addressbook#email> "richard49@hotmail.com" .
	// <http://learningsparql.com/ns/data#i0432> <http://learningsparql.com/ns/addressbook#firstname> "Richard" .
	// <http://learningsparql.com/ns/data#i0432> <http://learningsparql.com/ns/addressbook#homeTel> "(229) 276-5135" .
	// <http://learningsparql.com/ns/data#i0432> <http://learningsparql.com/ns/addressbook#lastname> "Mutt" .
	// <http://learningsparql.com/ns/data#i0432> <http://learningsparql.com/ns/addressbook#nickname> "Dick" .
	// <http://learningsparql.com/ns/data#i8301> <http://learningsparql.com/ns/addressbook#email> "c.ellis@usairwaysgroup.com" .
	// <http://learningsparql.com/ns/data#i8301> <http://learningsparql.com/ns/addressbook#email> "craigellis@yahoo.com" .
	// <http://learningsparql.com/ns/data#i8301> <http://learningsparql.com/ns/addressbook#firstname> "Craig" .
	// <http://learningsparql.com/ns/data#i8301> <http://learningsparql.com/ns/addressbook#homeTel> "(194) 966-1505" .
	// <http://learningsparql.com/ns/data#i8301> <http://learningsparql.com/ns/addressbook#lastname> "Ellis" .
	// <http://learningsparql.com/ns/data#i9771> <http://learningsparql.com/ns/addressbook#email> "cindym@gmail.com" .
	// <http://learningsparql.com/ns/data#i9771> <http://learningsparql.com/ns/addressbook#firstname> "Cindy" .
	// <http://learningsparql.com/ns/data#i9771> <http://learningsparql.com/ns/addressbook#homeTel> "(245) 646-5488" .
	// <http://learningsparql.com/ns/data#i9771> <http://learningsparql.com/ns/addressbook#lastname> "Marshall" .
	// <http://learningsparql.com/ns/data#i9771> <http://learningsparql.com/ns/addressbook#mobileTel> "(245) 732-8991" .
}
