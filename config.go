package swolf

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	tagName  = "sw"
	tagKey   = "key"
	tagValue = "tenant"
)

// Config for database manager.
type Config struct {
	Driver         string
	Connection     string         // Connection string with user, password and connection address
	MasterDatabase string         // Name of the master database
	MasterTable    string         // Name of the master table in master database
	Template       string         // Database template, works only for postgresql
	MasterData     MasterResolver // Names of control columns
	Mapper         FMapper        // Mapping functions
}

// MasterResolver is used to configure key and tenant fields.
type MasterResolver interface {
	getKey() string
	getTenant() string
}

type masterDataResolver struct {
	data interface{}
}

type masterFieldResolver struct {
	key    string
	tenant string
}

// MasterFieldResolver sets key and dependent fields from string.
func MasterFieldResolver(k, t string) MasterResolver {
	return &masterFieldResolver{
		key:    k,
		tenant: t,
	}
}

// MasterDataResolver is used to configure key and tenant fields from given type.
// Given struct must contain tags:
// `sw:"key"` for key
// `sw:"tenant"` for tenant database
func MasterDataResolver(in interface{}) MasterResolver {
	return &masterDataResolver{
		data: in,
	}
}

func (mdr *masterDataResolver) getKey() string {
	t := reflect.TypeOf(mdr.data).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)
		if tag != "" {
			pairs := strings.Split(tag, ",")
			if len(pairs) != 2 {
				panic(fmt.Sprintf(panicTagInWrongFormat, tag))
			}

			if pairs[0] == tagKey {
				return pairs[1]
			} else if pairs[0] == tagValue {
			} else {
				panic(fmt.Sprintf(panicTagInWrongFormat, tag))
			}
		}
		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
	panic(panicTagNotFound)
}

func (mdr *masterDataResolver) getTenant() string {
	t := reflect.TypeOf(mdr.data).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)
		if tag != "" {
			pairs := strings.Split(tag, ",")
			if len(pairs) != 2 {
				panic(fmt.Sprintf(panicTagInWrongFormat, tag))
			}

			if pairs[0] == tagKey {
			} else if pairs[0] == tagValue {
				return pairs[1]
			} else {
				panic(fmt.Sprintf(panicTagInWrongFormat, tag))
			}
		}
		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
	panic(panicTagNotFound)
}

func (mfr *masterFieldResolver) getKey() string {
	return mfr.key
}

func (mfr *masterFieldResolver) getTenant() string {
	return mfr.tenant
}
