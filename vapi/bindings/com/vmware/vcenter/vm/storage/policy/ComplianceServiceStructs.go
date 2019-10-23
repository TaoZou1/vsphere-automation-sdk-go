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

package policy

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "time"
)



// The ``Status`` enumeration class defines the storage compliance status of a virtual machine and its applicable entities.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Compliance_Status string

const (
    // Entity is in compliance.
     Compliance_Status_COMPLIANT Compliance_Status = "COMPLIANT"
    // Entity is out of compliance.
     Compliance_Status_NON_COMPLIANT Compliance_Status = "NON_COMPLIANT"
    // Compliance status of the entity is not known.
     Compliance_Status_UNKNOWN_COMPLIANCE Compliance_Status = "UNKNOWN_COMPLIANCE"
    // Compliance computation is not applicable for this entity because it does not have any storage requirements that apply to the datastore on which it is placed.
     Compliance_Status_NOT_APPLICABLE Compliance_Status = "NOT_APPLICABLE"
    // The Compliance status becomes out-of-date when the profile associated with the entity is edited but not applied. The compliance status remains out-of-date until the edited policy is applied to the entity.
     Compliance_Status_OUT_OF_DATE Compliance_Status = "OUT_OF_DATE"
)

func (s Compliance_Status) Compliance_Status() bool {
    switch s {
        case Compliance_Status_COMPLIANT:
            return true
        case Compliance_Status_NON_COMPLIANT:
            return true
        case Compliance_Status_UNKNOWN_COMPLIANCE:
            return true
        case Compliance_Status_NOT_APPLICABLE:
            return true
        case Compliance_Status_OUT_OF_DATE:
            return true
        default:
            return false
    }
}





// The ``VmComplianceInfo`` class contains information about storage policy compliance associated with a virtual machine.
 type ComplianceVmComplianceInfo struct {
    // Status of the compliance operation.
    Status Compliance_Status
    // Date and time of the most recent compliance check.
    CheckTime time.Time
    // Identifier of the storage policy associated with the virtual machine.
    Policy *string
    // The exception that caused the compliance check to fail. There can be more than one cause, since a policy can contain capabilities from multiple providers. If empty, it implies no failures while retrieving compliance.
    FailureCause []std.LocalizableMessage
}



func (ComplianceVmComplianceInfo ComplianceVmComplianceInfo) Error() string {
    return "com.vmware.vcenter.vm.storage.policy.vm_compliance_info"
}



// The ``Info`` class contains information about the storage policy compliance of a virtual machine, including information about it's home directory and/or it's virtual disks.
 type ComplianceInfo struct {
    // The overall compliance status of the virtual machine and all it's entities.
    OverallCompliance Compliance_Status
    // The storage policy compliance information ComplianceVmComplianceInfo for the virtual machine's home directory.
    VmHome *ComplianceVmComplianceInfo
    // The compliance information ComplianceVmComplianceInfo for the virtual machine's virtual disks that are currently associated with a storage policy.
    Disks map[string]ComplianceVmComplianceInfo
}



func (ComplianceInfo ComplianceInfo) Error() string {
    return "com.vmware.vcenter.vm.storage.policy.info"
}



// The ``CheckSpec`` class contains properties used to specify the entities on which the storage policy compliance check is to be invoked.
 type ComplianceCheckSpec struct {
    // Invoke compliance check on the virtual machine home directory if set to true.
    VmHome bool
    // Identifiers of the virtual machine's virtual disks for which compliance should be checked.
    Disks map[string]bool
}



func (ComplianceCheckSpec ComplianceCheckSpec) Error() string {
    return "com.vmware.vcenter.vm.storage.policy.check_spec"
}






func complianceGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func complianceGetOutputType() bindings.BindingType {
    return bindings.NewOptionalType(bindings.NewReferenceType(ComplianceInfoBindingType))
}

func complianceGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func complianceCheckInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["check_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ComplianceCheckSpecBindingType))
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["check_spec"] = "CheckSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func complianceCheckOutputType() bindings.BindingType {
    return bindings.NewOptionalType(bindings.NewReferenceType(ComplianceInfoBindingType))
}

func complianceCheckRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func ComplianceVmComplianceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.vm.storage.policy.compliance.status", reflect.TypeOf(Compliance_Status(Compliance_Status_COMPLIANT)))
    fieldNameMap["status"] = "Status"
    fields["check_time"] = bindings.NewDateTimeType()
    fieldNameMap["check_time"] = "CheckTime"
    fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, ""))
    fieldNameMap["policy"] = "Policy"
    fields["failure_cause"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["failure_cause"] = "FailureCause"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.compliance.vm_compliance_info",fields, reflect.TypeOf(ComplianceVmComplianceInfo{}), fieldNameMap, validators)
}

func ComplianceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["overall_compliance"] = bindings.NewEnumType("com.vmware.vcenter.vm.storage.policy.compliance.status", reflect.TypeOf(Compliance_Status(Compliance_Status_COMPLIANT)))
    fieldNameMap["overall_compliance"] = "OverallCompliance"
    fields["vm_home"] = bindings.NewOptionalType(bindings.NewReferenceType(ComplianceVmComplianceInfoBindingType))
    fieldNameMap["vm_home"] = "VmHome"
    fields["disks"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(ComplianceVmComplianceInfoBindingType),reflect.TypeOf(map[string]ComplianceVmComplianceInfo{}))
    fieldNameMap["disks"] = "Disks"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.compliance.info",fields, reflect.TypeOf(ComplianceInfo{}), fieldNameMap, validators)
}

func ComplianceCheckSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_home"] = bindings.NewBooleanType()
    fieldNameMap["vm_home"] = "VmHome"
    fields["disks"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["disks"] = "Disks"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.compliance.check_spec",fields, reflect.TypeOf(ComplianceCheckSpec{}), fieldNameMap, validators)
}


