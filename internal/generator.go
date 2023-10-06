package internal

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"text/template"
)

//go:embed gen_zap_object_marshaler.tmpl
var objectMarshalerTmpl string

// ParseAndGenerate parse file and generate new file where the MarshalLogObject methods are written.
func ParseAndGenerate(path string) ([]byte, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseFile: %w", err)
	}
	structFields := make(map[string][]*fieldInfo)
	ast.Inspect(file, func(node ast.Node) bool {
		ts, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return false
		}
		structName := ts.Name.String()
		fields := make([]*fieldInfo, 0, len(st.Fields.List))
		for _, f := range st.Fields.List {
			fi, ok := getFieldInfo(f)
			if !ok {
				continue
			}
			fields = append(fields, fi)
		}
		structFields[structName] = fields
		return true
	})
	packageName := file.Name.Name
	buf := new(bytes.Buffer)
	tmpl := template.Must(template.New("").
		Funcs(template.FuncMap{
			"zeroValueForType":         getZeroValueForType,
			"zapBasicEncoderFunc":      getZapBasicEncoderFunc,
			"zapBasicArrayEncoderFunc": getBasicZapArrayEncoderFunc,
		}).Parse(objectMarshalerTmpl))
	if err := tmpl.Execute(buf, map[string]any{
		"TargetStructs": structFields,
		"PackageName":   packageName,
	}); err != nil {
		return nil, fmt.Errorf("tmpl.Execute: %w", err)
	}
	data, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("format.Source: %w", err)
	}
	return data, nil
}

var (
	zeroValuesForType = map[string]string{
		"bool":       "false",
		"int":        "0",
		"int8":       "0",
		"int16":      "0",
		"int32":      "0",
		"int64":      "0",
		"uint":       "0",
		"uint8":      "0",
		"uint16":     "0",
		"uint32":     "0",
		"uint64":     "0",
		"float32":    "0",
		"float64":    "0",
		"complex64":  "0",
		"complex128": "0",
		"byte":       "0",
		"rune":       "0",
		"string":     `""`,
	}
	zapBasicEncoderFuncForType = map[string]string{
		"bool":          "AddBool",
		"complex128":    "AddComplex128",
		"complex64":     "AddComplex64",
		"float64":       "AddFloat64",
		"float32":       "AddFloat32",
		"int":           "AddInt",
		"int64":         "AddInt64",
		"int32":         "AddInt32",
		"int16":         "AddInt16",
		"int8":          "AddInt8",
		"uint":          "AddUint",
		"uint64":        "AddUint64",
		"uint32":        "AddUint32",
		"uint16":        "AddUint16",
		"uint8":         "AddUint8",
		"string":        "AddString",
		"uintptr":       "AddUintptr",
		"time.Time":     "AddTime",
		"time.Duration": "AddDuration",
		"[]byte":        "AddBinary",
	}
	zapBasicArrayEncoderFuncForType = map[string]string{
		"bool":          "AppendBool",
		"complex128":    "AppendComplex128",
		"complex64":     "AppendComplex64",
		"float64":       "AppendFloat64",
		"float32":       "AppendFloat32",
		"int":           "AppendInt",
		"int64":         "AppendInt64",
		"int32":         "AppendInt32",
		"int16":         "AppendInt16",
		"int8":          "AppendInt8",
		"uint":          "AppendUint",
		"uint64":        "AppendUint64",
		"uint32":        "AppendUint32",
		"uint16":        "AppendUint16",
		"uint8":         "AppendUint8",
		"string":        "AppendString",
		"uintptr":       "AppendUintptr",
		"[]byte":        "AppendByteString",
		"time.Time":     "AppendTime",
		"time.Duration": "AppendDuration",
	}
)

// getZeroValueForType returns zero value for type. if input type is not primitive type, return empty string
func getZeroValueForType(typ string) string {
	if v, ok := zeroValuesForType[typ]; ok {
		return v
	}
	return ""
}

// getZapBasicEncoderFunc returns zapcore.Encoder.AddXXX func name for type. if not found appropriate encoder func, return empty string
func getZapBasicEncoderFunc(typ string) string {
	if v, ok := zapBasicEncoderFuncForType[typ]; ok {
		return v
	}
	return ""
}

// getBasicZapArrayEncoderFunc returns zapcore.ArrayEncoder.AppendXXX func name for type. if not found appropriate encoder func, return empty string
func getBasicZapArrayEncoderFunc(typ string) string {
	if v, ok := zapBasicArrayEncoderFuncForType[typ]; ok {
		return v
	}
	return ""
}
