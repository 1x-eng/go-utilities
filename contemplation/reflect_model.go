package contemplation

import (
	"reflect"
)


type DTOReflection struct {
	FieldName       interface{}
	FieldType       interface{}
	FieldValue      interface{}
	TagFieldName    interface{} 
}

type StructToReflect[T any] struct {
	AModel		T
	TagName		string
}

func ReflectStructFieldsByTag(model StructToReflect) ([]DTOReflection, error) {
	
	valueOfProperty := reflect.ValueOf(model.AModel)
	typeOfProperty := valueOfProperty.Type()
	
	reflectedStruct := []DTOReflection
	
	for i := 0; i < valueOfProperty.NumField(); i++ {
		reflectedStructField, ok := reflect.TypeOf(model.AModel).FieldByName(typeOfProperty.Field(i).Name)
		
		if !ok {
			panic("Unable to reflect. Field not found")
		}
		
		reflectedStruct = append(reflectedStruct, DTOReflection{
			FieldName: typeOfProperty.Field(i).Name,
			FieldType: typeOfProperty.Field(i).Type,
			FieldValue: valueOfProperty.Field(i).Interface(),
			TagFieldName: getStructTag(reflectedStructField, model.TagName)
		})
	}
	return reflectedStruct
}

func getStructTag(f reflect.StructField, tagName string) string {
	return string(f.Tag.Get(tagName))
}