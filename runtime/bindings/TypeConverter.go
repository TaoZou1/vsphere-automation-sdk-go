/* Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package bindings

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/l10n"
	"reflect"
)

type ConverterMode string

var REST ConverterMode = "REST"
var JSONRPC ConverterMode = "JSONRPC"

// TypeConverter converts between Golang Native data model and API runtime data model.
type TypeConverter struct {
	// mode controls the way native types are rendered to DataValue. Most importantly if Maps will be represented as
	// struct or list
	options *typeConverterOptions
}

type typeConverterOptions struct {
	mode ConverterMode
}

type TypeConverterOption func(*typeConverterOptions)

func NewTypeConverter(options ...TypeConverterOption) *TypeConverter {
	doptions := &typeConverterOptions{
		mode: JSONRPC,
	}

	for _, o := range options {
		o(doptions)
	}

	return &TypeConverter{options: doptions}
}

func InRestMode() TypeConverterOption {
	return func(opts *typeConverterOptions) {
		opts.mode = REST
	}
}

// ConvertToGolang converts vapiValue which is an API runtime representation to its equivalent golang native representation
// with the help of bindingType
func (t *TypeConverter) ConvertToGolang(vapiValue data.DataValue, bindingType BindingType) (interface{}, []error) {
	if bindingType == nil {
		return nil, []error{l10n.NewRuntimeErrorNoParam("vapi.bindings.typeconverter.nil.type")}
	}
	var nativeConverter = NewDataValueToNativeConverter(vapiValue, t)
	err := nativeConverter.visit(bindingType)
	if err != nil {
		return nil, err
	}
	return nativeConverter.OutputValue(), nil
}

// ConvertToVapi converts golangValue which is native golang value to its equivalent api runtime representation
// with the help of bindingType.
func (t *TypeConverter) ConvertToVapi(golangValue interface{}, bindingType BindingType) (data.DataValue, []error) {
	if bindingType == nil {
		return nil, []error{l10n.NewRuntimeErrorNoParam("vapi.bindings.typeconverter.nil.type")}
	}
	if golangValue != nil {
		if reflect.TypeOf(golangValue).Kind() == reflect.Ptr && reflect.ValueOf(golangValue).IsNil() {
			golangValue = nil
		} else if reflect.TypeOf(golangValue).Kind() == reflect.Slice && reflect.ValueOf(golangValue).IsNil() {
			golangValue = nil
		} else if reflect.TypeOf(golangValue).Kind() == reflect.Map && reflect.ValueOf(golangValue).IsNil() {
			golangValue = nil
		}
	}

	if t.options.mode == REST {
		visitor := NewGolangToRestDataValueVisitor(golangValue)
		err := visitor.visit(bindingType)
		if err != nil {
			return nil, err
		}
		return visitor.OutputValue(), nil
	} else {
		visitor := NewGolangToVapiDataValueVisitor(golangValue)
		err := visitor.visit(bindingType)
		if err != nil {
			return nil, err
		}
		return visitor.OutputValue(), nil
	}
}

// ConvertToDataDefinition outputs DataDefinition representation of bindingType.
func (t *TypeConverter) ConvertToDataDefinition(bindingType BindingType) (data.DataDefinition, []error) {
	if bindingType == nil {
		return nil, []error{l10n.NewRuntimeErrorNoParam("vapi.bindings.typeconverter.nil.type")}
	}
	referenceResolver := data.NewReferenceResolver()
	seenStructures := map[string]bool{}
	var visitor = NewBindingTypeToDataDefinitionConverter(bindingType, referenceResolver, seenStructures)
	err := visitor.convert(bindingType)
	if err != nil {
		return nil, err
	}
	err = referenceResolver.ResolveReferences()
	if err != nil {
		return nil, err
	}
	return visitor.OutputValue(), nil
}

// Deprecated: SetPermissive does nothing and is left to keep code compatiblity
func (t *TypeConverter) SetPermissive(permissive bool) {
	// Deprecated
}

// SetMode controls the way native types are rendered to DataValue. Most importantly if Maps will be represented as
// struct or list
func (t *TypeConverter) SetMode(mode ConverterMode) {
	t.options.mode = mode
}

func (t *TypeConverter) Mode() ConverterMode {
	return t.options.mode
}
