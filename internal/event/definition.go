package event

// --------------------------------------------------------------------------------------------------------------------
// Field definition
// --------------------------------------------------------------------------------------------------------------------

type FieldType int

const (
	FIELD_STRING   FieldType = 1
	FIELD_NUMBER   FieldType = 2
	FIELD_BOOL     FieldType = 3
	FIELD_DATETIME FieldType = 4
)

type FieldDefinition struct {
	name      string
	fieldType FieldType
}

func NewFieldDefinition(name string, fieldType FieldType) *FieldDefinition {
	return &FieldDefinition{
		name:      name,
		fieldType: fieldType,
	}
}

func (field *FieldDefinition) GetName() string {
	return field.name
}

func (field *FieldDefinition) GetFieldType() FieldType {
	return field.fieldType
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
func (def *Definition) GetFielsNum() int {
	return len(def.fields)
}

// Return a field definition or nil
func (def *Definition) GetField(fieldName string) *FieldDefinition {
	if def, ok := def.fields[fieldName]; ok {
		return def
	}
	return nil
}
