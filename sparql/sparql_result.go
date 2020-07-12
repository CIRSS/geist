package sparql

type SparqlResult struct {
	Head struct {
		Vars []string
	}
	Results struct {
		Bindings []Binding
	}
}

func (sr *SparqlResult) Vars() []string {
	return sr.Head.Vars
}

func (sr *SparqlResult) Bindings() []Binding {
	return sr.Results.Bindings
}

type Binding map[string]TypedValue

type TypedValue struct {
	Type  string
	Value string
}
