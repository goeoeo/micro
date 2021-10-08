package config

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/camelcase"
	"github.com/fatih/structs"
)

// EnvironmentLoader satisifies the loader interface. It loads the
// configuration from the environment variables in the form of
// STRUCTNAME_FIELDNAME.
type EnvironmentLoader struct {
	// Prefix prepends given string to every environment variable
	// {STRUCTNAME}_FIELDNAME will be {PREFIX}_FIELDNAME
	Prefix string

	// CamelCase adds a separator for field names in camelcase form. A
	// fieldname of "AccessKey" would generate a environment name of
	// "STRUCTNAME_ACCESSKEY". If CamelCase is enabled, the environment name
	// will be generated in the form of "STRUCTNAME_ACCESS_KEY"
	CamelCase bool
}

func (e *EnvironmentLoader) getPrefix(s *structs.Struct) string {
	if e.Prefix != "" {
		return e.Prefix
	}

	return s.Name()
}

// Load loads the source into the config defined by struct s
func (e *EnvironmentLoader) Load(s interface{}) error {
	strct := structs.New(s)
	strctMap := strct.Map()
	prefix := e.getPrefix(strct)

	for key, val := range strctMap {
		field := strct.Field(key)

		if err := e.processField(prefix, field, key, val); err != nil {
			return err
		}
	}

	return nil
}

// processField gets leading name for the env variable and combines the current
// field's name and generates environment variable names recursively
func (e *EnvironmentLoader) processField(prefix string, field *structs.Field, name string, strctMap interface{}) error {
	fieldName := e.generateFieldName(prefix, name)

	switch strctMap.(type) {
	case map[string]interface{}:
		for key, val := range strctMap.(map[string]interface{}) {
			field := field.Field(key)

			if err := e.processField(fieldName, field, key, val); err != nil {
				return err
			}
		}
	default:
		v := os.Getenv(fieldName)
		if v == "" {
			return nil
		}

		if err := fieldSet(field, v); err != nil {
			return err
		}
	}

	return nil
}

// PrintEnvs prints the generated environment variables to the std out.
func (e *EnvironmentLoader) PrintEnvs(s interface{}) {
	strct := structs.New(s)
	strctMap := strct.Map()
	prefix := e.getPrefix(strct)

	keys := make([]string, 0, len(strctMap))
	for key := range strctMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		field := strct.Field(key)
		e.printField(prefix, field, key, strctMap[key])
	}
}

// printField prints the field of the config struct for the flag.Usage
func (e *EnvironmentLoader) printField(prefix string, field *structs.Field, name string, strctMap interface{}) {
	fieldName := e.generateFieldName(prefix, name)

	switch strctMap.(type) {
	case map[string]interface{}:
		smap := strctMap.(map[string]interface{})
		keys := make([]string, 0, len(smap))
		for key := range smap {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			field := field.Field(key)
			e.printField(fieldName, field, key, smap[key])
		}
	default:
		fmt.Println("  ", fieldName)
	}
}

// generateFieldName generates the field name combined with the prefix and the
// struct's field name
func (e *EnvironmentLoader) generateFieldName(prefix string, name string) string {
	fieldName := strings.ToUpper(name)
	if e.CamelCase {
		fieldName = strings.ToUpper(strings.Join(camelcase.Split(name), "_"))
	}

	return strings.ToUpper(prefix) + "_" + fieldName
}

// fieldSet sets field value from the given string value. It converts the
// string value in a sane way and is usefulf or environment variables or flags
// which are by nature in string types.
func fieldSet(field *structs.Field, v string) error {
	switch f := field.Value().(type) {
	case flag.Value:
		if v := reflect.ValueOf(field.Value()); v.IsNil() {
			typ := v.Type()
			if typ.Kind() == reflect.Ptr {
				typ = typ.Elem()
			}

			if err := field.Set(reflect.New(typ).Interface()); err != nil {
				return err
			}

			f = field.Value().(flag.Value)
		}

		return f.Set(v)
	}

	// TODO: add support for other types
	switch field.Kind() {
	case reflect.Bool:
		val, err := strconv.ParseBool(v)
		if err != nil {
			return err
		}

		if err := field.Set(val); err != nil {
			return err
		}
	case reflect.Int:
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}

		if err := field.Set(i); err != nil {
			return err
		}
	case reflect.String:
		if err := field.Set(v); err != nil {
			return err
		}
	case reflect.Slice:
		switch t := field.Value().(type) {
		case []string:
			if err := field.Set(strings.Split(v, ",")); err != nil {
				return err
			}
		case []int:
			var list []int
			for _, in := range strings.Split(v, ",") {
				i, err := strconv.Atoi(in)
				if err != nil {
					return err
				}

				list = append(list, i)
			}

			if err := field.Set(list); err != nil {
				return err
			}
		default:
			return fmt.Errorf("multiconfig: field '%s' of type slice is unsupported: %s (%T)",
				field.Name(), field.Kind(), t)
		}
	case reflect.Float64:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}

		if err := field.Set(f); err != nil {
			return err
		}
	case reflect.Int64:
		switch t := field.Value().(type) {
		case time.Duration:
			d, err := time.ParseDuration(v)
			if err != nil {
				return err
			}

			if err := field.Set(d); err != nil {
				return err
			}
		case int64:
			p, err := strconv.ParseInt(v, 10, 0)
			if err != nil {
				return err
			}

			if err := field.Set(p); err != nil {
				return err
			}
		default:
			return fmt.Errorf("multiconfig: field '%s' of type int64 is unsupported: %s (%T)",
				field.Name(), field.Kind(), t)
		}

	default:
		return fmt.Errorf("multiconfig: field '%s' has unsupported type: %s", field.Name(), field.Kind())
	}

	return nil
}
