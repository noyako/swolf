package builder

import "strings"

type Arg interface {
	Build() string
}

type SQLBuilder struct {
	builder strings.Builder
}

type Eq struct {
	Key   string
	Value string
}

type Op struct {
	Operation string
}

func (e Eq) Build() string {
	return e.Key + "=" + e.Value + " "
}

func (o Op) Build() string {
	return o.Operation + " "
}

func (b *SQLBuilder) Select(what ...string) *SQLBuilder {
	b.builder.WriteString("SELECT ")

	for i := 0; i < len(what)-1; i++ {
		b.builder.WriteString(what[i])
		b.builder.WriteString(", ")
	}
	b.builder.WriteString(what[len(what)-1])
	b.builder.WriteString(" ")

	return b
}

func (b *SQLBuilder) From(what string) *SQLBuilder {
	b.builder.WriteString("FROM ")

	b.builder.WriteString(what)
	b.builder.WriteString(" ")

	return b
}

func (b *SQLBuilder) Where(args ...Arg) *SQLBuilder {
	b.builder.WriteString("WHERE ")

	for i := range args {
		b.builder.WriteString(args[i].Build())
	}

	return b
}
