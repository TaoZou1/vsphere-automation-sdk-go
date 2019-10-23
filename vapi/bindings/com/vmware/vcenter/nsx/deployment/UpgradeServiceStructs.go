/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Upgrade.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package deployment

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "net/url"
    "time"
)



// The ``ExecutionMode`` enumeration class defines the Execution mode of NSX environment. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Upgrade_ExecutionMode string

const (
    // The clusters with NSX enabled will be upgraded in parallel. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_ExecutionMode_BATCH Upgrade_ExecutionMode = "BATCH"
    // The clusters with NSX enabled will be upgraded in sequence one at a time. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_ExecutionMode_SERIAL Upgrade_ExecutionMode = "SERIAL"
)

func (e Upgrade_ExecutionMode) Upgrade_ExecutionMode() bool {
    switch e {
        case Upgrade_ExecutionMode_BATCH:
            return true
        case Upgrade_ExecutionMode_SERIAL:
            return true
        default:
            return false
    }
}




// The ``NSXNodeType`` enumeration class defines the type of the node in a cluster. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Upgrade_NSXNodeType string

const (
    // The node type to identify NSX Manager. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_NSXNodeType_MANAGER Upgrade_NSXNodeType = "MANAGER"
    // The node type to identify NSX Edge VM. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_NSXNodeType_EDGE Upgrade_NSXNodeType = "EDGE"
    // The node type to identify a Host. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_NSXNodeType_HOST Upgrade_NSXNodeType = "HOST"
)

func (n Upgrade_NSXNodeType) Upgrade_NSXNodeType() bool {
    switch n {
        case Upgrade_NSXNodeType_MANAGER:
            return true
        case Upgrade_NSXNodeType_EDGE:
            return true
        case Upgrade_NSXNodeType_HOST:
            return true
        default:
            return false
    }
}




// The ``DeploymentStatus`` enumeration class defines the overall upgrade status of the deployment, including the clusters and the nodes. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Upgrade_DeploymentStatus string

const (
    // The upgrade is not yet started for the deployment. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_DeploymentStatus_NOT_STARTED Upgrade_DeploymentStatus = "NOT_STARTED"
    // The upgrade is in progress for the deployment. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_DeploymentStatus_IN_PROGRESS Upgrade_DeploymentStatus = "IN_PROGRESS"
    // The upgrade will be paused after the current operation is completed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_DeploymentStatus_PAUSE_PENDING Upgrade_DeploymentStatus = "PAUSE_PENDING"
    // The upgrade is paused for the deployment. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_DeploymentStatus_PAUSED Upgrade_DeploymentStatus = "PAUSED"
    // The upgrade is completed for the deployment. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_DeploymentStatus_COMPLETED Upgrade_DeploymentStatus = "COMPLETED"
)

func (d Upgrade_DeploymentStatus) Upgrade_DeploymentStatus() bool {
    switch d {
        case Upgrade_DeploymentStatus_NOT_STARTED:
            return true
        case Upgrade_DeploymentStatus_IN_PROGRESS:
            return true
        case Upgrade_DeploymentStatus_PAUSE_PENDING:
            return true
        case Upgrade_DeploymentStatus_PAUSED:
            return true
        case Upgrade_DeploymentStatus_COMPLETED:
            return true
        default:
            return false
    }
}




// The ``ClusterStatus`` enumeration class defines the upgrade status of the cluster. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Upgrade_ClusterStatus string

const (
    // The upgrade is not yet started for the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_ClusterStatus_NOT_STARTED Upgrade_ClusterStatus = "NOT_STARTED"
    // The upgrade is in queue, and not yet started for the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_ClusterStatus_IN_QUEUE Upgrade_ClusterStatus = "IN_QUEUE"
    // The upgrade is in progress for the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_ClusterStatus_IN_PROGRESS Upgrade_ClusterStatus = "IN_PROGRESS"
    // The upgrade has failed for the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_ClusterStatus_FAILED Upgrade_ClusterStatus = "FAILED"
    // The upgrade is completed for the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_ClusterStatus_COMPLETED Upgrade_ClusterStatus = "COMPLETED"
)

func (c Upgrade_ClusterStatus) Upgrade_ClusterStatus() bool {
    switch c {
        case Upgrade_ClusterStatus_NOT_STARTED:
            return true
        case Upgrade_ClusterStatus_IN_QUEUE:
            return true
        case Upgrade_ClusterStatus_IN_PROGRESS:
            return true
        case Upgrade_ClusterStatus_FAILED:
            return true
        case Upgrade_ClusterStatus_COMPLETED:
            return true
        default:
            return false
    }
}




