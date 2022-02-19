package contemplation

import (
	"reflect"
)

type StructToReflect struct {
	AModel  interface{}
	TagName string
}

type reflectedModel struct {
	FieldName    string
	FieldType    interface{}
	FieldValue   interface{}
	TagFieldName string
}

func ReflectStructFieldsByTag(model StructToReflect) ([]reflectedModel, error) {

	valueOfProperty := reflect.ValueOf(model.AModel)
	typeOfProperty := valueOfProperty.Type()

	reflectedStruct := []reflectedModel{}

	for i := 0; i < valueOfProperty.NumField(); i++ {
		reflectedStructField, ok := reflect.TypeOf(model.AModel).FieldByName(typeOfProperty.Field(i).Name)

		if !ok {
			panic("Unable to reflect. Field not found")
		}

		reflectedStruct = append(reflectedStruct, reflectedModel{
			FieldName:    typeOfProperty.Field(i).Name,
			FieldType:    typeOfProperty.Field(i).Type,
			FieldValue:   valueOfProperty.Field(i).Interface(),
			TagFieldName: getStructTag(reflectedStructField, model.TagName),
		})
	}
	return reflectedStruct, nil
}

func getStructTag(f reflect.StructField, tagName string) string {
	return string(f.Tag.Get(tagName))
}
