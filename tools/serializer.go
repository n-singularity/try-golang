package tools

import (
	"reflect"
)

//TODO: cara remove serializerStruct2
func Serialize(model, serializerStructAddress, serializerStruct interface{}) interface{} {
	array := reflect.ValueOf(model)

	if array.Kind() != reflect.Slice {
		return SingleSerialize(model, serializerStructAddress)
	}

	elemType := reflect.TypeOf(serializerStruct)
	result := reflect.MakeSlice(reflect.SliceOf(elemType), 0, 1000)
	for i := 0; i < array.Len(); i++ {
		var value = SingleSerialize(array.Index(i).Interface(), serializerStructAddress)

		result = reflect.Append(result, reflect.ValueOf(value))
	}

	return result.Interface()
}

func SingleSerialize(model, serializerStruct interface{}) interface{} {
	result := serializerStruct
	resultReflectSerializer := reflect.ValueOf(result).Elem()
	reflectModel := reflect.TypeOf(model)
	modelValue := reflect.ValueOf(model)

	fieldNames := make(map[string]struct{})
	collectFieldNames(reflect.TypeOf(serializerStruct), fieldNames)

	for fieldName := range fieldNames {
		_, fieldExist := reflectModel.FieldByName(fieldName)
		_, methodExist := reflectModel.MethodByName(fieldName)

		if fieldExist {
			resultReflectSerializer.FieldByName(fieldName).Set(modelValue.FieldByName(fieldName))
		} else if methodExist {
			meth := modelValue.MethodByName(fieldName)
			res := meth.Call(nil)
			ret := res[0]
			resultReflectSerializer.FieldByName(fieldName).Set(ret)
		}
	}

	value := reflect.ValueOf(result)

	if value.Type().Kind() == reflect.Ptr {
		ptr := value
		value = ptr.Elem() // acquire value referenced by pointer
	} else {
		ptr := reflect.New(reflect.TypeOf(result)) // create new pointer
		temp := ptr.Elem()                         // create variable to value of pointer
		temp.Set(value)                            // set value of variable to our passed in value
	}

	return value.Interface()
}

func collectFieldNames(t reflect.Type, m map[string]struct{}) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		m[sf.Name] = struct{}{}
		if sf.Anonymous {
			collectFieldNames(sf.Type, m)
		}
	}
}
