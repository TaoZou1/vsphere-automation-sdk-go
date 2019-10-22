/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Scsi.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package adapter

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/vm/hardware"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for the virtual SCSI adapter device.
const Scsi_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.ScsiAdapter"


// The ``Type`` enumeration class defines the valid emulation types for a virtual SCSI adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Scsi_Type string

const (
    // BusLogic host bus adapter.
     Scsi_Type_BUSLOGIC Scsi_Type = "BUSLOGIC"
    // LSI Logic host bus adapter.
     Scsi_Type_LSILOGIC Scsi_Type = "LSILOGIC"
    // LSI Logic SAS 1068 host bus adapter.
     Scsi_Type_LSILOGICSAS Scsi_Type = "LSILOGICSAS"
    // Paravirtualized host bus adapter.
     Scsi_Type_PVSCSI Scsi_Type = "PVSCSI"
)

func (t Scsi_Type) Scsi_Type() bool {
    switch t {
        case Scsi_Type_BUSLOGIC:
            return true
        case Scsi_Type_LSILOGIC:
            return true
        case Scsi_Type_LSILOGICSAS:
            return true
        case Scsi_Type_PVSCSI:
            return true
        default:
            return false
    }
}




// The ``Sharing`` enumeration class defines the valid bus sharing modes for a virtual SCSI adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Scsi_Sharing string

const (
    // The virtual SCSI bus is not shared.
     Scsi_Sharing_NONE Scsi_Sharing = "NONE"
    // The virtual SCSI bus is shared between two or more virtual machines. In this case, no physical machine is involved.
     Scsi_Sharing_VIRTUAL Scsi_Sharing = "VIRTUAL"
    // The virtual SCSI bus is shared between two or more virtual machines residing on different physical hosts.
     Scsi_Sharing_PHYSICAL Scsi_Sharing = "PHYSICAL"
)

func (s Scsi_Sharing) Scsi_Sharing() bool {
    switch s {
        case Scsi_Sharing_NONE:
            return true
        case Scsi_Sharing_VIRTUAL:
            return true
        case Scsi_Sharing_PHYSICAL:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information about a virtual SCSI adapter.
 type ScsiInfo struct {
    // Device label.
    Label string
    // Adapter type.
    Type_ Scsi_Type
    // Address of the SCSI adapter on the SCSI bus.
    Scsi hardware.ScsiAddressInfo
    // Address of the SCSI adapter on the PCI bus. If the PCI address is invalid, the server will change it when the VM is started or as the device is hot added.
    PciSlotNumber *int64
    // Bus sharing mode.
    Sharing Scsi_Sharing
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual SCSI adapter.
 type ScsiCreateSpec struct {
    // Adapter type.
    Type_ *Scsi_Type
    // SCSI bus number.
    Bus *int64
    // Address of the SCSI adapter on the PCI bus. If the PCI address is invalid, the server will change it when the VM is started or as the device is hot added.
    PciSlotNumber *int64
    // Bus sharing mode.
    Sharing *Scsi_Sharing
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual SCSI adapter.
 type ScsiUpdateSpec struct {
    // Bus sharing mode. 
//
//  This property may only be modified if the virtual machine is not powered on.
    Sharing *Scsi_Sharing
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Summary`` class contains commonly used information about a Virtual SCSI adapter.
 type ScsiSummary struct {
    // Identifier of the virtual SCSI adapter.
    Adapter string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func scsiListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func scsiListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ScsiSummaryBindingType), reflect.TypeOf([]ScsiSummary{}))
}

func scsiListRestMetadata() protocol.OperationRestMetadata {
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


func scsiGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["adapter"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["adapter"] = "Adapter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func scsiGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ScsiInfoBindingType)
}

func scsiGetRestMetadata() protocol.OperationRestMetadata {
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


func scsiCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewReferenceType(ScsiCreateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func scsiCreateOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, "")
}

func scsiCreateRestMetadata() protocol.OperationRestMetadata {
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


func scsiUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["adapter"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, "")
    fields["spec"] = bindings.NewReferenceType(ScsiUpdateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["adapter"] = "Adapter"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func scsiUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func scsiUpdateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func scsiDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["adapter"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, "")
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["adapter"] = "Adapter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func scsiDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func scsiDeleteRestMetadata() protocol.OperationRestMetadata {
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



func ScsiInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.adapter.scsi.type", reflect.TypeOf(Scsi_Type(Scsi_Type_BUSLOGIC)))
    fieldNameMap["type"] = "Type_"
    fields["scsi"] = bindings.NewReferenceType(hardware.ScsiAddressInfoBindingType)
    fieldNameMap["scsi"] = "Scsi"
    fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pci_slot_number"] = "PciSlotNumber"
    fields["sharing"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.adapter.scsi.sharing", reflect.TypeOf(Scsi_Sharing(Scsi_Sharing_NONE)))
    fieldNameMap["sharing"] = "Sharing"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.scsi.info",fields, reflect.TypeOf(ScsiInfo{}), fieldNameMap, validators)
}

func ScsiCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.adapter.scsi.type", reflect.TypeOf(Scsi_Type(Scsi_Type_BUSLOGIC))))
    fieldNameMap["type"] = "Type_"
    fields["bus"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["bus"] = "Bus"
    fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pci_slot_number"] = "PciSlotNumber"
    fields["sharing"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.adapter.scsi.sharing", reflect.TypeOf(Scsi_Sharing(Scsi_Sharing_NONE))))
    fieldNameMap["sharing"] = "Sharing"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.scsi.create_spec",fields, reflect.TypeOf(ScsiCreateSpec{}), fieldNameMap, validators)
}

func ScsiUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sharing"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.adapter.scsi.sharing", reflect.TypeOf(Scsi_Sharing(Scsi_Sharing_NONE))))
    fieldNameMap["sharing"] = "Sharing"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.scsi.update_spec",fields, reflect.TypeOf(ScsiUpdateSpec{}), fieldNameMap, validators)
}

func ScsiSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["adapter"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, "")
    fieldNameMap["adapter"] = "Adapter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.scsi.summary",fields, reflect.TypeOf(ScsiSummary{}), fieldNameMap, validators)
}


