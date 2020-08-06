package compiler

// Builder creates files for the specified language
type Builder struct {
	language string
	options  map[string]string
}

// NewBuilder ctor for Builder
func NewBuilder(language string, options map[string]string) *Builder {
	b := new(Builder)
	b.language = language
	b.options = options

	return b
}

// Build creates files by Ast
func (b *Builder) Build(ast *Ast) {

}
