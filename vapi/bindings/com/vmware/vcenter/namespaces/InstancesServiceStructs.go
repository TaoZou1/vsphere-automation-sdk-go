/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Instances.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package namespaces

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The resource type for namespace. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
const Instances_RESOURCE_TYPE = "com.vmware.vcenter.namespaces.Instance"


// The ``ConfigStatus`` enumeration class describes the status of reaching the desired state configuration for the namespace. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Instances_ConfigStatus string

const (
    // The configuration is being applied to the namespace. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Instances_ConfigStatus_CONFIGURING Instances_ConfigStatus = "CONFIGURING"
    // The configuration is being removed and namespace is being deleted. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Instances_ConfigStatus_REMOVING Instances_ConfigStatus = "REMOVING"
    // The namespace is configured correctly. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Instances_ConfigStatus_RUNNING Instances_ConfigStatus = "RUNNING"
    // Failed to apply the configuration to the namespace, user intervention needed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Instances_ConfigStatus_ERROR Instances_ConfigStatus = "ERROR"
)

func (c Instances_ConfigStatus) Instances_ConfigStatus() bool {
    switch c {
        case Instances_ConfigStatus_CONFIGURING:
            return true
        case Instances_ConfigStatus_REMOVING:
            return true
        case Instances_ConfigStatus_RUNNING:
            return true
        case Instances_ConfigStatus_ERROR:
            return true
        default:
            return false
    }
}





// The ``Access`` class contains the access control information for a subject on a namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type InstancesAccess struct {
    // Type of the subject. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SubjectType Access_SubjectType
    // Name of the subject. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Subject string
    // Domain of the subject. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Domain string
    // Role of the subject on the namespace instance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Role Access_Role
}



func (InstancesAccess InstancesAccess) Error() string {
    return "com.vmware.vcenter.namespaces.access"
}



// The ``StorageSpec`` class contains the specification required to configure storage associated with a namespace. Information in this class will result in storage quotas on the Kubernetes namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type InstancesStorageSpec struct {
    // ID of the storage policy. A Kubernetes storage class is created for this storage policy if it does not exist already. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Policy string
    // The maximum amount of storage (in mebibytes) which can be utilized by the namespace for this specification. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Limit *int64
}



func (InstancesStorageSpec InstancesStorageSpec) Error() string {
    return "com.vmware.vcenter.namespaces.storage_spec"
}



// The ``Message`` class contains the information about the object configuration. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type InstancesMessage struct {
    // Type of the message. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Severity InstancesMessage_MessageSeverity
    // Details about the message. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Details *std.LocalizableMessage
}



func (InstancesMessage InstancesMessage) Error() string {
    return "com.vmware.vcenter.namespaces.message"
}

    
    // The ``MessageSeverity`` enumeration class represents the severity of the message. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type InstancesMessage_MessageSeverity string

    const (
        // Informational message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         InstancesMessage_MessageSeverity_INFO InstancesMessage_MessageSeverity = "INFO"
        // Warning message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         InstancesMessage_MessageSeverity_WARNING InstancesMessage_MessageSeverity = "WARNING"
        // Error message. This is accompanied by vCenter event and/or alarm. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         InstancesMessage_MessageSeverity_ERROR InstancesMessage_MessageSeverity = "ERROR"
    )

    func (m InstancesMessage_MessageSeverity) InstancesMessage_MessageSeverity() bool {
        switch m {
            case InstancesMessage_MessageSeverity_INFO:
                return true
            case InstancesMessage_MessageSeverity_WARNING:
                return true
            case InstancesMessage_MessageSeverity_ERROR:
                return true
            default:
                return false
        }
    }



// The ``Stats`` class contains the basic runtime statistics about the namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type InstancesStats struct {
    // Overall CPU usage of the namespace, in MHz. This is the sum of CPU usage across all pods in the Kubernetes namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuUsed int64
    // Overall memory usage of the namespace (in mebibytes). This is the sum of memory usage across all pods. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MemoryUsed int64
    // Overall storage used by the namespace (in mebibytes). This is the sum of storage used by pods across all datastores in the cluster associated with storage policies configured for the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StorageUsed int64
}



func (InstancesStats InstancesStats) Error() string {
    return "com.vmware.vcenter.namespaces.stats"
}



// The ``Summary`` class contains basic information about the namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type InstancesSummary struct {
    // Identifier of the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Namespace string
    // Description of the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Description string
    // Identifier for the cluster hosting the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Cluster string
    // Current setting for ``ConfigStatus``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ConfigStatus Instances_ConfigStatus
    // Basic runtime statistics for the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Stats InstancesStats
}



func (InstancesSummary InstancesSummary) Error() string {
    return "com.vmware.vcenter.namespaces.summary"
}



