package builder

import "strings"

// Equals represents equals statement.
// I.e key=value
type Equals struct {
	Key   string
	Value string
}

// Opearion represents operaion statement.
type Opearion struct {
	Operation string
}

// Closer represents set of statements in parentheses.
type Closer struct {
	Value []Arg
}

// Eq constructs Equals object.
func Eq(k, v string) *Equals {
	return &Equals{k, v}
}

// Cl constructs Closer object.
func Cl(c ...Arg) *Closer {
	return &Closer{c}
}

// Op constructs Opearion object.
func Op(o string) *Opearion {
	return &Opearion{o}
}

// Build returns a string with constructed query.
func (e Equals) Build() string {
	return "(" + e.Key + "=" + e.Value + ") "
}

// Build returns a string with constructed query.
func (o Opearion) Build() string {
	return o.Operation + " "
}

// Build returns a string with constructed query.
func (c Closer) Build() string {
	b := strings.Builder{}
	b.WriteString("(")

	for arg := range c.Value {
		b.WriteString(c.Value[arg].Build())
	}

	b.WriteString(") ")
	return b.String()
}
