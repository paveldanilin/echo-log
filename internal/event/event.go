package event

import (
	"errors"
	"fmt"
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
	ValueType() ValueType
	Raw() interface{}
	String() string
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
	fieldType ValueType
	raw       interface{}
}

func (fv *fieldValue) ValueType() ValueType {
	return fv.fieldType
}

func (fv *fieldValue) Raw() interface{} {
	return fv.raw
}

func (fv *fieldValue) String() string {
	return fmt.Sprintf("%v", fv.raw)
}

// String
type stringValue struct {
	fieldValue
}

func (s *stringValue) Value() string {
	return s.raw.(string)
}

// Int
type intValue struct {
	fieldValue
}

func (i *intValue) Value() int {
	return i.raw.(int)
}

// Float
type floatValue struct {
	fieldValue
}

func (f *floatValue) Value() float64 {
	return f.raw.(float64)
}

// Bool
type boolValue struct {
	fieldValue
}

func (b *boolValue) Value() bool {
	return b.raw.(bool)
}

// Datetime
type datetimeValue struct {
	fieldValue
}

func (dt *datetimeValue) Value() time.Time {
	return dt.raw.(time.Time)
}

// Creates a new feild value
func NewFieldValue(raw interface{}, fieldType ValueType, parameters map[string]string) (FieldValue, error) {

	if isString(raw) {
		r, err := parseValue(raw.(string), fieldType, parameters)
		if err != nil {
			return nil, err
		}
		raw = r
	}

	v := fieldValue{fieldType: fieldType, raw: raw}

	switch fieldType {
	case VALUE_STRING:
		return &stringValue{fieldValue: v}, nil
	case VALUE_INT:
		return &intValue{fieldValue: v}, nil
	case VALUE_FLOAT:
		return &floatValue{fieldValue: v}, nil
	case VALUE_BOOL:
		return &boolValue{fieldValue: v}, nil
	case VALUE_DATETIME:
		return &datetimeValue{fieldValue: v}, nil
	}

	return nil, errors.New("unknown field type")
}

// Is value string?
func isString(v interface{}) bool {
	if _, ok := v.(string); ok {
		return true
	}
	return false
}

// Converts string to a specific type
func parseValue(s string, fieldType ValueType, parameters map[string]string) (interface{}, error) {
	switch fieldType {
	case VALUE_STRING:
		return s, nil
	case VALUE_INT:
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		return i, nil
	case VALUE_FLOAT:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		return f, nil
	case VALUE_BOOL:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return nil, err
		}
		return b, nil
	case VALUE_DATETIME:
		var dt carbon.Carbon
		if mapHas(parameters, "format") {
			if mapHas(parameters, "tz") {
				dt = carbon.ParseByFormat(s, mapGet(parameters, "format"), mapGet(parameters, "tz"))
			} else {
				dt = carbon.ParseByFormat(s, mapGet(parameters, "format"))
			}
		} else {
			dt = carbon.Parse(s)
		}

		if dt.Error != nil {
			return nil, dt.Error
		}
		return dt.Carbon2Time(), nil
	}
	return nil, errors.New("unknown field type")
}

func mapGet(m map[string]string, name string) string {
	return m[name]
}

func mapHas(m map[string]string, name string) bool {
	_, ok := m[name]
	return ok
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
	switch fv.ValueType() {
	case VALUE_STRING:
		return fv.(StringValue).Value()
	case VALUE_INT:
		return fv.(IntValue).Value()
	case VALUE_FLOAT:
		return fv.(FloatValue).Value()
	case VALUE_BOOL:
		return fv.(BoolValue).Value()
	case VALUE_DATETIME:
		return fv.(DatetimeValue).Value()
	}
	return nil
}

func (event *Event) String(fieldName string) string {
	fv := event.Field(fieldName)
	if fv == nil {
		return ""
	}
	return fv.String()
}

func (event *Event) StringValue(fieldName string) StringValue {
	fv := event.Field(fieldName)
	if fv == nil {
		return nil
	}
	if fv.ValueType() == VALUE_STRING {
		return fv.(StringValue)
	}
	return nil
}

func (event *Event) IntValue(fieldName string) IntValue {
	fv := event.Field(fieldName)
	if fv == nil {
		return nil
	}
	if fv.ValueType() == VALUE_INT {
		return fv.(IntValue)
	}
	return nil
}

func (event *Event) FloatValue(fieldName string) FloatValue {
	fv := event.Field(fieldName)
	if fv == nil {
		return nil
	}
	if fv.ValueType() == VALUE_FLOAT {
		return fv.(FloatValue)
	}
	return nil
}

func (event *Event) BoolValue(fieldName string) BoolValue {
	fv := event.Field(fieldName)
	if fv == nil {
		return nil
	}
	if fv.ValueType() == VALUE_FLOAT {
		return fv.(BoolValue)
	}
	return nil
}

func (event *Event) DatetimeValue(fieldName string) DatetimeValue {
	fv := event.Field(fieldName)
	if fv == nil {
		return nil
	}
	if fv.ValueType() == VALUE_DATETIME {
		return fv.(DatetimeValue)
	}
	return nil
}

func (event *Event) SetField(fieldName string, value FieldValue) {
	event.fields[fieldName] = value
}

func (event *Event) SetValue(fieldName string, value interface{}, fieldType ValueType, parameters map[string]string) error {
	if parameters == nil {
		parameters = map[string]string{}
	}
	v, err := NewFieldValue(value, fieldType, parameters)
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
