package sparql

type SparqlResult struct {
	Head    Head
	Results Results
}

type Head struct {
	Vars []string
}

type Results struct {
	Bindings []Binding
}

type Binding map[string]struct {
	Type  string
	Value string
}

func (sr *SparqlResult) Vars() []string {
	return sr.Head.Vars
}

func (sr *SparqlResult) Bindings() []Binding {
	return sr.Results.Bindings
}

func (sr *SparqlResult) Row(rowIndex int) (values []string) {
	binding := sr.Bindings()[rowIndex]
	for columnIndex, varName := range sr.Vars() {
		values[columnIndex] = binding[varName].Value
	}
	return values
}
