package event

// --------------------------------------------------------------------------------------------------------------------
// Field definition
// --------------------------------------------------------------------------------------------------------------------

type FieldType int

const (
	FIELD_STRING   FieldType = 1
	FIELD_NUMBER   FieldType = 2
	FIELD_DATETIME FieldType = 3
)

type FieldDefinition struct {
	name       string
	fieldType  FieldType
	parameters map[string]interface{}
}

func NewFieldDefinition(name string, fieldType FieldType) *FieldDefinition {
	return &FieldDefinition{
		name:       name,
		fieldType:  fieldType,
		parameters: make(map[string]interface{}),
	}
}

func (field *FieldDefinition) GetName() string {
	return field.name
}

func (field *FieldDefinition) GetFieldType() FieldType {
	return field.fieldType
}

func (field *FieldDefinition) SetParam(name string, value interface{}) {
	field.parameters[name] = value
}

func (field *FieldDefinition) GetParam(name string) interface{} {
	if v, ok := field.parameters[name]; ok {
		return v
	}
	return nil
}

func (field *FieldDefinition) HasParam(name string) bool {
	_, ok := field.parameters[name]
	return ok
}

func (field *FieldDefinition) GetIntParam(name string) int {
	return field.GetParam(name).(int)
}

// --------------------------------------------------------------------------------------------------------------------
// Event Definition
// --------------------------------------------------------------------------------------------------------------------

type Definition struct {
	fields map[string]*FieldDefinition
}

func NewDefinition(fieldDefinitions []*FieldDefinition) *Definition {
	def := &Definition{
		fields: make(map[string]*FieldDefinition),
	}
	for _, field := range fieldDefinitions {
		def.fields[field.GetName()] = field
	}
	return def
}

// Returns a number of fields
func (def *Definition) GetSize() int {
	return len(def.fields)
}

// Return a field definition or nil
func (def *Definition) GetFieldDefinition(fieldName string) *FieldDefinition {
	if def, ok := def.fields[fieldName]; ok {
		return def
	}
	return nil
}
