package builder

type Equals struct {
	Key   string
	Value string
}

type Op struct {
	Operation string
}

func Eq(k, v string) *Equals {
	return &Equals{k, v}
}

func (e Equals) Build() string {
	return e.Key + "=" + e.Value + " "
}

func (o Op) Build() string {
	return o.Operation + " "
}
