package event

import "errors"

// ParametersMap
type ParametersMap struct {
	parameters map[string]interface{}
}

func NewParametersMap(parameters map[string]interface{}) *ParametersMap {
	return &ParametersMap{
		parameters: parameters,
	}
}

func (p *ParametersMap) Set(name string, value interface{}) {
	p.parameters[name] = value
}

func (p *ParametersMap) Get(name string) interface{} {
	if v, ok := p.parameters[name]; ok {
		return v
	}
	return nil
}

func (p *ParametersMap) Has(name string) bool {
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

func NewFieldDefinition(name string, fieldType FieldType, parameters map[string]interface{}) *FieldDefinition {
	return &FieldDefinition{
		name:          name,
		fieldType:     fieldType,
		parametersMap: NewParametersMap(parameters),
	}
}

func (fieldDef *FieldDefinition) GetName() string {
	return fieldDef.name
}

func (fieldDef *FieldDefinition) GetFieldType() FieldType {
	return fieldDef.fieldType
}

func (fieldDef *FieldDefinition) GetParam(name string) interface{} {
	return fieldDef.parametersMap.Get(name)
}

func (fieldDef *FieldDefinition) GetIntParam(name string) int {
	return fieldDef.parametersMap.Get(name).(int)
}

func (fieldDef *FieldDefinition) HasParam(name string) bool {
	return fieldDef.parametersMap.Has(name)
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
