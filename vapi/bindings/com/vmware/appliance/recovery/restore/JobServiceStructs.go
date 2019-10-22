/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Job.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package restore

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``ReturnStatus`` enumeration class Defines the state of precheck
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Job_ReturnStatus string

const (
    // Check failed
     Job_ReturnStatus_FAIL Job_ReturnStatus = "FAIL"
    // Passed with warnings
     Job_ReturnStatus_WARNING Job_ReturnStatus = "WARNING"
    // Check passed
     Job_ReturnStatus_OK Job_ReturnStatus = "OK"
)

func (r Job_ReturnStatus) Job_ReturnStatus() bool {
    switch r {
        case Job_ReturnStatus_FAIL:
            return true
        case Job_ReturnStatus_WARNING:
            return true
        case Job_ReturnStatus_OK:
            return true
        default:
            return false
    }
}




// ``LocationType`` enumeration class Defines type of all locations for backup/restore
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Job_LocationType string

const (
    // Destination is FTP server
     Job_LocationType_FTP Job_LocationType = "FTP"
    // Destination is HTTP server
     Job_LocationType_HTTP Job_LocationType = "HTTP"
    // Destination is FTPS server
     Job_LocationType_FTPS Job_LocationType = "FTPS"
    // Destination is HTTPS server
     Job_LocationType_HTTPS Job_LocationType = "HTTPS"
    // Destination is SSH server
     Job_LocationType_SCP Job_LocationType = "SCP"
    // Destination is SFTP server
     Job_LocationType_SFTP Job_LocationType = "SFTP"
    // Destination is NFS server. This constant field was added in vSphere API 6.7.2.
     Job_LocationType_NFS Job_LocationType = "NFS"
    // Destination is SMB server. This constant field was added in vSphere API 6.7.2.
     Job_LocationType_SMB Job_LocationType = "SMB"
)

func (l Job_LocationType) Job_LocationType() bool {
    switch l {
        case Job_LocationType_FTP:
            return true
        case Job_LocationType_HTTP:
            return true
        case Job_LocationType_FTPS:
            return true
        case Job_LocationType_HTTPS:
            return true
        case Job_LocationType_SCP:
            return true
        case Job_LocationType_SFTP:
            return true
        case Job_LocationType_NFS:
            return true
        case Job_LocationType_SMB:
            return true
        default:
            return false
    }
}




// ``BackupRestoreProcessState`` enumeration class Defines state of backup/restore process
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Job_BackupRestoreProcessState string

const (
    // Failed
     Job_BackupRestoreProcessState_FAILED Job_BackupRestoreProcessState = "FAILED"
    // In progress
     Job_BackupRestoreProcessState_INPROGRESS Job_BackupRestoreProcessState = "INPROGRESS"
    // Not started
     Job_BackupRestoreProcessState_NONE Job_BackupRestoreProcessState = "NONE"
    // Completed successfully
     Job_BackupRestoreProcessState_SUCCEEDED Job_BackupRestoreProcessState = "SUCCEEDED"
)

func (b Job_BackupRestoreProcessState) Job_BackupRestoreProcessState() bool {
    switch b {
        case Job_BackupRestoreProcessState_FAILED:
            return true
        case Job_BackupRestoreProcessState_INPROGRESS:
            return true
        case Job_BackupRestoreProcessState_NONE:
            return true
        case Job_BackupRestoreProcessState_SUCCEEDED:
            return true
        default:
            return false
    }
}





// ``LocalizableMessage`` class Structure representing message
 type JobLocalizableMessage struct {
    // id in message bundle
    Id string
    // text in english
    DefaultMessage string
    // nested data
    Args []string
}



//


// ``ReturnResult`` class Structure representing precheck result
 type JobReturnResult struct {
    // Check status
    Status Job_ReturnStatus
    // List of messages
    Messages []JobLocalizableMessage
}



//


// ``RestoreRequest`` class Structure representing requested restore piece
 type JobRestoreRequest struct {
    // a password for a backup piece
    BackupPassword *string
    // a type of location
    LocationType Job_LocationType
    // path or url
    Location string
    // username for location
    LocationUser *string
    // password for location
    LocationPassword *string
    // Administrators Username for SSO. This property was added in vSphere API 6.7.
    SsoAdminUserName *string
    // The password for SSO admin user. This property was added in vSphere API 6.7.
    SsoAdminUserPassword *string
    // The flag to ignore warnings during restore. This property was added in vSphere API 6.7.
    IgnoreWarnings *bool
}



//


// ``RestoreJobStatus`` class Structure representing backup restore status
 type JobRestoreJobStatus struct {
    // process state
    State Job_BackupRestoreProcessState
    // list of messages
    Messages []JobLocalizableMessage
    // percentage complete
    Progress int64
}



//





func jobCancelInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func jobCancelOutputType() bindings.BindingType {
    return bindings.NewReferenceType(JobReturnResultBindingType)
}

func jobCancelRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
}


func jobCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["piece"] = bindings.NewReferenceType(JobRestoreRequestBindingType)
    fieldNameMap["piece"] = "Piece"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func jobCreateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(JobRestoreJobStatusBindingType)
}

func jobCreateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"FeatureInUse": 400,"NotAllowedInCurrentState": 400,"Error": 500})
}


func jobGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func jobGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(JobRestoreJobStatusBindingType)
}

func jobGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
}



func JobLocalizableMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["default_message"] = bindings.NewStringType()
    fieldNameMap["default_message"] = "DefaultMessage"
    fields["args"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["args"] = "Args"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.restore.job.localizable_message",fields, reflect.TypeOf(JobLocalizableMessage{}), fieldNameMap, validators)
}

func JobReturnResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.recovery.restore.job.return_status", reflect.TypeOf(Job_ReturnStatus(Job_ReturnStatus_FAIL)))
    fieldNameMap["status"] = "Status"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(JobLocalizableMessageBindingType), reflect.TypeOf([]JobLocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.restore.job.return_result",fields, reflect.TypeOf(JobReturnResult{}), fieldNameMap, validators)
}

func JobRestoreRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["backup_password"] = "BackupPassword"
    fields["location_type"] = bindings.NewEnumType("com.vmware.appliance.recovery.restore.job.location_type", reflect.TypeOf(Job_LocationType(Job_LocationType_FTP)))
    fieldNameMap["location_type"] = "LocationType"
    fields["location"] = bindings.NewStringType()
    fieldNameMap["location"] = "Location"
    fields["location_user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["location_user"] = "LocationUser"
    fields["location_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["location_password"] = "LocationPassword"
    fields["sso_admin_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["sso_admin_user_name"] = "SsoAdminUserName"
    fields["sso_admin_user_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["sso_admin_user_password"] = "SsoAdminUserPassword"
    fields["ignore_warnings"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ignore_warnings"] = "IgnoreWarnings"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.restore.job.restore_request",fields, reflect.TypeOf(JobRestoreRequest{}), fieldNameMap, validators)
}

func JobRestoreJobStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewEnumType("com.vmware.appliance.recovery.restore.job.backup_restore_process_state", reflect.TypeOf(Job_BackupRestoreProcessState(Job_BackupRestoreProcessState_FAILED)))
    fieldNameMap["state"] = "State"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(JobLocalizableMessageBindingType), reflect.TypeOf([]JobLocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["progress"] = bindings.NewIntegerType()
    fieldNameMap["progress"] = "Progress"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.restore.job.restore_job_status",fields, reflect.TypeOf(JobRestoreJobStatus{}), fieldNameMap, validators)
}


