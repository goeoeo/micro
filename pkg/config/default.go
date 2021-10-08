package config

import (
	"reflect"

	"github.com/fatih/structs"
)

//载入对象的默认值
type Default struct {
	DefaultTagName string
}

func NewDefault() *Default {
	return &Default{
		DefaultTagName: "default",
	}
}

func (d *Default) Load(s interface{}) (err error) {
	if s == nil {
		return
	}

	for _, field := range structs.Fields(s) {
		if err = d.processTagField(field); err != nil {
			return err
		}
	}

	return nil
}

func (d *Default) processTagField(field *structs.Field) error {
	switch field.Kind() {
	case reflect.Struct:
		for _, f := range field.Fields() {
			if err := d.processTagField(f); err != nil {
				return err
			}
		}
	default:
		defaultVal := field.Tag(d.DefaultTagName)
		if defaultVal == "" {
			return nil
		}

		err := fieldSet(field, defaultVal)
		if err != nil {
			return err
		}
	}

	return nil
}
