package swolf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMasterDataResolver(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		data interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantKey   string
		wantValue string
	}{
		{
			name: "No columns",
			args: args{
				data: &struct {
				}{},
			},
		},
		{
			name: "Key column",
			args: args{
				data: &struct {
					k int `sw:"key,kw"`
				}{},
			},
			wantKey: "kw",
		},
		{
			name: "Value column",
			args: args{
				data: &struct {
					v int `sw:"tenant,tw"`
				}{},
			},
			wantValue: "tw",
		},
		{
			name: "All columns",
			args: args{
				data: &struct {
					k int `sw:"key,kw"`
					v int `sw:"tenant,tw"`
				}{},
			},
			wantKey:   "kw",
			wantValue: "tw",
		},
		{
			name: "Bad tags",
			args: args{
				data: &struct {
					k int `sw:"key1,kw"`
					v int `sw:"tenant1,tw"`
				}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolver := MasterDataResolver(tt.args.data)
			if tt.wantKey == "" {
				assert.Panics(func() { resolver.getKey() })
			} else {
				assert.Equal(tt.wantKey, resolver.getKey())
			}
			if tt.wantValue == "" {
				assert.Panics(func() { resolver.getTenant() })
			} else {
				assert.Equal(tt.wantValue, resolver.getTenant())
			}
		})
	}
}
