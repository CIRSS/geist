package sparql

type SparqlResult struct {
	Head    HeadT
	Results ResultsT
}

type HeadT struct {
	Vars []string
}

type ResultsT struct {
	Bindings []BindingT
}

type BindingT map[string]TypedValueT

type TypedValueT struct {
	Type  string
	Value string
}

func (sr *SparqlResult) Vars() []string {
	return sr.Head.Vars
}

func (sr *SparqlResult) Bindings() []BindingT {
	return sr.Results.Bindings
}

func (sr *SparqlResult) Row(rowIndex int) (values []string) {
	binding := sr.Bindings()[rowIndex]
	for columnIndex, varName := range sr.Vars() {
		values[columnIndex] = binding[varName].Value
	}
	return values
}
