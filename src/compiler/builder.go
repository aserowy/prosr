package compiler

import (
	"bytes"
	"io/ioutil"
	"strings"
	"text/template"

	"prosr/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/markbates/pkger"
)

// Builder creates files for the specified language
type Builder struct {
	language string
	packages []string
	content  parser.IContentContext
}

// NewBuilder ctor for Builder
func NewBuilder(language string, packages []string, content parser.IContentContext) *Builder {
	b := new(Builder)
	b.language = language
	b.packages = packages
	b.content = content

	return b
}

// Build creates files by Ast
func (b *Builder) Build() (*string, string) {
	tmpl, fe := buildTemplate(b)
	out := executeParse(tmpl, b.content)
	out = strings.Trim(out, " \r\n")

	if *fe == "" {
		panic("No file extension in template set!")
	}

	return fe, out
}

func buildTemplate(b *Builder) (*template.Template, *string) {
	fe := ""
	tm := map[string]string{}

	fm := template.FuncMap{
		"registerFileExtension": func(extension string) string { fe = extension; return "" },
		"addType":               func(key string, value string) string { tm[key] = value; return "" },
		"type":                  func(key string) string { return resolveType(tm, key) },

		"capitalizeFirstLetter": capitalizeFirstLetter,
		"is":                    isOfType,
		"unifyCallDefinition":   unifyCallDefinition,
	}

	// pkger.Include("/templates/") is called in main.go
	f, err := pkger.Open("/templates/" + b.language + ".tmpl")
	if err != nil {
		panic(err)
	}

	r, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.
		New(b.language + ".tmpl").
		Funcs(fm).
		Parse(string(r))

	if err != nil {
		panic(err)
	}

	return tmpl, &fe
}

func executeParse(tmpl *template.Template, data interface{}) string {
	b := new(bytes.Buffer)
	err := tmpl.Execute(b, data)
	if err != nil {
		panic(err)
	}

	return b.String()
}

func capitalizeFirstLetter(v string) string {
	switch len(v) {
	case 0:
		return v
	case 1:
		return strings.ToUpper(string(v[0]))
	default:
		return strings.ToUpper(string(v[0])) + v[1:]
	}
}

var ruleToIndexMap = map[string]int{
	"content":           0,
	"syntaxDefinition":  1,
	"bodyDefinition":    2,
	"packageDefinition": 3,
	"hubDefinition":     4,
	"actionDefinition":  5,
	"callsDefinition":   6,
	"messageDefinition": 7,
	"fieldDefinition":   8,
	"fullIdent":         9,
	"typeIdent":         10,
}

func isOfType(ctx antlr.RuleContext, typeDefinition string) bool {
	indx, cntns := ruleToIndexMap[typeDefinition]
	if !cntns {
		return false
	}

	return ctx.GetRuleIndex() == indx
}

func unifyCallDefinition(ctx parser.HubDefinitionContext) []*parser.CallsDefinitionContext {
	m := map[string]*parser.CallsDefinitionContext{}
	for _, vl := range ctx.AllActionDefinition() {
		action := vl.(*parser.ActionDefinitionContext)
		if action.CallsDefinition() == nil {
			continue
		}

		calls := action.CallsDefinition().(*parser.CallsDefinitionContext)
		key := calls.IDENT().GetText()
		if _, ok := m[key]; ok {
			continue
		}

		m[key] = calls
	}

	for _, vl := range ctx.AllCallsDefinition() {
		calls := vl.(*parser.CallsDefinitionContext)
		key := calls.IDENT().GetText()
		if _, ok := m[key]; ok {
			continue
		}

		m[key] = calls
	}

	r := []*parser.CallsDefinitionContext{}
	for _, vl := range m {
		r = append(r, vl)
	}

	return r
}

func resolveType(options map[string]string, key string) string {
	if v, ok := options[key]; ok {
		return v
	}

	return key
}