// The ``Info`` class contains detailed information about the namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type InstancesInfo struct {
    // Identifier for the cluster hosting the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Cluster string
    // Current setting for ``ConfigStatus``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ConfigStatus Instances_ConfigStatus
    // Current set of messages associated with the object. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Messages []InstancesMessage
    // Basic runtime statistics for the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Stats InstancesStats
    // Description of the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Description string
    // Quotas on the namespace resources. Refer to null for the type of the value for this field. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ResourceSpec *data.StructValue
    // Access controls associated with the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AccessList []InstancesAccess
    // Storage associated with the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StorageSpecs []InstancesStorageSpec
}



func (InstancesInfo InstancesInfo) Error() string {
    return "com.vmware.vcenter.namespaces.info"
}



// The ``UpdateSpec`` class contains the specification required to update the configuration on the namespace. This class is applied partially, and only the specified fields will replace or modify their existing counterparts. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type InstancesUpdateSpec struct {
    // Description for the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Description *string
    // Resource quota updates on the namespace. Refer to null and use null for retrieving the type for the value for this field. For an example of this, see ResourceQuotaOptionsV1Update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ResourceSpec *data.StructValue
    // Access control associated with the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AccessList []InstancesAccess
    // Storage associated with the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StorageSpecs []InstancesStorageSpec
}



func (InstancesUpdateSpec InstancesUpdateSpec) Error() string {
    return "com.vmware.vcenter.namespaces.update_spec"
}



// The ``SetSpec`` class contains the specification required to set a new configuration on the namespace. This class is applied in entirety, replacing the current specification fully. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type InstancesSetSpec struct {
    // Description for the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Description *string
    // Resource quota for the namespace. This will replace the existing resource constraints on the namespace in entirety. Refer to null and use null for retrieving the type for the value for this field. For an example of this, see ResourceQuotaOptionsV1. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ResourceSpec *data.StructValue
    // Access control associated with the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AccessList []InstancesAccess
    // Storage associated with the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StorageSpecs []InstancesStorageSpec
}



func (InstancesSetSpec InstancesSetSpec) Error() string {
    return "com.vmware.vcenter.namespaces.set_spec"
}



// The ``CreateSpec`` class contains the specification required to set up a namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type InstancesCreateSpec struct {
    // Identifier of the namespace. This has DNS_LABEL restrictions as specified in ` <https://tools.ietf.org/html/rfc1123>`_. This must be an alphanumeric (a-z and 0-9) string and with maximum length of 63 characters and with the '-' character allowed anywhere except the first or last character. This name is unique across all Namespaces in this vCenter server. In this version, this maps to the name of a Kubernetes namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Namespace string
    // Identifier of the cluster on which the namespace is being created. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Cluster string
    // Description for the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Description *string
    // Resource quota on the namespace. Refer to null and use null for retrieving the type for the value for this field. For an example of this, see ResourceQuotaOptionsV1. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ResourceSpec *data.StructValue
    // Access controls associated with the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AccessList []InstancesAccess
    // Storage associated with the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StorageSpecs []InstancesStorageSpec
}



func (InstancesCreateSpec InstancesCreateSpec) Error() string {
    return "com.vmware.vcenter.namespaces.create_spec"
}






func instancesCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(InstancesCreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instancesCreateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func instancesCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(InstancesCreateSpecBindingType)
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
      "/vcenter/namespaces/instances",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unsupported": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func instancesDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instancesDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func instancesDeleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    pathParams["namespace"] = "namespace"
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
      "/vcenter/namespaces/instances/{namespace}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func instancesGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instancesGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InstancesInfoBindingType)
}

func instancesGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    pathParams["namespace"] = "namespace"
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
      "/vcenter/namespaces/instances/{namespace}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func instancesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instancesListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(InstancesSummaryBindingType), reflect.TypeOf([]InstancesSummary{}))
}

func instancesListRestMetadata() protocol.OperationRestMetadata {
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
      "GET",
      "/vcenter/namespaces/instances",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func instancesSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fields["spec"] = bindings.NewReferenceType(InstancesSetSpecBindingType)
    fieldNameMap["namespace"] = "Namespace"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instancesSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func instancesSetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(InstancesSetSpecBindingType)
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    pathParams["namespace"] = "namespace"
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
      "/vcenter/namespaces/instances/{namespace}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func instancesUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fields["spec"] = bindings.NewReferenceType(InstancesUpdateSpecBindingType)
    fieldNameMap["namespace"] = "Namespace"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instancesUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func instancesUpdateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(InstancesUpdateSpecBindingType)
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    pathParams["namespace"] = "namespace"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PATCH",
      "/vcenter/namespaces/instances/{namespace}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func InstancesAccessBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["subject_type"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.subject_type", reflect.TypeOf(Access_SubjectType(Access_SubjectType_USER)))
    fieldNameMap["subject_type"] = "SubjectType"
    fields["subject"] = bindings.NewStringType()
    fieldNameMap["subject"] = "Subject"
    fields["domain"] = bindings.NewStringType()
    fieldNameMap["domain"] = "Domain"
    fields["role"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.role", reflect.TypeOf(Access_Role(Access_Role_EDIT)))
    fieldNameMap["role"] = "Role"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.access",fields, reflect.TypeOf(InstancesAccess{}), fieldNameMap, validators)
}

func InstancesStorageSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["policy"] = "Policy"
    fields["limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["limit"] = "Limit"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.storage_spec",fields, reflect.TypeOf(InstancesStorageSpec{}), fieldNameMap, validators)
}

func InstancesMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.instances.message.message_severity", reflect.TypeOf(InstancesMessage_MessageSeverity(InstancesMessage_MessageSeverity_INFO)))
    fieldNameMap["severity"] = "Severity"
    fields["details"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.message",fields, reflect.TypeOf(InstancesMessage{}), fieldNameMap, validators)
}

func InstancesStatsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cpu_used"] = bindings.NewIntegerType()
    fieldNameMap["cpu_used"] = "CpuUsed"
    fields["memory_used"] = bindings.NewIntegerType()
    fieldNameMap["memory_used"] = "MemoryUsed"
    fields["storage_used"] = bindings.NewIntegerType()
    fieldNameMap["storage_used"] = "StorageUsed"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.stats",fields, reflect.TypeOf(InstancesStats{}), fieldNameMap, validators)
}

func InstancesSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.instances.config_status", reflect.TypeOf(Instances_ConfigStatus(Instances_ConfigStatus_CONFIGURING)))
    fieldNameMap["config_status"] = "ConfigStatus"
    fields["stats"] = bindings.NewReferenceType(InstancesStatsBindingType)
    fieldNameMap["stats"] = "Stats"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.summary",fields, reflect.TypeOf(InstancesSummary{}), fieldNameMap, validators)
}

func InstancesInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.instances.config_status", reflect.TypeOf(Instances_ConfigStatus(Instances_ConfigStatus_CONFIGURING)))
    fieldNameMap["config_status"] = "ConfigStatus"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(InstancesMessageBindingType), reflect.TypeOf([]InstancesMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["stats"] = bindings.NewReferenceType(InstancesStatsBindingType)
    fieldNameMap["stats"] = "Stats"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["resource_spec"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["resource_spec"] = "ResourceSpec"
    fields["access_list"] = bindings.NewListType(bindings.NewReferenceType(InstancesAccessBindingType), reflect.TypeOf([]InstancesAccess{}))
    fieldNameMap["access_list"] = "AccessList"
    fields["storage_specs"] = bindings.NewListType(bindings.NewReferenceType(InstancesStorageSpecBindingType), reflect.TypeOf([]InstancesStorageSpec{}))
    fieldNameMap["storage_specs"] = "StorageSpecs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.info",fields, reflect.TypeOf(InstancesInfo{}), fieldNameMap, validators)
}

func InstancesUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["resource_spec"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["resource_spec"] = "ResourceSpec"
    fields["access_list"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(InstancesAccessBindingType), reflect.TypeOf([]InstancesAccess{})))
    fieldNameMap["access_list"] = "AccessList"
    fields["storage_specs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(InstancesStorageSpecBindingType), reflect.TypeOf([]InstancesStorageSpec{})))
    fieldNameMap["storage_specs"] = "StorageSpecs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.update_spec",fields, reflect.TypeOf(InstancesUpdateSpec{}), fieldNameMap, validators)
}

func InstancesSetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["resource_spec"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["resource_spec"] = "ResourceSpec"
    fields["access_list"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(InstancesAccessBindingType), reflect.TypeOf([]InstancesAccess{})))
    fieldNameMap["access_list"] = "AccessList"
    fields["storage_specs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(InstancesStorageSpecBindingType), reflect.TypeOf([]InstancesStorageSpec{})))
    fieldNameMap["storage_specs"] = "StorageSpecs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.set_spec",fields, reflect.TypeOf(InstancesSetSpec{}), fieldNameMap, validators)
}

func InstancesCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["resource_spec"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["resource_spec"] = "ResourceSpec"
    fields["access_list"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(InstancesAccessBindingType), reflect.TypeOf([]InstancesAccess{})))
    fieldNameMap["access_list"] = "AccessList"
    fields["storage_specs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(InstancesStorageSpecBindingType), reflect.TypeOf([]InstancesStorageSpec{})))
    fieldNameMap["storage_specs"] = "StorageSpecs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.create_spec",fields, reflect.TypeOf(InstancesCreateSpec{}), fieldNameMap, validators)
}


