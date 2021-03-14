package blaze

type DatasetProperties struct {
	Name      string
	Inference string
	Quads     bool
}

func NewDatasetProperties(name string) *DatasetProperties {
	dp := new(DatasetProperties)
	dp.Name = name
	dp.Inference = "none"
	return dp
}

func (dp *DatasetProperties) String() string {

	s := "com.bigdata.rdf.sail.namespace=" + dp.Name + "\n"

	switch dp.Inference {
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

	if dp.Quads {
		s += "com.bigdata.rdf.store.AbstractTripleStore.quads=true\n"
	} else {
		s += "com.bigdata.rdf.store.AbstractTripleStore.quads=false\n"
	}

	return s
}
