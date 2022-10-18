package event

type ValueType string

const (
	VALUE_STRING   ValueType = "string"
	VALUE_INT      ValueType = "int"
	VALUE_FLOAT    ValueType = "float"
	VALUE_BOOL     ValueType = "bool"
	VALUE_DATETIME ValueType = "datetime"
)

// --------------------------------------------------------------------------------------------------------------------

type FieldDefinition interface {
	Name() string
	ValueType() ValueType
}

// Base fieldDefinition definition struct
type fieldDefinition struct {
	name      string
	fieldType ValueType
}

func NewFieldDefinition(name string, fieldType ValueType) FieldDefinition {
	return &fieldDefinition{
		name:      name,
		fieldType: fieldType,
	}
}

func (field *fieldDefinition) Name() string {
	return field.name
}

func (field *fieldDefinition) ValueType() ValueType {
	return field.fieldType
}

// --------------------------------------------------------------------------------------------------------------------

type Definition interface {
	SetField(field FieldDefinition)
	Field(fieldName string) FieldDefinition
	FieldsNum() int
	Fields() map[string]FieldDefinition
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
	def.fields[field.Name()] = field
}

// Return a field definition or nil
func (def *definition) Field(fieldName string) FieldDefinition {
	if def, ok := def.fields[fieldName]; ok {
		return def
	}
	return nil
}

// Returns a number of fields
func (def *definition) FieldsNum() int {
	return len(def.fields)
}

func (def *definition) Fields() map[string]FieldDefinition {
	return def.fields
}
