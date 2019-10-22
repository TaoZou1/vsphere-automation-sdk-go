/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Nodes.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package namespace_management

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The resource types allowed for node. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
var Nodes_NODE_TYPES = []string{"HostSystem", "VirtualMachine"}


// The ``MaintenanceActionType`` class contains actions to be taken when a node enters maintenance mode. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Nodes_MaintenanceActionType string

const (
    // The node is getting permanently removed. Move applications, rebuild storage on other nodes before allowing to proceed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Nodes_MaintenanceActionType_PermanentNodeRemoval Nodes_MaintenanceActionType = "PermanentNodeRemoval"
    // The node is going down temporarily for maintenance. Still need to ensure application availability and storage accessibility at least in a degraded level. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Nodes_MaintenanceActionType_TemporaryEnsureAccessibility Nodes_MaintenanceActionType = "TemporaryEnsureAccessibility"
    // Admin override to not delay or stop the node from entering maintenance mode. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Nodes_MaintenanceActionType_NoAction Nodes_MaintenanceActionType = "NoAction"
)

func (m Nodes_MaintenanceActionType) Nodes_MaintenanceActionType() bool {
    switch m {
        case Nodes_MaintenanceActionType_PermanentNodeRemoval:
            return true
        case Nodes_MaintenanceActionType_TemporaryEnsureAccessibility:
            return true
        case Nodes_MaintenanceActionType_NoAction:
            return true
        default:
            return false
    }
}





// The ``NodeIdentity`` class contains information about a node. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type NodesNodeIdentity struct {
    // Type for the node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NodeType string
    // Identifier for the node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NodeId string
}



//





func nodesEnterMaintenanceModeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["node"] = bindings.NewReferenceType(NodesNodeIdentityBindingType)
    fields["action"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.namespace_management.nodes.maintenance_action_type", reflect.TypeOf(Nodes_MaintenanceActionType(Nodes_MaintenanceActionType_PermanentNodeRemoval))))
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["node"] = "Node"
    fieldNameMap["action"] = "Action"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nodesEnterMaintenanceModeOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"Task"}, "")
}

func nodesEnterMaintenanceModeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"AlreadyInDesiredState": 400,"NotFound": 404,"Unsupported": 400,"InvalidElementConfiguration": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func nodesExitMaintenanceModeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["node"] = bindings.NewReferenceType(NodesNodeIdentityBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["node"] = "Node"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nodesExitMaintenanceModeOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"Task"}, "")
}

func nodesExitMaintenanceModeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"AlreadyInDesiredState": 400,"NotFound": 404,"Unsupported": 400,"InvalidElementConfiguration": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func nodesRemoveInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["node"] = bindings.NewReferenceType(NodesNodeIdentityBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["node"] = "Node"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nodesRemoveOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"Task"}, "")
}

func nodesRemoveRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"AlreadyInDesiredState": 400,"NotFound": 404,"Unsupported": 400,"InvalidElementConfiguration": 400,"Unauthenticated": 401,"Unauthorized": 403})
}



func NodesNodeIdentityBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["node_type"] = bindings.NewStringType()
    fieldNameMap["node_type"] = "NodeType"
    fields["node_id"] = bindings.NewIdType([]string {"HostSystem", "VirtualMachine"}, "node_type")
    fieldNameMap["node_id"] = "NodeId"
    var validators = []bindings.Validator{}
    isv1 := bindings.NewIsOneOfValidator(
        "node_type",
        []string {
             "HostSystem",
             "VirtualMachine",
        },
    )
    validators = append(validators, isv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.nodes.node_identity",fields, reflect.TypeOf(NodesNodeIdentity{}), fieldNameMap, validators)
}


