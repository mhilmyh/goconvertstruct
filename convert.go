package goconvertstruct

import (
	"reflect"
)

var availableType = map[reflect.Kind]interface{}{
	reflect.Bool:       map[bool]interface{}{},
	reflect.Int:        map[int]interface{}{},
	reflect.Int8:       map[int8]interface{}{},
	reflect.Int16:      map[int16]interface{}{},
	reflect.Int32:      map[int32]interface{}{},
	reflect.Int64:      map[int64]interface{}{},
	reflect.Uint:       map[uint]interface{}{},
	reflect.Uint8:      map[uint8]interface{}{},
	reflect.Uint16:     map[uint16]interface{}{},
	reflect.Uint32:     map[uint32]interface{}{},
	reflect.Uint64:     map[uint64]interface{}{},
	reflect.Uintptr:    map[uintptr]interface{}{},
	reflect.Float32:    map[float32]interface{}{},
	reflect.Float64:    map[float64]interface{}{},
	reflect.Complex64:  map[complex64]interface{}{},
	reflect.Complex128: map[complex128]interface{}{},
	reflect.Interface:  map[interface{}]interface{}{},
	reflect.String:     map[string]interface{}{},
}

type ConverterOption struct {
	MapKeyFromTag    string
	FallBackEmptyTag bool
}

func defaultOption() *ConverterOption {
	return &ConverterOption{
		MapKeyFromTag:    "json",
		FallBackEmptyTag: true,
	}
}

func initMap(value reflect.Value) reflect.Value {
	var dummy interface{}
	if dummyType, ok := availableType[value.Type().Key().Kind()]; ok {
		dummy = dummyType
	}
	return reflect.MakeMap(reflect.TypeOf(dummy))
}

func Convert(target interface{}, option *ConverterOption) interface{} {
	if option == nil {
		option = defaultOption()
	}
	original := reflect.ValueOf(target)
	if target == nil || (original.Kind() == reflect.Ptr && original.IsNil()) {
		return nil
	}
	value := reflect.Indirect(original)
	switch value.Kind() {
	case reflect.Struct:
		result := make(map[string]interface{})
		typeOf := value.Type()
		for i := 0; i < typeOf.NumField(); i++ {
			field := typeOf.Field(i)
			name := field.Tag.Get(option.MapKeyFromTag)
			if name == "" && option.FallBackEmptyTag {
				name = field.Name
			}
			if !field.IsExported() {
				continue
			}
			result[name] = Convert(value.Field(i).Interface(), option)
		}
		return result
	case reflect.Array, reflect.Slice:
		if value.Kind() == reflect.Slice && value.IsNil() {
			return nil
		}
		result := make([]interface{}, value.Len())
		for i := 0; i < value.Len(); i++ {
			item := value.Index(i)
			result[i] = Convert(item.Interface(), option)
		}
		return result
	case reflect.Map:
		if value.IsNil() {
			return nil
		}
		result := initMap(value)
		for _, key := range value.MapKeys() {
			origin := value.MapIndex(key)
			converted := Convert(origin.Interface(), option)
			result.SetMapIndex(key, reflect.ValueOf(converted))
		}
		return result.Interface()
	case reflect.Ptr:
		if value.IsNil() {
			return nil
		}
		return Convert(value.Interface(), option)
	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return nil
	default:
		return value.Interface()
	}
}
