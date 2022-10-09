package event

import "strings"

type Parser interface {
	Parse(text string) *Event
}

type csvParser struct {
	def *Definition
}

func NewCsvParser(def *Definition) Parser {
	return &csvParser{
		def: def,
	}
}

func (parser *csvParser) Parse(text string) *Event {
	fields := strings.Split(text, ";")

	e := New()

	for _, f := range parser.def.fieldDefinitions {
		e.SetField(f.GetName(), fields[f.GetIntParam("csv_column_index")])
	}

	return e
}
