package blazegraph

type Properties struct {
	Name      string
	Inference string
	Quads     bool
}

func NewProperties(name string) *Properties {
	p := new(Properties)
	p.Name = name
	p.Inference = "none"
	return p
}

func (p *Properties) String() string {

	s := "com.bigdata.rdf.sail.namespace=" + p.Name + "\n"

	switch p.Inference {
	case "none":
		s += "com.bigdata.rdf.sail.truthMaintenance=false\n"
		s += "com.bigdata.rdf.store.AbstractTripleStore.axiomsClass=com.bigdata.rdf.axioms.NoAxioms\n"
	case "rdfs":
		s += "com.bigdata.rdf.sail.truthMaintenance=true\n"
		s += "com.bigdata.rdf.store.AbstractTripleStore.axiomsClass=com.bigdata.rdf.axioms.RdfsAxioms\n"
	case "owl":
		s += "com.bigdata.rdf.sail.truthMaintenance=true\n"
		s += "com.bigdata.rdf.store.AbstractTripleStore.axiomsClass=com.bigdata.rdf.axioms.OwlAxioms\n"
	}

	if p.Quads {
		s += "com.bigdata.rdf.store.AbstractTripleStore.quads=true\n"
	} else {
		s += "com.bigdata.rdf.store.AbstractTripleStore.quads=false\n"
	}

	return s
}