// The ``Status`` enumeration class defines the current status of any operation. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Upgrade_Status string

const (
    // The operation is not yet started. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_Status_NOT_STARTED Upgrade_Status = "NOT_STARTED"
    // The operation is in progress. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_Status_IN_PROGRESS Upgrade_Status = "IN_PROGRESS"
    // The operation is paused. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_Status_PAUSED Upgrade_Status = "PAUSED"
    // The operation has failed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_Status_FAILED Upgrade_Status = "FAILED"
    // The operation is completed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Upgrade_Status_COMPLETED Upgrade_Status = "COMPLETED"
)

func (s Upgrade_Status) Upgrade_Status() bool {
    switch s {
        case Upgrade_Status_NOT_STARTED:
            return true
        case Upgrade_Status_IN_PROGRESS:
            return true
        case Upgrade_Status_PAUSED:
            return true
        case Upgrade_Status_FAILED:
            return true
        case Upgrade_Status_COMPLETED:
            return true
        default:
            return false
    }
}





// The ``Messages`` class contains information about the info, warnings and erros if any. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpgradeMessages struct {
    // List of info messages. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Info []std.LocalizableMessage
    // List of warning messages. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Warnings []std.LocalizableMessage
    // List of error messages. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Errors []std.LocalizableMessage
}



func (UpgradeMessages UpgradeMessages) Error() string {
    return "com.vmware.vcenter.nsx.deployment.messages"
}



// The ``NSXNodeInfo`` class contains the information about the nodes in the cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpgradeNSXNodeInfo struct {
    // Unique identifier of the given node TODO: Based on the node type, this ID can be "HostSystem" (Host) or "VirtualMachine" (Edge, Manager). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Id string
    // Type of the node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NodeType Upgrade_NSXNodeType
    // The percent completion of the operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PercentComplete *float64
    // The precheck messages corresponding to the node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PrecheckMessages UpgradeMessages
    // The upgrade messages corresponding to the node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpgradeMessages UpgradeMessages
}



func (UpgradeNSXNodeInfo UpgradeNSXNodeInfo) Error() string {
    return "com.vmware.vcenter.nsx.deployment.NSX_node_info"
}



// The ``ClusterInfo`` class contains information about the cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpgradeClusterInfo struct {
    // Unique identifier of the given cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Id string
    // Name of the cluster as displayed in vSphere client. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DisplayName string
    // Current NSX version. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CurrentVersion string
    // The time when the cluster was last updated. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    LastUpdated *time.Time
    // The upgrade status of the cluster. Valid value: Not Started, Running, In Queue, Completed, Failed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Status Upgrade_ClusterStatus
    // The percent completion of the operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PercentComplete *float64
    // The list of Edge nodes in the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EdgeList []UpgradeNSXNodeInfo
    // The list of Host nodes in the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    HostList []UpgradeNSXNodeInfo
    // The precheck messages corresponding to the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PrecheckMessages UpgradeMessages
    // The upgrade messages corresponding to the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpgradeMessages UpgradeMessages
}



func (UpgradeClusterInfo UpgradeClusterInfo) Error() string {
    return "com.vmware.vcenter.nsx.deployment.cluster_info"
}



// The ``DeploymentInfo`` class contains information about the complete NSX deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpgradeDeploymentInfo struct {
    // The list of clusters availanble in the deployment. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Clusters []UpgradeClusterInfo
    // The list of manager appliances. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Managers []UpgradeNSXNodeInfo
    // The overall upgrade status of the deployment. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Status Upgrade_DeploymentStatus
    // The precheck messages corresponding to the deployment. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PrecheckMessages UpgradeMessages
    // The upgrade messages corresponding to the deployment. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpgradeMessages UpgradeMessages
}



func (UpgradeDeploymentInfo UpgradeDeploymentInfo) Error() string {
    return "com.vmware.vcenter.nsx.deployment.deployment_info"
}



// The ``Release`` class contains information about the NSX Release. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpgradeRelease struct {
    // The version of the NSX release. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Version string
    // The display name of the NSX release. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // A URI pointing to the NSX upgrade release notes. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ReleaseNotes *url.URL
}



func (UpgradeRelease UpgradeRelease) Error() string {
    return "com.vmware.vcenter.nsx.deployment.release"
}



