/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Sata.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package adapter

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for the virtual SATA adapter device.
const Sata_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.SataAdapter"


// The ``Type`` enumeration class defines the valid emulation types for a virtual SATA adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Sata_Type string

const (
    // AHCI host bus adapter.
     Sata_Type_AHCI Sata_Type = "AHCI"
)

func (t Sata_Type) Sata_Type() bool {
    switch t {
        case Sata_Type_AHCI:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information about a virtual SATA adapter.
 type SataInfo struct {
    // Device label.
    Label string
    // Adapter type.
    Type_ Sata_Type
    // SATA bus number.
    Bus int64
    // Address of the SATA adapter on the PCI bus.
    PciSlotNumber *int64
}



func (SataInfo SataInfo) Error() string {
    return "com.vmware.vcenter.vm.hardware.adapter.info"
}



// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual SATA adapter.
 type SataCreateSpec struct {
    // Adapter type.
    Type_ *Sata_Type
    // SATA bus number.
    Bus *int64
    // Address of the SATA adapter on the PCI bus.
    PciSlotNumber *int64
}



func (SataCreateSpec SataCreateSpec) Error() string {
    return "com.vmware.vcenter.vm.hardware.adapter.create_spec"
}



// The ``Summary`` class contains commonly used information about a Virtual SATA adapter.
 type SataSummary struct {
    // Identifier of the virtual SATA adapter.
    Adapter string
}



func (SataSummary SataSummary) Error() string {
    return "com.vmware.vcenter.vm.hardware.adapter.summary"
}






func sataListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sataListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SataSummaryBindingType), reflect.TypeOf([]SataSummary{}))
}

func sataListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func sataGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["adapter"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SataAdapter"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["adapter"] = "Adapter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sataGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SataInfoBindingType)
}

func sataGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func sataCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewReferenceType(SataCreateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sataCreateOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SataAdapter"}, "")
}

func sataCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"UnableToAllocateResource": 500,"ResourceInUse": 400,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}


func sataDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["adapter"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SataAdapter"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["adapter"] = "Adapter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sataDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func sataDeleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func SataInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.adapter.sata.type", reflect.TypeOf(Sata_Type(Sata_Type_AHCI)))
    fieldNameMap["type"] = "Type_"
    fields["bus"] = bindings.NewIntegerType()
    fieldNameMap["bus"] = "Bus"
    fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pci_slot_number"] = "PciSlotNumber"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.sata.info",fields, reflect.TypeOf(SataInfo{}), fieldNameMap, validators)
}

func SataCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.adapter.sata.type", reflect.TypeOf(Sata_Type(Sata_Type_AHCI))))
    fieldNameMap["type"] = "Type_"
    fields["bus"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["bus"] = "Bus"
    fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pci_slot_number"] = "PciSlotNumber"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.sata.create_spec",fields, reflect.TypeOf(SataCreateSpec{}), fieldNameMap, validators)
}

func SataSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["adapter"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SataAdapter"}, "")
    fieldNameMap["adapter"] = "Adapter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.sata.summary",fields, reflect.TypeOf(SataSummary{}), fieldNameMap, validators)
}


