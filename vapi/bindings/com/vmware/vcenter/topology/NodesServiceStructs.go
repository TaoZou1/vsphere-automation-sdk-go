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

package topology

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``ApplianceType`` enumeration class defines values for valid appliance types for the vCenter and Platform Services Controller node. See NodesInfo.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Nodes_ApplianceType string

const (
    // vCenter Server Appliance with an embedded Platform Services Controller.
     Nodes_ApplianceType_VCSA_EMBEDDED Nodes_ApplianceType = "VCSA_EMBEDDED"
    // vCenter Server Appliance with an external Platform Services Controller.
     Nodes_ApplianceType_VCSA_EXTERNAL Nodes_ApplianceType = "VCSA_EXTERNAL"
    // An external Platform Services Controller.
     Nodes_ApplianceType_PSC_EXTERNAL Nodes_ApplianceType = "PSC_EXTERNAL"
)

func (a Nodes_ApplianceType) Nodes_ApplianceType() bool {
    switch a {
        case Nodes_ApplianceType_VCSA_EMBEDDED:
            return true
        case Nodes_ApplianceType_VCSA_EXTERNAL:
            return true
        case Nodes_ApplianceType_PSC_EXTERNAL:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains vCenter or Platform Services Controller node details.
 type NodesInfo struct {
    // Domain name of the node.
    Domain string
    // Appliance type of the node.
    Type_ Nodes_ApplianceType
    // List of replication partners' node identifiers. Identifiers can be either IP address or DNS resolvable name of the partner node.
    ReplicationPartners []string
    // Identifier of the affinitized Platform Services Controller node. Identifier can be either IP address or DNS resolvable name of the affinitized node.
    ClientAffinity *string
}



func (NodesInfo NodesInfo) Error() string {
    return "com.vmware.vcenter.topology.info"
}



// The ``Summary`` class contains commonly used information of vCenter or Platform Services Controller node.
 type NodesSummary struct {
    // Identifier for the vCenter or Platform Services Controller node. Identifier can be either IP address or DNS resolvable name of the node.
    Node string
    // Appliance type of the node.
    Type_ Nodes_ApplianceType
    // List of replication partners' node identifiers. Identifiers can be either IP address or DNS resolvable name of the partner node.
    ReplicationPartners []string
    // Identifier of the affinitized Platform Services Controller node. Identifier can be either IP address or DNS resolvable name of the affinitized node.
    ClientAffinity *string
}



func (NodesSummary NodesSummary) Error() string {
    return "com.vmware.vcenter.topology.summary"
}



// The ``FilterSpec`` class contains property used to filter the results when listing vCenter and Platform Services Controller nodes (see Nodes#list).
 type NodesFilterSpec struct {
    // Types of the appliance that a vCenter and Platform Services Controller node must be to match the filter (see Nodes_ApplianceType.
    Types map[Nodes_ApplianceType]bool
}



func (NodesFilterSpec NodesFilterSpec) Error() string {
    return "com.vmware.vcenter.topology.filter_spec"
}






func nodesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(NodesFilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nodesListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(NodesSummaryBindingType), reflect.TypeOf([]NodesSummary{}))
}

func nodesListRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"InvalidArgument": 400})
}


func nodesGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["node"] = bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, "")
    fieldNameMap["node"] = "Node"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nodesGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(NodesInfoBindingType)
}

func nodesGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404})
}



func NodesInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["domain"] = bindings.NewStringType()
    fieldNameMap["domain"] = "Domain"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.topology.nodes.appliance_type", reflect.TypeOf(Nodes_ApplianceType(Nodes_ApplianceType_VCSA_EMBEDDED)))
    fieldNameMap["type"] = "Type_"
    fields["replication_partners"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, ""), reflect.TypeOf([]string{})))
    fieldNameMap["replication_partners"] = "ReplicationPartners"
    fields["client_affinity"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, ""))
    fieldNameMap["client_affinity"] = "ClientAffinity"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "VCSA_EMBEDDED": []bindings.FieldData {
                 bindings.NewFieldData("replication_partners", true),
            },
            "PSC_EXTERNAL": []bindings.FieldData {
                 bindings.NewFieldData("replication_partners", true),
            },
            "VCSA_EXTERNAL": []bindings.FieldData {
                 bindings.NewFieldData("client_affinity", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.topology.nodes.info",fields, reflect.TypeOf(NodesInfo{}), fieldNameMap, validators)
}

func NodesSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["node"] = bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, "")
    fieldNameMap["node"] = "Node"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.topology.nodes.appliance_type", reflect.TypeOf(Nodes_ApplianceType(Nodes_ApplianceType_VCSA_EMBEDDED)))
    fieldNameMap["type"] = "Type_"
    fields["replication_partners"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, ""), reflect.TypeOf([]string{})))
    fieldNameMap["replication_partners"] = "ReplicationPartners"
    fields["client_affinity"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, ""))
    fieldNameMap["client_affinity"] = "ClientAffinity"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "VCSA_EMBEDDED": []bindings.FieldData {
                 bindings.NewFieldData("replication_partners", true),
            },
            "PSC_EXTERNAL": []bindings.FieldData {
                 bindings.NewFieldData("replication_partners", true),
            },
            "VCSA_EXTERNAL": []bindings.FieldData {
                 bindings.NewFieldData("client_affinity", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.topology.nodes.summary",fields, reflect.TypeOf(NodesSummary{}), fieldNameMap, validators)
}

func NodesFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.topology.nodes.appliance_type", reflect.TypeOf(Nodes_ApplianceType(Nodes_ApplianceType_VCSA_EMBEDDED))), reflect.TypeOf(map[Nodes_ApplianceType]bool{})))
    fieldNameMap["types"] = "Types"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.topology.nodes.filter_spec",fields, reflect.TypeOf(NodesFilterSpec{}), fieldNameMap, validators)
}


