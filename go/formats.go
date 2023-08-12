package main

import (
	"sort"
)

type format struct {
	Extension   string
	PerLanguage bool
	Template    string
}

type formatMap map[string]*format

var formats = formatMap{ //nolint:gochecknoglobals
	"go":         &golang,
	"goi18n":     &goi18n,
	"typescript": &typescript,
}

// Keys is a sorted list of names.
func (f *formatMap) Keys() []string {
	languages := []string{}

	for k := range *f {
		languages = append(languages, k)
	}

	sort.Strings(languages)

	return languages
}
