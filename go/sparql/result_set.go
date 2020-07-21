package sparql

import "encoding/json"

type ResultSet struct {
	Head    Head    `json:"head"`
	Results Results `json:"results"`
}

type Head struct {
	Vars []string `json:"vars"`
}

type Results struct {
	Bindings []Binding `json:"bindings"`
}

type Binding map[string]struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (b Binding) DelimitedValue(name string) (delimitedValue string) {
	value := b[name].Value
	switch b[name].Type {
	case "uri":
		delimitedValue = "<" + value + ">"
	case "literal":
		delimitedValue = "\"" + value + "\""
	}
	return
}

func (rs *ResultSet) JSONString() (jsonString string, err error) {
	jsonBytes, err := json.Marshal(*rs)
	jsonString = string(jsonBytes)
	return
}

func (sr *ResultSet) Vars() []string {
	return sr.Head.Vars
}

func (sr *ResultSet) Bindings() []Binding {
	return sr.Results.Bindings
}

func (sr *ResultSet) BoundValues(bindingIndex int) []string {
	variables := sr.Vars()
	values := make([]string, len(variables))
	binding := sr.Bindings()[bindingIndex]
	for columnIndex, varName := range variables {
		values[columnIndex] = binding[varName].Value
	}
	return values
}
