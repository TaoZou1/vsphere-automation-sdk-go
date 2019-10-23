/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Witness.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package cluster

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/vcha"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``CheckSpec`` class contains placement information for validation.
 type WitnessCheckSpec struct {
    // Contains the active node's management vCenter server credentials.
    VcSpec *vcha.CredentialsSpec
    // Contains the node's placement information for validation.
    Placement vcha.PlacementSpec
}



func (WitnessCheckSpec WitnessCheckSpec) Error() string {
    return "com.vmware.vcenter.vcha.cluster.check_spec"
}



// The ``CheckResult`` class contains the warnings and errors that will occur during the clone operation.
 type WitnessCheckResult struct {
    // A list of problems which may require attention, but which are not fatal.
    Warnings []std.LocalizableMessage
    // A list of problems which are fatal to the operation and the operation will fail.
    Errors []std.LocalizableMessage
}



func (WitnessCheckResult WitnessCheckResult) Error() string {
    return "com.vmware.vcenter.vcha.cluster.check_result"
}



// The ``RedeploySpec`` class contains the redeploy specification.
 type WitnessRedeploySpec struct {
    // Contains the active node's management vCenter server credentials.
    VcSpec *vcha.CredentialsSpec
    // Contains the node's placement information.
    Placement vcha.PlacementSpec
    // Contains the VCHA HA network configuration of the node. All cluster communication (state replication, heartbeat, cluster messages) happens over this network.
    HaIp *vcha.IpSpec
}



func (WitnessRedeploySpec WitnessRedeploySpec) Error() string {
    return "com.vmware.vcenter.vcha.cluster.redeploy_spec"
}






func witnessCheckInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(WitnessCheckSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func witnessCheckOutputType() bindings.BindingType {
    return bindings.NewReferenceType(WitnessCheckResultBindingType)
}

func witnessCheckRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"InvalidArgument": 400,"UnverifiedPeer": 400,"NotFound": 404,"InvalidElementConfiguration": 400,"NotAllowedInCurrentState": 400,"Unauthorized": 403,"Error": 500})
}


func witnessRedeployInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(WitnessRedeploySpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func witnessRedeployOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func witnessRedeployRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"InvalidArgument": 400,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}



func WitnessCheckSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.CredentialsSpecBindingType))
    fieldNameMap["vc_spec"] = "VcSpec"
    fields["placement"] = bindings.NewReferenceType(vcha.PlacementSpecBindingType)
    fieldNameMap["placement"] = "Placement"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.witness.check_spec",fields, reflect.TypeOf(WitnessCheckSpec{}), fieldNameMap, validators)
}

func WitnessCheckResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.witness.check_result",fields, reflect.TypeOf(WitnessCheckResult{}), fieldNameMap, validators)
}

func WitnessRedeploySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.CredentialsSpecBindingType))
    fieldNameMap["vc_spec"] = "VcSpec"
    fields["placement"] = bindings.NewReferenceType(vcha.PlacementSpecBindingType)
    fieldNameMap["placement"] = "Placement"
    fields["ha_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.IpSpecBindingType))
    fieldNameMap["ha_ip"] = "HaIp"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.witness.redeploy_spec",fields, reflect.TypeOf(WitnessRedeploySpec{}), fieldNameMap, validators)
}