// The ``Job`` class contains information about the job to be performed. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpgradeJob struct {
    // The name of the job. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // The status of the job. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Status Upgrade_Status
    // The time when the job was started. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StartTime *time.Time
    // The time when the job ended. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EndTime *time.Time
    // The percent completion of the operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PercentComplete *float64
}



func (UpgradeJob UpgradeJob) Error() string {
    return "com.vmware.vcenter.nsx.deployment.job"
}



// The ``Operation`` class contains information about the operation to be performed. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpgradeOperation struct {
    // The name of the operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // The status of the operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Status Upgrade_Status
    // The time when the operation was started. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StartTime *time.Time
    // The time when the operation ended. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EndTime *time.Time
    // The percent completion of the operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PercentComplete *float64
    // The list of jobs associated with the current operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Jobs []UpgradeJob
}



func (UpgradeOperation UpgradeOperation) Error() string {
    return "com.vmware.vcenter.nsx.deployment.operation"
}



// The ``Info`` class contains information about NSX deployment on the environment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpgradeInfo struct {
    // The version of the NSX deployed with the current vCenter server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CurrentVersion string
    // Available target version of the NSX. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AvailableRelease *UpgradeRelease
    // Available current operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CurrentOperation *UpgradeOperation
    // Information about the deployment, includng cluster and managers. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Deployment *UpgradeDeploymentInfo
}



func (UpgradeInfo UpgradeInfo) Error() string {
    return "com.vmware.vcenter.nsx.deployment.info"
}






func upgradeGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(UpgradeInfoBindingType)
}

func upgradeGetRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/nsx/deployment/upgrade",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"Error": 500})
}


func upgradeCheckInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeCheckOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func upgradeCheckRestMetadata() protocol.OperationRestMetadata {
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
      "POST",
      "/vcenter/nsx/deployment/upgrade",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"Error": 500})
}


func upgradeStartInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.execution_mode", reflect.TypeOf(Upgrade_ExecutionMode(Upgrade_ExecutionMode_BATCH)))
    fieldNameMap["mode"] = "Mode"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeStartOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func upgradeStartRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["mode"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.execution_mode", reflect.TypeOf(Upgrade_ExecutionMode(Upgrade_ExecutionMode_BATCH)))
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "mode",
      "POST",
      "/vcenter/nsx/deployment/upgrade",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"AlreadyInDesiredState": 400,"Error": 500})
}


func upgradePauseInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradePauseOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func upgradePauseRestMetadata() protocol.OperationRestMetadata {
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
      "POST",
      "/vcenter/nsx/deployment/upgrade",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"AlreadyInDesiredState": 400,"Error": 500})
}


func upgradeResumeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeResumeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func upgradeResumeRestMetadata() protocol.OperationRestMetadata {
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
      "POST",
      "/vcenter/nsx/deployment/upgrade",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"Error": 500})
}



func UpgradeMessagesBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["info"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["info"] = "Info"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.messages",fields, reflect.TypeOf(UpgradeMessages{}), fieldNameMap, validators)
}

func UpgradeNSXNodeInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.deployment.node"}, "")
    fieldNameMap["id"] = "Id"
    fields["node_type"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.NSX_node_type", reflect.TypeOf(Upgrade_NSXNodeType(Upgrade_NSXNodeType_MANAGER)))
    fieldNameMap["node_type"] = "NodeType"
    fields["percent_complete"] = bindings.NewOptionalType(bindings.NewDoubleType())
    fieldNameMap["percent_complete"] = "PercentComplete"
    fields["precheck_messages"] = bindings.NewReferenceType(UpgradeMessagesBindingType)
    fieldNameMap["precheck_messages"] = "PrecheckMessages"
    fields["upgrade_messages"] = bindings.NewReferenceType(UpgradeMessagesBindingType)
    fieldNameMap["upgrade_messages"] = "UpgradeMessages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.NSX_node_info",fields, reflect.TypeOf(UpgradeNSXNodeInfo{}), fieldNameMap, validators)
}

func UpgradeClusterInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["id"] = "Id"
    fields["display_name"] = bindings.NewStringType()
    fieldNameMap["display_name"] = "DisplayName"
    fields["current_version"] = bindings.NewStringType()
    fieldNameMap["current_version"] = "CurrentVersion"
    fields["last_updated"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["last_updated"] = "LastUpdated"
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.cluster_status", reflect.TypeOf(Upgrade_ClusterStatus(Upgrade_ClusterStatus_NOT_STARTED)))
    fieldNameMap["status"] = "Status"
    fields["percent_complete"] = bindings.NewOptionalType(bindings.NewDoubleType())
    fieldNameMap["percent_complete"] = "PercentComplete"
    fields["edge_list"] = bindings.NewListType(bindings.NewReferenceType(UpgradeNSXNodeInfoBindingType), reflect.TypeOf([]UpgradeNSXNodeInfo{}))
    fieldNameMap["edge_list"] = "EdgeList"
    fields["host_list"] = bindings.NewListType(bindings.NewReferenceType(UpgradeNSXNodeInfoBindingType), reflect.TypeOf([]UpgradeNSXNodeInfo{}))
    fieldNameMap["host_list"] = "HostList"
    fields["precheck_messages"] = bindings.NewReferenceType(UpgradeMessagesBindingType)
    fieldNameMap["precheck_messages"] = "PrecheckMessages"
    fields["upgrade_messages"] = bindings.NewReferenceType(UpgradeMessagesBindingType)
    fieldNameMap["upgrade_messages"] = "UpgradeMessages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.cluster_info",fields, reflect.TypeOf(UpgradeClusterInfo{}), fieldNameMap, validators)
}

func UpgradeDeploymentInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["clusters"] = bindings.NewListType(bindings.NewReferenceType(UpgradeClusterInfoBindingType), reflect.TypeOf([]UpgradeClusterInfo{}))
    fieldNameMap["clusters"] = "Clusters"
    fields["managers"] = bindings.NewListType(bindings.NewReferenceType(UpgradeNSXNodeInfoBindingType), reflect.TypeOf([]UpgradeNSXNodeInfo{}))
    fieldNameMap["managers"] = "Managers"
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.deployment_status", reflect.TypeOf(Upgrade_DeploymentStatus(Upgrade_DeploymentStatus_NOT_STARTED)))
    fieldNameMap["status"] = "Status"
    fields["precheck_messages"] = bindings.NewReferenceType(UpgradeMessagesBindingType)
    fieldNameMap["precheck_messages"] = "PrecheckMessages"
    fields["upgrade_messages"] = bindings.NewReferenceType(UpgradeMessagesBindingType)
    fieldNameMap["upgrade_messages"] = "UpgradeMessages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.deployment_info",fields, reflect.TypeOf(UpgradeDeploymentInfo{}), fieldNameMap, validators)
}

func UpgradeReleaseBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["release_notes"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["release_notes"] = "ReleaseNotes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.release",fields, reflect.TypeOf(UpgradeRelease{}), fieldNameMap, validators)
}

func UpgradeJobBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.status", reflect.TypeOf(Upgrade_Status(Upgrade_Status_NOT_STARTED)))
    fieldNameMap["status"] = "Status"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["percent_complete"] = bindings.NewOptionalType(bindings.NewDoubleType())
    fieldNameMap["percent_complete"] = "PercentComplete"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.job",fields, reflect.TypeOf(UpgradeJob{}), fieldNameMap, validators)
}

func UpgradeOperationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.status", reflect.TypeOf(Upgrade_Status(Upgrade_Status_NOT_STARTED)))
    fieldNameMap["status"] = "Status"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["percent_complete"] = bindings.NewOptionalType(bindings.NewDoubleType())
    fieldNameMap["percent_complete"] = "PercentComplete"
    fields["jobs"] = bindings.NewListType(bindings.NewReferenceType(UpgradeJobBindingType), reflect.TypeOf([]UpgradeJob{}))
    fieldNameMap["jobs"] = "Jobs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.operation",fields, reflect.TypeOf(UpgradeOperation{}), fieldNameMap, validators)
}

func UpgradeInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["current_version"] = bindings.NewStringType()
    fieldNameMap["current_version"] = "CurrentVersion"
    fields["available_release"] = bindings.NewOptionalType(bindings.NewReferenceType(UpgradeReleaseBindingType))
    fieldNameMap["available_release"] = "AvailableRelease"
    fields["current_operation"] = bindings.NewOptionalType(bindings.NewReferenceType(UpgradeOperationBindingType))
    fieldNameMap["current_operation"] = "CurrentOperation"
    fields["deployment"] = bindings.NewOptionalType(bindings.NewReferenceType(UpgradeDeploymentInfoBindingType))
    fieldNameMap["deployment"] = "Deployment"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.info",fields, reflect.TypeOf(UpgradeInfo{}), fieldNameMap, validators)
}


