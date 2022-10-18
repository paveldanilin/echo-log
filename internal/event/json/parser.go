package json

import (
	"errors"
	"strings"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"github.com/paveldanilin/logwatch/internal/event"
)

// JSON event parser
// Example:
// 		- {"eventdate": 12345678, "level": "ERROR", "message": "Fatal error"}

// --------------------------------------------------------------------------------------------------------------------

type FieldDefinition struct {
	event.FieldDefinition
	pathExpr jp.Expr
}

func NewFieldDefinition(name string, fieldType event.ValueType, jsonPath string) *FieldDefinition {
	return &FieldDefinition{
		FieldDefinition: event.NewFieldDefinition(name, fieldType),
		pathExpr:        jp.MustParseString(jsonPath),
	}
}

// --------------------------------------------------------------------------------------------------------------------

type EventDefinition struct {
	event.Definition
}

func NewEventDefition() *EventDefinition {
	return &EventDefinition{
		Definition: event.NewDefinition(),
	}
}

func (def *EventDefinition) SetField(field *FieldDefinition) {
	def.Definition.SetField(field)
}

func (def *EventDefinition) GetField(fieldName string) *FieldDefinition {
	f := def.Definition.Field(fieldName)
	return f.(*FieldDefinition)
}

// --------------------------------------------------------------------------------------------------------------------

type parser struct {
	definition *EventDefinition
}

func NewParser(definition *EventDefinition) event.Parser {
	return &parser{
		definition: definition,
	}
}

func (p *parser) Parse(text string) (*event.Event, error) {
	text = strings.TrimSpace(text)
	if len(text) == 0 {
		// Nothing to parse, text is empty
		return nil, errors.New("text is empty")
	}

	obj, err := oj.ParseString(text)
	if err != nil {
		return nil, err
	}

	e := event.New()

	for fieldName, fieldDefinition := range p.definition.Fields() {
		v := fieldDefinition.(*FieldDefinition).pathExpr.First(obj)
		if v == nil {
			v = ""
		}
		e.SetValue(fieldName, v, fieldDefinition.ValueType())
	}

	return e, nil
}
