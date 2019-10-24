/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Customization.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package guest

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/guest"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"time"
)

// The ``SetSpec`` class contains specification information that has to be applied to a virtual machine.
type CustomizationSetSpec struct {
    // The name of the customization specification that has be retrieved from the virtual center inventory and applied for the virtual machine. Either one of ``name`` or ``spec`` or none of them should be specified.
	Name *string
    // The customization specification that has to be applied for the virtual machine. Either one of ``name`` or ``spec`` or none of them should be specified.
	Spec *guest.CustomizationSpec
}

// The ``Info`` class contains the status of a customization operation applied to a virtual machine. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CustomizationInfo struct {
    // The status of the customization operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Status CustomizationInfoStatus
    // Description of the error if the CustomizationInfo#status of customization operation is CustomizationInfoStatus#Status_FAILED. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Error_ *string
    // Time when the customization process has started inside the guest operating system. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	StartTime *time.Time
    // Time when the customization process has completed inside the guest operating system. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	EndTime *time.Time
}

// The ``Status`` enumeration class defines the status values that can be reported for the customization operation. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CustomizationInfoStatus string

const (
    // The customization process has not yet started inside the guest operating system. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CustomizationInfoStatus_PENDING CustomizationInfoStatus = "PENDING"
    // The customization process is currently running inside the guest operating system. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CustomizationInfoStatus_RUNNING CustomizationInfoStatus = "RUNNING"
    // The customization process has completed successfully inside the guest operating system. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CustomizationInfoStatus_SUCCEEDED CustomizationInfoStatus = "SUCCEEDED"
    // The customizatio process has failed inside the guest operating system. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CustomizationInfoStatus_FAILED CustomizationInfoStatus = "FAILED"
)

func (s CustomizationInfoStatus) CustomizationInfoStatus() bool {
	switch s {
	case CustomizationInfoStatus_PENDING:
		return true
	case CustomizationInfoStatus_RUNNING:
		return true
	case CustomizationInfoStatus_SUCCEEDED:
		return true
	case CustomizationInfoStatus_FAILED:
		return true
	default:
		return false
	}
}



func customizationSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(CustomizationSetSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func customizationSetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(CustomizationSetSpecBindingType)
	paramsTypeMap["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	paramsTypeMap["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	pathParams["vm"] = "vm"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"spec",
		"PUT",
		"/vcenter/vm/{vm}/guest/customization",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CustomizationInfoBindingType)
}

func customizationGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	paramsTypeMap["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	pathParams["vm"] = "vm"
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
		"/vcenter/vm/{vm}/guest/customization",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func CustomizationSetSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(guest.CustomizationSpecBindingType))
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.guest.customization.set_spec", fields, reflect.TypeOf(CustomizationSetSpec{}), fieldNameMap, validators)
}

func CustomizationInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest.customization.info.status", reflect.TypeOf(CustomizationInfoStatus(CustomizationInfoStatus_PENDING)))
	fieldNameMap["status"] = "Status"
	fields["error"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["error"] = "Error_"
	fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("status",
	    map[string][]bindings.FieldData{
	        "FAILED": []bindings.FieldData{
	             bindings.NewFieldData("error", false),
	             bindings.NewFieldData("start_time", false),
	             bindings.NewFieldData("end_time", false),
	        },
	        "RUNNING": []bindings.FieldData{
	             bindings.NewFieldData("start_time", false),
	        },
	        "SUCCEEDED": []bindings.FieldData{
	             bindings.NewFieldData("start_time", false),
	             bindings.NewFieldData("end_time", false),
	        },
	        "PENDING": []bindings.FieldData{},
	    },
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.guest.customization.info", fields, reflect.TypeOf(CustomizationInfo{}), fieldNameMap, validators)
}

