package event

import (
	"regexp"
	"strings"
)

const csv_field_name_index = "csv.field_index"

type Parser interface {
	Parse(text string) *Event
}

// --------------------------------------------------------------------------------------------------------------------
// CSV parser
// --------------------------------------------------------------------------------------------------------------------

type csvParser struct {
	def       *Definition
	separator string
}

func NewCsvParser(def *Definition, separator string) Parser {
	return &csvParser{
		def:       def,
		separator: separator,
	}
}

func (parser *csvParser) Parse(text string) *Event {
	fields := strings.Split(text, parser.separator)

	e := New()

	for _, f := range parser.def.fields {
		fieldIndex := f.GetIntParam(csv_field_name_index)
		e.SetField(f.GetName(), fields[fieldIndex])
	}

	return e
}

// --------------------------------------------------------------------------------------------------------------------
// Pattern parser
// --------------------------------------------------------------------------------------------------------------------

type patternDefinition struct {
	def *Definition
	re  *regexp.Regexp
}

type patternParser struct {
	patterns []*patternDefinition
}

func NewPatternParser() Parser {
	return &patternParser{
		patterns: make([]*patternDefinition, 0),
	}
}

func (parser *patternParser) RegisterPattern(pattern string, def *Definition) {
	parser.patterns = append(parser.patterns, &patternDefinition{
		def: def,
		re:  regexp.MustCompile(pattern),
	})
}

func (parser *patternParser) Parse(text string) *Event {
	pat, matches := parser.findPattern(text)

	if pat == nil {
		return nil
	}

	e := New()

	for _, fieldDef := range pat.def.fields {
		nameIndex := pat.re.SubexpIndex(fieldDef.GetName())
		e.SetField(fieldDef.GetName(), strings.Trim(matches[nameIndex], " "))
	}

	return e
}

func (parser *patternParser) findPattern(text string) (*patternDefinition, []string) {
	for _, pat := range parser.patterns {
		matches := pat.re.FindStringSubmatch(text)
		if matches != nil {
			return pat, matches
		}
	}

	return nil, nil
}
