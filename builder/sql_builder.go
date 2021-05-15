// Package builder provides methods for creating
// simple SQL queries.
package builder

import "strings"

// Arg is an interface for any type supported by this package.
type Arg interface {
	Build() string
}

// SQLBuilder represents a starting builder for
// any SQL query.
type SQLBuilder struct {
	builder strings.Builder
}

// Select builds SELECT query with output columns
// as what argument.
func (b *SQLBuilder) Select(what ...string) *SQLBuilder {
	b.builder.WriteString("SELECT ")

	for i := 0; i < len(what)-1; i++ {
		b.builder.WriteString(what[i])
		b.builder.WriteString(", ")
	}
	if len(what) == 0 {
		b.builder.WriteString("*")
	} else {
		b.builder.WriteString(what[len(what)-1])
	}
	b.builder.WriteString(" ")

	return b
}

// From adds FROM statement to query.
func (b *SQLBuilder) From(what string) *SQLBuilder {
	b.builder.WriteString("FROM ")

	b.builder.WriteString(what)
	b.builder.WriteString(" ")

	return b
}

// Where adds WHERE statement to query.
// Equation or/and operators are specified as args parameter.
func (b *SQLBuilder) Where(args ...Arg) *SQLBuilder {
	b.builder.WriteString("WHERE ")

	for i := range args {
		b.builder.WriteString(args[i].Build())
	}

	return b
}

// Insert adds INSERT TO statement to query.
func (b *SQLBuilder) Insert(table string, what ...string) *SQLBuilder {
	b.builder.WriteString("INSERT INTO ")
	b.builder.WriteString(table)
	b.builder.WriteString("(")

	for i := 0; i < len(what)-1; i++ {
		b.builder.WriteString(what[i])
		b.builder.WriteString(", ")
	}
	if len(what) != 0 {
		b.builder.WriteString(what[len(what)-1])
	}

	b.builder.WriteString(") ")

	return b
}

// Values adds VALUES() statement to query.
func (b *SQLBuilder) Values(args ...string) *SQLBuilder {
	b.builder.WriteString("VALUES (")

	for i := 0; i < len(args)-1; i++ {
		b.builder.WriteString(args[i])
		b.builder.WriteString(", ")
	}
	if len(args) != 0 {
		b.builder.WriteString(args[len(args)-1])
	}
	b.builder.WriteString(") ")

	return b
}

// Delete adds DELETE statement to query.
func (b *SQLBuilder) Delete() *SQLBuilder {
	b.builder.WriteString("DELETE ")

	return b
}

// Create adds CREATE DATABASE statement to query.
func (b *SQLBuilder) Create(dbname string) *SQLBuilder {
	b.builder.WriteString("CREATE DATABASE ")
	b.builder.WriteString(dbname)
	b.builder.WriteString(" ")

	return b
}

// Template adds TEMPLATE statement to query.
func (b *SQLBuilder) Template(template string) *SQLBuilder {
	b.builder.WriteString("TEMPLATE ")
	b.builder.WriteString(template)
	b.builder.WriteString(" ")

	return b
}

// Drop adds DROP DATABASE statement to query.
func (b *SQLBuilder) Drop(dbname string) *SQLBuilder {
	b.builder.WriteString("DROP DATABASE ")
	b.builder.WriteString(dbname)
	b.builder.WriteString(" ")

	return b
}

// Build returns a string with constructed query.
func (b *SQLBuilder) Build() string {
	return b.builder.String()
}

// NewBuilder creates empty builder.
func NewBuilder() *SQLBuilder {
	return &SQLBuilder{}
}
