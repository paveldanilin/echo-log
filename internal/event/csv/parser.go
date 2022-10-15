package csv

import (
	"strings"

	"github.com/paveldanilin/logwatch/internal/event"
)

type FieldDefinition struct {
	event.FieldDefinition
	columnIndex int
}

func NewField(name string, fieldType event.FieldType, columnIndex int) *FieldDefinition {
	field := &FieldDefinition{FieldDefinition: *event.NewFieldDefinition(name, fieldType), columnIndex: columnIndex}
	return field
}

func (field *FieldDefinition) GetColumnIndex() int {
	return field.columnIndex
}

// CSV parser
type parser struct {
	def       *Definition
	separator string
}

func NewCsvParser(def *Definition, separator string) event.Parser {
	return &csvParser{
		def:       def,
		separator: separator,
	}
}

func (parser *csvParser) Parse(text string) *Event {
	csvFields := strings.Split(text, parser.separator)

	e := New()

	for _, f := range parser.def.fields {
		fieldIndex := f.
			e.SetField(f.GetName(), csvFields[fieldIndex])
	}

	return e
}
