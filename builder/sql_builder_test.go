package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		columns []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "No columns",
			want: "SELECT * ",
		},
		{
			name: "One column",
			args: args{[]string{
				"id",
			}},
			want: "SELECT id ",
		},
		{
			name: "Multiple columns",
			args: args{[]string{
				"id",
				"name",
				"date",
			}},
			want: "SELECT id, name, date ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := NewBuilder().Select(tt.args.columns...).Build()
			assert.Equal(tt.want, req)
		})
	}
}

func TestFrom(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		table string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal",
			args: args{"user"},
			want: "FROM user ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := NewBuilder().From(tt.args.table).Build()
			assert.Equal(tt.want, req)
		})
	}
}

func TestOp(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		op string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "And",
			args: args{And},
			want: "AND ",
		},
		{
			name: "Or",
			args: args{Or},
			want: "OR ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := Op(tt.args.op).Build()
			assert.Equal(tt.want, req)
		})
	}
}

func TestEq(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		l string
		r string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal",
			args: args{"key", "value"},
			want: "(key=value) ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := Eq(tt.args.l, tt.args.r).Build()
			assert.Equal(tt.want, req)
		})
	}
}

func TestWhere(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		conditions []Arg
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "No conditions",
			want: "WHERE ",
		},
		{
			name: "One conditions",
			args: args{[]Arg{
				Eq("key", "value"),
			}},
			want: "WHERE (key=value) ",
		},
		{
			name: "Multiple conditions",
			args: args{[]Arg{
				Eq("key1", "value1"),
				Op(And),
				Eq("key2", "value2"),
				Op(Or),
				Eq("key3", "value3"),
			}},
			want: "WHERE (key1=value1) AND (key2=value2) OR (key3=value3) ",
		},
		{
			name: "Nested conditions",
			args: args{[]Arg{
				Cl(
					Eq("a", "b"),
					Op(And),
					Eq("b", "c"),
				),
				Op(Or),
				Eq("c", "d"),
			}},
			want: "WHERE ((a=b) AND (b=c) ) OR (c=d) ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := NewBuilder().Where(tt.args.conditions...).Build()
			assert.Equal(tt.want, req)
		})
	}
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		table   string
		columns []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "No columns",
			args: args{
				table: "main_table",
			},
			want: "INSERT INTO main_table() ",
		},
		{
			name: "One column",
			args: args{
				"main_table",
				[]string{
					"id",
				}},
			want: "INSERT INTO main_table(id) ",
		},
		{
			name: "Multiple columns",
			args: args{
				"main_table",
				[]string{
					"id",
					"name",
					"date",
				}},
			want: "INSERT INTO main_table(id, name, date) ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := NewBuilder().Insert(tt.args.table, tt.args.columns...).Build()
			assert.Equal(tt.want, req)
		})
	}
}

func TestValues(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "No value",
			want: "VALUES () ",
		},
		{
			name: "One column",
			args: args{
				[]string{
					"id",
				}},
			want: "VALUES (id) ",
		},
		{
			name: "Multiple columns",
			args: args{
				[]string{
					"id",
					"name",
					"date",
				}},
			want: "VALUES (id, name, date) ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := NewBuilder().Values(tt.args.values...).Build()
			assert.Equal(tt.want, req)
		})
	}
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name string
		want string
	}{
		{
			name: "Normal",
			want: "DELETE ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := NewBuilder().Delete().Build()
			assert.Equal(tt.want, req)
		})
	}
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal",
			args: args{"my_database"},
			want: "CREATE DATABASE my_database ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := NewBuilder().Create(tt.args.name).Build()
			assert.Equal(tt.want, req)
		})
	}
}

func TestDrop(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal",
			args: args{"my_database"},
			want: "DROP DATABASE my_database ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := NewBuilder().Drop(tt.args.name).Build()
			assert.Equal(tt.want, req)
		})
	}
}

func TestTemplate(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal",
			args: args{"template1"},
			want: "TEMPLATE template1 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := NewBuilder().Template(tt.args.name).Build()
			assert.Equal(tt.want, req)
		})
	}
}
