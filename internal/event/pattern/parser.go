package pattern

import (
	"errors"
	"regexp"
	"strings"

	"github.com/paveldanilin/logwatch/internal/event"
)

// --------------------------------------------------------------------------------------------------------------------

// Pattern field definition
type FieldDefinition struct {
	event.FieldDefinition
	groupName string // Regex group name
}

func NewFieldDefinition(name string, fieldType event.ValueType, groupName string) *FieldDefinition {
	return &FieldDefinition{
		FieldDefinition: event.NewFieldDefinition(name, fieldType),
		groupName:       groupName,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// Pattern event defintiion
type EventDefinition struct {
	event.Definition
	re *regexp.Regexp
}

func NewEventDefition(pattern string) *EventDefinition {
	return &EventDefinition{
		Definition: event.NewDefinition(),
		re:         regexp.MustCompile(pattern),
	}
}

func (def *EventDefinition) SetField(field *FieldDefinition) event.FieldDefinition {
	return def.Definition.SetField(field)
}

func (def *EventDefinition) GetField(fieldName string) *FieldDefinition {
	f := def.Definition.Field(fieldName)
	return f.(*FieldDefinition)
}

// --------------------------------------------------------------------------------------------------------------------

// Pattern parser
type parser struct {
	definitions []*EventDefinition
}

func NewParser(definitions []*EventDefinition) event.Parser {
	return &parser{
		definitions: definitions,
	}
}

// TODO: add multiline support
// Returns an event object from text
func (p *parser) Parse(text string) (*event.Event, error) {
	text = strings.TrimSpace(text)
	if len(text) == 0 {
		// Nothing to parse, text is empty
		return nil, errors.New("text is empty")
	}

	def, matches := p.findDefinition(text)

	if def == nil {
		return nil, errors.New("event defintion not found")
	}

	e := event.New()

	for fieldName, fieldDefinition := range def.Fields() {
		groupIndex := def.re.SubexpIndex(fieldDefinition.(*FieldDefinition).groupName)
		rawText := strings.TrimSpace(matches[groupIndex])
		err := e.SetValue(fieldName, rawText, fieldDefinition.ValueType(), fieldDefinition.Parameters())
		if err != nil {
			return nil, err
		}
	}

	return e, nil
}

// Returns an event defintion as first value and captured group values
func (p *parser) findDefinition(text string) (*EventDefinition, []string) {
	for _, d := range p.definitions {
		matches := d.re.FindStringSubmatch(text)
		if matches != nil {
			return d, matches
		}
	}

	return nil, nil
}
