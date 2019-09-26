package bindings

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)

/* **********************************************************
 * Copyright 2018 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

/**
 * Given a BindingType, builds its corresponding StructValue instance.
 */
type StructValueBuilder struct {
	structType    StructType
	typeConverter *TypeConverter
	fields        map[string]interface{}
}

func NewStructValueBuilder(structType StructType, converter *TypeConverter) *StructValueBuilder {
	return &StructValueBuilder{structType: structType, typeConverter: converter, fields: map[string]interface{}{}}
}

func (s *StructValueBuilder) AddStructField(name string, value interface{}) {
	s.fields[name] = value
}

/**
 * Returns StructValue after converting go native values into their corresponding DataValue type.
 */
func (s *StructValueBuilder) GetStructValue() (*data.StructValue, []error) {
	sv := data.NewStructValue(s.structType.Name(), nil)
	for _, fieldName := range s.structType.FieldNames() {
		fieldBindingType := s.structType.Field(fieldName)
		fieldGoValue := s.fields[s.structType.canonicalFieldMap[fieldName]]
		fieldDataValue, err := s.typeConverter.ConvertToVapi(fieldGoValue, fieldBindingType)
		if err != nil {
			return nil, err
		}
		sv.SetField(fieldName, fieldDataValue)
	}
	msgs := s.structType.Validate(sv)
	if msgs != nil {
		return nil, msgs
	}
	return sv, nil
}
