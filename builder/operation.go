package builder

import "strings"

type Equals struct {
	Key   string
	Value string
}

type Opearion struct {
	Operation string
}

type Closer struct {
	Value []Arg
}

func Eq(k, v string) *Equals {
	return &Equals{k, v}
}

func Cl(c ...Arg) *Closer {
	return &Closer{c}
}

func Op(o string) *Opearion {
	return &Opearion{o}
}

func (e Equals) Build() string {
	return "(" + e.Key + "=" + e.Value + ") "
}

func (o Opearion) Build() string {
	return o.Operation + " "
}

func (c Closer) Build() string {
	b := strings.Builder{}
	b.WriteString("(")

	for arg := range c.Value {
		b.WriteString(c.Value[arg].Build())
	}

	b.WriteString(") ")
	return b.String()
}
