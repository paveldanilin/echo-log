package csv

import (
	"errors"
	"fmt"
	"strings"

	"github.com/paveldanilin/logwatch/internal/event"
)

// --------------------------------------------------------------------------------------------------------------------

// CSV field definition
type FieldDefinition struct {
	event.FieldDefinition
	columnIndex int
}

func NewFieldDefinition(name string, fieldType event.FieldType, columnIndex int) *FieldDefinition {
	return &FieldDefinition{
		FieldDefinition: event.NewFieldDefinition(name, fieldType),
		columnIndex:     columnIndex,
	}
}

func (field *FieldDefinition) ColumnIndex() int {
	return field.columnIndex
}

// --------------------------------------------------------------------------------------------------------------------

// CSV event definition
type EventDefinition struct {
	event.Definition
}

func NewEventDefition() *EventDefinition {
	return &EventDefinition{Definition: event.NewDefinition()}
}

func (def *EventDefinition) SetField(field *FieldDefinition) {
	def.Definition.SetField(field)
}

func (def *EventDefinition) GetField(fieldName string) *FieldDefinition {
	f := def.Definition.GetField(fieldName)
	return f.(*FieldDefinition)
}

// --------------------------------------------------------------------------------------------------------------------

// CSV parser
type parser struct {
	definition *EventDefinition
	separator  string
}

func NewParser(eventDef *EventDefinition, separator string) event.Parser {
	return &parser{
		definition: eventDef,
		separator:  separator,
	}
}

func (p *parser) Parse(text string) (*event.Event, error) {
	text = strings.TrimSpace(text)
	if len(text) == 0 {
		// Nothing to parse, text is empty
		return nil, errors.New("text is empty")
	}

	csvFields := strings.Split(text, p.separator)
	csvFieldsNum := len(csvFields)

	e := event.New()

	for fieldName, fieldDefinition := range p.definition.GetFields() {
		fieldIndex := fieldDefinition.(*FieldDefinition).columnIndex
		if fieldIndex > csvFieldsNum {
			return nil, fmt.Errorf("[%s]: column index `%d` does not exist", fieldName, fieldIndex)
		}
		err := e.SetValue(fieldName, csvFields[fieldIndex], fieldDefinition.GetType())
		if err != nil {
			return nil, err
		}
	}

	return e, nil
}
