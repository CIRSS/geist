package rdf

import "strings"

type Uri struct {
	Value    string
	Prefixed bool
}

func (u Uri) String() string {
	if u.Prefixed && !strings.Contains(u.Value, "/") {
		return u.Value
	}
	return "<" + u.Value + ">"
}
