package internal

import (
	"go/ast"
	"go/types"
	"reflect"
	"strings"
)

type fieldKind int

const (
	fieldKindNormal fieldKind = iota
	fieldKindPrt
	fieldKindSlice
	fieldKindByteSlice
)

func (k fieldKind) IsPtr() bool {
	return k == fieldKindPrt
}

func (k fieldKind) IsSlice() bool {
	return k == fieldKindSlice
}

func (k fieldKind) IsByteSlice() bool {
	return k == fieldKindByteSlice
}

func (k fieldKind) IsNormal() bool {
	return k == fieldKindNormal
}

type fieldInfo struct {
	ZapKey               string
	FieldName            string
	FieldType            string
	OriginalFieldType    string
	FieldKind            fieldKind
	ElementType          string
	OriginalElementType  string
	ElementKind          fieldKind
	OmitEmpty            bool
	ForceObjectMarshaler bool
}

type tagOptions string

// parseTag splits a struct fieldInfo's json tag into its name and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {
	tag, opt, _ := strings.Cut(tag, ",")
	return tag, tagOptions(opt)
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var name string
		name, s, _ = strings.Cut(s, ",")
		if name == optionName {
			return true
		}
	}
	return false
}

func getFieldInfo(f *ast.Field) (*fieldInfo, bool) {
	if len(f.Names) == 0 {
		return nil, false
	}
	fieldName := f.Names[0].Name
	var zapKey string
	var opts tagOptions
	if f.Tag != nil {
		tagValue := strings.ReplaceAll(f.Tag.Value, "`", "")
		tags := reflect.StructTag(tagValue)
		tag := tags.Get("zap")
		if tag == "-" {
			return nil, false
		}
		zapKey, opts = parseTag(tag)
	}

	if zapKey == "" {
		zapKey = fieldName
	}
	fKind := fieldKindNormal
	eKind := fieldKindNormal
	fieldType := types.ExprString(f.Type)
	origFieldType := fieldType
	var elementType string
	var origElementType string
	switch v := f.Type.(type) {
	case *ast.StarExpr:
		fKind = fieldKindPrt
		origFieldType = strings.TrimPrefix(fieldType, "*")
	case *ast.ArrayType:
		if fieldType == "[]byte" {
			fKind = fieldKindByteSlice
		} else {
			fKind = fieldKindSlice
			if _, ok := v.Elt.(*ast.StarExpr); ok {
				eKind = fieldKindPrt
			}
			elementType = types.ExprString(v.Elt)
			origElementType = strings.TrimPrefix(elementType, "*")
		}
	}
	return &fieldInfo{
		ZapKey:              zapKey,
		FieldKind:           fKind,
		FieldName:           fieldName,
		FieldType:           fieldType,
		OriginalFieldType:   origFieldType,
		ElementType:         elementType,
		ElementKind:         eKind,
		OriginalElementType: origElementType,
		OmitEmpty:           opts.Contains("omitempty"),
	}, true
}
