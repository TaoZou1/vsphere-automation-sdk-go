/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Compliance.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package policies

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// This enumeration defines the set of status values for a compliance operation. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Compliance_Status string

const (
    // Entity is in compliance. This constant field was added in vSphere API 6.7.
     Compliance_Status_COMPLIANT Compliance_Status = "COMPLIANT"
    // Entity is out of compliance. This constant field was added in vSphere API 6.7.
     Compliance_Status_NON_COMPLIANT Compliance_Status = "NON_COMPLIANT"
    // Compliance status of the entity is not known. This constant field was added in vSphere API 6.7.
     Compliance_Status_UNKNOWN Compliance_Status = "UNKNOWN"
    // Compliance computation is not applicable for this entity because it does not have any storage requirement that apply to the object-based datastore on which the entity is placed. This constant field was added in vSphere API 6.7.
     Compliance_Status_NOT_APPLICABLE Compliance_Status = "NOT_APPLICABLE"
    // Compliance status becomes out of date when the profile associated with the entity is edited and not applied. The compliance status will remain out of date until the latest policy is applied to the entity. This constant field was added in vSphere API 6.7.
     Compliance_Status_OUT_OF_DATE Compliance_Status = "OUT_OF_DATE"
)

func (s Compliance_Status) Compliance_Status() bool {
    switch s {
        case Compliance_Status_COMPLIANT:
            return true
        case Compliance_Status_NON_COMPLIANT:
            return true
        case Compliance_Status_UNKNOWN:
            return true
        case Compliance_Status_NOT_APPLICABLE:
            return true
        case Compliance_Status_OUT_OF_DATE:
            return true
        default:
            return false
    }
}





// Provides the details of a virtual machine and its associated entities which match the given compliance statuses. This class was added in vSphere API 6.7.
 type ComplianceSummary struct {
    // Identifier of virtual machine. This property was added in vSphere API 6.7.
    Vm string
    // Compliance status of the virtual machine home. This property was added in vSphere API 6.7.
    VmHome *Compliance_Status
    // List of the virtual hard disk. This property was added in vSphere API 6.7.
    Disks map[string]Compliance_Status
}



//


// The ``FilterSpec`` class contains complianceStatus used to filter the results when listing entities (see Compliance#list). This class was added in vSphere API 6.7.
 type ComplianceFilterSpec struct {
    // Compliance Status that a virtual machine must have to match the filter. This property was added in vSphere API 6.7.
    Status map[Compliance_Status]bool
}



//





func complianceListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewReferenceType(ComplianceFilterSpecBindingType)
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func complianceListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ComplianceSummaryBindingType), reflect.TypeOf([]ComplianceSummary{}))
}

func complianceListRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func ComplianceSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    fields["vm_home"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.storage.policies.compliance.status", reflect.TypeOf(Compliance_Status(Compliance_Status_COMPLIANT))))
    fieldNameMap["vm_home"] = "VmHome"
    fields["disks"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewEnumType("com.vmware.vcenter.storage.policies.compliance.status", reflect.TypeOf(Compliance_Status(Compliance_Status_COMPLIANT))),reflect.TypeOf(map[string]Compliance_Status{})))
    fieldNameMap["disks"] = "Disks"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.storage.policies.compliance.summary",fields, reflect.TypeOf(ComplianceSummary{}), fieldNameMap, validators)
}

func ComplianceFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.storage.policies.compliance.status", reflect.TypeOf(Compliance_Status(Compliance_Status_COMPLIANT))), reflect.TypeOf(map[Compliance_Status]bool{}))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.storage.policies.compliance.filter_spec",fields, reflect.TypeOf(ComplianceFilterSpec{}), fieldNameMap, validators)
}


