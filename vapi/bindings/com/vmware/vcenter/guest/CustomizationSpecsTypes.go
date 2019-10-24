/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: CustomizationSpecs.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package guest

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"time"
)

// The resource type for a vCenter guest customization specification
const CustomizationSpecs_RESOURCE_TYPE = "com.vmware.vcenter.guest.CustomizationSpec"


// The ``OsType`` enumeration class defines the types of guest operating systems for which guest customization is supported.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CustomizationSpecsOsType string

const (
    // A customization specification for a Windows guest operating system
	CustomizationSpecsOsType_WINDOWS CustomizationSpecsOsType = "WINDOWS"
    // A customization specification for a Linux guest operating system
	CustomizationSpecsOsType_LINUX CustomizationSpecsOsType = "LINUX"
)

func (o CustomizationSpecsOsType) CustomizationSpecsOsType() bool {
	switch o {
	case CustomizationSpecsOsType_WINDOWS:
		return true
	case CustomizationSpecsOsType_LINUX:
		return true
	default:
		return false
	}
}


// The ``Format`` enumeration class specifies the formats the customization specification can be exported to.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CustomizationSpecsFormat string

const (
    // JSON format.
	CustomizationSpecsFormat_JSON CustomizationSpecsFormat = "JSON"
    // XML format.
	CustomizationSpecsFormat_XML CustomizationSpecsFormat = "XML"
)

func (f CustomizationSpecsFormat) CustomizationSpecsFormat() bool {
	switch f {
	case CustomizationSpecsFormat_JSON:
		return true
	case CustomizationSpecsFormat_XML:
		return true
	default:
		return false
	}
}


// The ``Metadata`` class contains metadata i.e. name and description related to a customization specification.
type CustomizationSpecsMetadata struct {
    // Description of the specification.
	Description string
    // Name of the specification.
	Name string
}

// The ``CreateSpec`` class contains specification information and specification object that can be passed to the CustomizationSpecs#create method.
type CustomizationSpecsCreateSpec struct {
    // The specification object.
	Spec CustomizationSpec
    // Description of the specification.
	Description string
    // Name of the specification.
	Name string
}

// The ``Spec`` class contains the specification information and specification object. This is passed to the CustomizationSpecs#set method.
type CustomizationSpecsSpec struct {
    // The fingerprint is a unique identifier for a given version of the configuration. Each change to the configuration will update this value. A client cannot change this value. If specified when updating a specification, the changes will only be applied if the current fingerprint matches the specified fingerprint. This field can be used to guard against updates that has happened between the specification content was read and until it is applied.
	Fingerprint string
    // The specification object.
	Spec CustomizationSpec
    // Description of the specification.
	Description string
    // Name of the specification.
	Name string
}

// The ``Info`` class describes a guest customization specification and the timestamp when it was last modified. This is returned by the CustomizationSpecs#get method.
type CustomizationSpecsInfo struct {
    // Time when the specification was last modified.
	LastModified time.Time
    // The Spec object including specification and metadata information.
	Spec CustomizationSpecsSpec
}

// The ``FilterSpec`` class contains properties used to filter the results when listing guest customization specifications (see CustomizationSpecs#list). If multiple properties are specified, only guest customization specifications matching all of the properties match the filter.
type CustomizationSpecsFilterSpec struct {
    // Names that guest customization specifications must have to match the filter (see CustomizationSpecsSummary#name).
	Names map[string]bool
    // Guest operating system type that guest customization specifications must have to match the filter (see CustomizationSpecsSummary#osType).
	OSType *CustomizationSpecsOsType
}

// The ``Summary`` class contains commonly used information about a guest customization specification.
type CustomizationSpecsSummary struct {
    // Name of the guest customization specification.
	Name string
    // Description of the guest customization specification.
	Description string
    // Guest operating system type for which that this guest customization specification applies.
	OSType CustomizationSpecsOsType
    // Date and tme when this guest customization specification was last modified.
	LastModified time.Time
}



func customizationSpecsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(CustomizationSpecsFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CustomizationSpecsSummaryBindingType), reflect.TypeOf([]CustomizationSpecsSummary{}))
}

func customizationSpecsListRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["filter.OS_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.os_type", reflect.TypeOf(CustomizationSpecsOsType(CustomizationSpecsOsType_WINDOWS))))
	paramsTypeMap["filter.names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, ""), reflect.TypeOf(map[string]bool{})))
	queryParams["filter.names"] = "names"
	queryParams["filter.OS_type"] = "os_type"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"GET",
		"/vcenter/guest/customization-specs",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationSpecsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(CustomizationSpecsCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
}

func customizationSpecsCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(CustomizationSpecsCreateSpecBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"spec",
		"POST",
		"/vcenter/guest/customization-specs",
		resultHeaders,
		201,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationSpecsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CustomizationSpecsInfoBindingType)
}

func customizationSpecsGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	pathParams["name"] = "name"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"GET",
		"/vcenter/guest/customization-specs/{name}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationSpecsSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fields["spec"] = bindings.NewReferenceType(CustomizationSpecsSpecBindingType)
	fieldNameMap["name"] = "Name"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func customizationSpecsSetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(CustomizationSpecsSpecBindingType)
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	pathParams["name"] = "name"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"spec",
		"POST",
		"/vcenter/guest/customization-specs/{name}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationSpecsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func customizationSpecsDeleteRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	pathParams["name"] = "name"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"DELETE",
		"/vcenter/guest/customization-specs/{name}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationSpecsExportInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fields["format"] = bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.format", reflect.TypeOf(CustomizationSpecsFormat(CustomizationSpecsFormat_JSON)))
	fieldNameMap["name"] = "Name"
	fieldNameMap["format"] = "Format"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsExportOutputType() bindings.BindingType {
	return bindings.NewStringType()
}

func customizationSpecsExportRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["format"] = bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.format", reflect.TypeOf(CustomizationSpecsFormat(CustomizationSpecsFormat_JSON)))
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	pathParams["name"] = "name"
	queryParams["format"] = "format"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"POST",
		"/vcenter/guest/customization-specs/{name}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationSpecsImportSpecificationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customization_spec"] = bindings.NewStringType()
	fieldNameMap["customization_spec"] = "CustomizationSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsImportSpecificationOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CustomizationSpecsCreateSpecBindingType)
}

func customizationSpecsImportSpecificationRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["customization_spec"] = bindings.NewStringType()
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"customization_spec",
		"POST",
		"/vcenter/guest/customization-specs",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func CustomizationSpecsMetadataBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.metadata", fields, reflect.TypeOf(CustomizationSpecsMetadata{}), fieldNameMap, validators)
}

func CustomizationSpecsCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(CustomizationSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.create_spec", fields, reflect.TypeOf(CustomizationSpecsCreateSpec{}), fieldNameMap, validators)
}

func CustomizationSpecsSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["fingerprint"] = bindings.NewStringType()
	fieldNameMap["fingerprint"] = "Fingerprint"
	fields["spec"] = bindings.NewReferenceType(CustomizationSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.spec", fields, reflect.TypeOf(CustomizationSpecsSpec{}), fieldNameMap, validators)
}

func CustomizationSpecsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["last_modified"] = bindings.NewDateTimeType()
	fieldNameMap["last_modified"] = "LastModified"
	fields["spec"] = bindings.NewReferenceType(CustomizationSpecsSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.info", fields, reflect.TypeOf(CustomizationSpecsInfo{}), fieldNameMap, validators)
}

func CustomizationSpecsFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["names"] = "Names"
	fields["OS_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.os_type", reflect.TypeOf(CustomizationSpecsOsType(CustomizationSpecsOsType_WINDOWS))))
	fieldNameMap["OS_type"] = "OSType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.filter_spec", fields, reflect.TypeOf(CustomizationSpecsFilterSpec{}), fieldNameMap, validators)
}

func CustomizationSpecsSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["OS_type"] = bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.os_type", reflect.TypeOf(CustomizationSpecsOsType(CustomizationSpecsOsType_WINDOWS)))
	fieldNameMap["OS_type"] = "OSType"
	fields["last_modified"] = bindings.NewDateTimeType()
	fieldNameMap["last_modified"] = "LastModified"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.summary", fields, reflect.TypeOf(CustomizationSpecsSummary{}), fieldNameMap, validators)
}

