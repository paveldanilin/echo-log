package event

type FieldType string

const (
	FIELD_STRING   FieldType = "string"
	FIELD_INT      FieldType = "int"
	FIELD_FLOAT    FieldType = "float"
	FIELD_BOOL     FieldType = "bool"
	FIELD_DATETIME FieldType = "datetime"
)

// --------------------------------------------------------------------------------------------------------------------

type FieldDefinition interface {
	GetName() string
	GetType() FieldType
}

// Base fieldDefinition definition struct
type fieldDefinition struct {
	name      string
	fieldType FieldType
}

func NewFieldDefinition(name string, fieldType FieldType) FieldDefinition {
	return &fieldDefinition{
		name:      name,
		fieldType: fieldType,
	}
}

func (field *fieldDefinition) GetName() string {
	return field.name
}

func (field *fieldDefinition) GetType() FieldType {
	return field.fieldType
}

// --------------------------------------------------------------------------------------------------------------------

type Definition interface {
	SetField(field FieldDefinition)
	GetField(fieldName string) FieldDefinition
	GetFielsNum() int
	GetFields() map[string]FieldDefinition
}

// Base event definition struct
type definition struct {
	fields map[string]FieldDefinition
}

func NewDefinition() Definition {
	return &definition{
		fields: make(map[string]FieldDefinition),
	}
}

func (def *definition) SetField(field FieldDefinition) {
	def.fields[field.GetName()] = field
}

// Return a field definition or nil
func (def *definition) GetField(fieldName string) FieldDefinition {
	if def, ok := def.fields[fieldName]; ok {
		return def
	}
	return nil
}

// Returns a number of fields
func (def *definition) GetFielsNum() int {
	return len(def.fields)
}

func (def *definition) GetFields() map[string]FieldDefinition {
	return def.fields
}
