package builder

import "strings"

type Arg interface {
	Build() string
}

type SQLBuilder struct {
	builder strings.Builder
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

func (b *SQLBuilder) Insert(table string, what ...string) *SQLBuilder {
	b.builder.WriteString("INSERT INTO ")
	b.builder.WriteString(table)
	b.builder.WriteString("(")

	for i := 0; i < len(what)-1; i++ {
		b.builder.WriteString(what[i])
		b.builder.WriteString(", ")
	}
	b.builder.WriteString(what[len(what)-1])
	b.builder.WriteString(") ")

	return b
}

func (b *SQLBuilder) Values(args ...string) *SQLBuilder {
	b.builder.WriteString("VALUES (")

	for i := 0; i < len(args)-1; i++ {
		b.builder.WriteString(args[i])
		b.builder.WriteString(", ")
	}
	b.builder.WriteString(args[len(args)-1])
	b.builder.WriteString(") ")

	return b
}

func (b *SQLBuilder) Delete(table string) *SQLBuilder {
	b.builder.WriteString("DELETE FROM ")
	b.builder.WriteString(table)
	b.builder.WriteString(" ")

	return b
}

func (b *SQLBuilder) Create(dbname string) *SQLBuilder {
	b.builder.WriteString("CREATE DATABASE ")
	b.builder.WriteString(dbname)
	b.builder.WriteString(" ")

	return b
}

func (b *SQLBuilder) Template(template string) *SQLBuilder {
	b.builder.WriteString("TEMPLATE ")
	b.builder.WriteString(template)
	b.builder.WriteString(" ")

	return b
}

func (b *SQLBuilder) Drop(dbname string) *SQLBuilder {
	b.builder.WriteString("DROP DATABASE ")
	b.builder.WriteString(dbname)
	b.builder.WriteString(" ")

	return b
}

func (b *SQLBuilder) Build() string {
	return b.builder.String()
}

func NewBuilder() *SQLBuilder {
	return &SQLBuilder{}
}
