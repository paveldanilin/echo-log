package event

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-module/carbon/v2"
)

type Parser interface {
	Parse(text string) (*Event, error)
}

// --------------------------------------------------------------------------------------------------------------------

// Field value
type FieldValue interface {
	GetType() FieldType
	Raw() string
}

type StringValue interface {
	Value() string
}

type IntValue interface {
	Value() int
}

type FloatValue interface {
	Value() float64
}

type BoolValue interface {
	Value() bool
}

type DatetimeValue interface {
	Value() time.Time
}

// Base field value
type fieldValue struct {
	fieldType FieldType
	raw       string
}

func (fv *fieldValue) GetType() FieldType {
	return fv.fieldType
}

func (fv *fieldValue) Raw() string {
	return fv.raw
}

// String
type stringValue struct {
	fieldValue
}

func (s *stringValue) Value() string {
	return s.raw
}

// Int
type intValue struct {
	fieldValue
	value int
}

func (i *intValue) Value() int {
	return i.value
}

// Float
type floatValue struct {
	fieldValue
	value float64
}

func (f *floatValue) Value() float64 {
	return f.value
}

// Bool
type boolValue struct {
	fieldValue
	value bool
}

func (b *boolValue) Value() bool {
	return b.value
}

// Datetime
type datetimeValue struct {
	fieldValue
	value time.Time
}

func (dt *datetimeValue) Value() time.Time {
	return dt.value
}

func NewFieldValue(raw string, fieldType FieldType) (FieldValue, error) {
	v := fieldValue{fieldType: fieldType, raw: raw}
	switch fieldType {
	case FIELD_STRING:
		return &stringValue{fieldValue: v}, nil
	case FIELD_INT:
		i, err := strconv.Atoi(raw)
		if err != nil {
			return nil, err
		}
		return &intValue{fieldValue: v, value: i}, nil
	case FIELD_FLOAT:
		f, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return nil, err
		}
		return &floatValue{fieldValue: v, value: f}, nil
	case FIELD_BOOL:
		b, err := strconv.ParseBool(raw)
		if err != nil {
			return nil, err
		}
		return &boolValue{fieldValue: v, value: b}, nil
	case FIELD_DATETIME:
		dt := carbon.Parse(raw)
		if dt.Error != nil {
			return nil, dt.Error
		}
		return &datetimeValue{fieldValue: v, value: dt.Carbon2Time()}, nil
	}
	return nil, errors.New("unknown field type")
}

// --------------------------------------------------------------------------------------------------------------------

// Event
type Event struct {
	fields map[string]FieldValue
}

func New() *Event {
	return &Event{
		fields: make(map[string]FieldValue),
	}
}

func (event *Event) Has(fieldName string) bool {
	_, ok := event.fields[fieldName]
	return ok
}

func (event *Event) Field(fieldName string) FieldValue {
	if v, ok := event.fields[fieldName]; ok {
		return v
	}
	return nil
}

func (event *Event) Value(fieldName string) interface{} {
	fv := event.Field(fieldName)
	if fv == nil {
		return nil
	}
	switch fv.GetType() {
	case FIELD_STRING:
		return fv.(StringValue).Value()
	case FIELD_INT:
		return fv.(IntValue).Value()
	case FIELD_FLOAT:
		return fv.(FloatValue).Value()
	case FIELD_BOOL:
		return fv.(BoolValue).Value()
	case FIELD_DATETIME:
		return fv.(DatetimeValue).Value()
	}
	return nil
}

func (event *Event) String(fieldName string) string {
	fv := event.Field(fieldName)
	if fv == nil {
		return ""
	}
	return fv.Raw()
}

func (event *Event) StringValue(fieldName string) StringValue {
	fv := event.Field(fieldName)
	if fv == nil {
		return nil
	}
	if fv.GetType() == FIELD_STRING {
		return fv.(StringValue)
	}
	return nil
}

func (event *Event) IntValue(fieldName string) IntValue {
	fv := event.Field(fieldName)
	if fv == nil {
		return nil
	}
	if fv.GetType() == FIELD_INT {
		return fv.(IntValue)
	}
	return nil
}

func (event *Event) FloatValue(fieldName string) FloatValue {
	fv := event.Field(fieldName)
	if fv == nil {
		return nil
	}
	if fv.GetType() == FIELD_FLOAT {
		return fv.(FloatValue)
	}
	return nil
}

func (event *Event) BoolValue(fieldName string) BoolValue {
	fv := event.Field(fieldName)
	if fv == nil {
		return nil
	}
	if fv.GetType() == FIELD_FLOAT {
		return fv.(BoolValue)
	}
	return nil
}

func (event *Event) DatetimeValue(fieldName string) DatetimeValue {
	fv := event.Field(fieldName)
	if fv == nil {
		return nil
	}
	if fv.GetType() == FIELD_DATETIME {
		return fv.(DatetimeValue)
	}
	return nil
}

func (event *Event) SetField(fieldName string, value FieldValue) {
	event.fields[fieldName] = value
}

func (event *Event) SetValue(fieldName string, value string, fieldType FieldType) error {
	v, err := NewFieldValue(value, fieldType)
	if err != nil {
		return err
	}
	event.fields[fieldName] = v
	return nil
}

func (event *Event) Fields() map[string]FieldValue {
	return event.fields
}

// Returns an array of field names
func (event *Event) FieldNames() []string {
	keys := make([]string, len(event.fields))
	i := 0
	for k := range event.fields {
		keys[i] = k
		i++
	}
	return keys
}

// Returns a values map
func (event *Event) Map() map[string]interface{} {
	vMap := make(map[string]interface{})
	for n := range event.fields {
		vMap[n] = event.Value(n)
	}
	return vMap
}
