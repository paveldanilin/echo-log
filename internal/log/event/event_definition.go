package event

import "errors"

// ParametersMap
type ParametersMap struct {
	parameters map[string]string
}

func NewParametersMap() *ParametersMap {
	return &ParametersMap{
		parameters: make(map[string]string),
	}
}

func (p *ParametersMap) SetParameter(name string, value string) {
	p.parameters[name] = value
}

func (p *ParametersMap) GetParameter(name string) (string, error) {
	if v, ok := p.parameters[name]; ok {
		return v, nil
	}
	return "", errors.New("parameter not found")
}

func (p *ParametersMap) HasParameter(name string) bool {
	_, ok := p.parameters[name]
	return ok
}

// Field definition

type FieldType int

const (
	FIELD_STRING   FieldType = 1
	FIELD_NUMBER   FieldType = 2
	FIELD_DATETIME FieldType = 3
)

type FieldDefinition struct {
	name          string
	fieldType     FieldType
	parametersMap *ParametersMap
}

func NewFieldDefinition(name string, fieldType FieldType) *FieldDefinition {
	return &FieldDefinition{
		name:          name,
		fieldType:     fieldType,
		parametersMap: NewParametersMap(),
	}
}

func (fieldDef *FieldDefinition) GetName() string {
	return fieldDef.name
}

func (fieldDef *FieldDefinition) GetFieldType() FieldType {
	return fieldDef.fieldType
}

func (fieldDef *FieldDefinition) GetParametersMap() *ParametersMap {
	return fieldDef.parametersMap
}

// Definition

type Definition struct {
	fieldDefinitions []*FieldDefinition
}

func NewDefinition(fieldDefinitions []*FieldDefinition) *Definition {
	return &Definition{
		fieldDefinitions: fieldDefinitions,
	}
}

func (def *Definition) GetFieldDefinition(fieldName string) (*FieldDefinition, error) {
	for _, fd := range def.fieldDefinitions {
		if fd.GetName() == fieldName {
			return fd, nil
		}
	}
	return nil, errors.New("field definition not found")
}
