package compiler

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

// Builder creates files for the specified language
type Builder struct {
	language string
	Options  map[string]string
	Ast      []Node
}

// NewBuilder ctor for Builder
func NewBuilder(language string, ast *Ast, Options map[string]string) *Builder {
	b := new(Builder)
	b.language = language
	b.Options = Options
	b.Ast = ast.Nodes

	return b
}

// Build creates files by Ast
func (b *Builder) Build() (*string, string) {
	tmpl, fe := buildTemplate(b)
	out := executeParse(tmpl, b)
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
		"type":                  func(key string) string { return resolveOption(tm, key) },

		"capitalizeFirstLetter": capitalizeFirstLetter,
		"resolveOption":         resolveOption,
		"unifyReturnings":       unifyReturnings,
	}

	tmpl, err := template.
		New(b.language + ".tmpl").
		Funcs(fm).
		ParseFiles("compiler/" + b.language + ".tmpl")

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

func resolveOption(Options map[string]string, key string) string {
	if v, ok := Options[key]; ok {
		return v
	}

	fmt.Println("Options key " + key + " was not found. Returned key instead.")

	return key
}

func unifyReturnings(nodes []Node) []*Returning {
	ns := map[string]*Returning{}
	for _, n := range nodes {
		re := resolveReturning(n)
		if re == nil {
			continue
		}

		ns[re.String()] = re
	}

	r := []*Returning{}
	for _, n := range ns {
		r = append(r, n)
	}

	return r
}

func resolveReturning(node Node) *Returning {
	switch n := node.(type) {
	case *Sending:
		if n.Calls == nil {
			return nil
		}

		return resolveReturning(n.Calls)
	case *Returning:
		return n
	default:
		return nil
	}
}