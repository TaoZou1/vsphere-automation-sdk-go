/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Backup.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package recovery

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``ReturnStatus`` enumeration class Defines the state of precheck
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Backup_ReturnStatus string

const (
    // Check failed
     Backup_ReturnStatus_FAIL Backup_ReturnStatus = "FAIL"
    // Passed with warnings
     Backup_ReturnStatus_WARNING Backup_ReturnStatus = "WARNING"
    // Check passed
     Backup_ReturnStatus_OK Backup_ReturnStatus = "OK"
)

func (r Backup_ReturnStatus) Backup_ReturnStatus() bool {
    switch r {
        case Backup_ReturnStatus_FAIL:
            return true
        case Backup_ReturnStatus_WARNING:
            return true
        case Backup_ReturnStatus_OK:
            return true
        default:
            return false
    }
}




// ``LocationType`` enumeration class Defines type of all locations for backup/restore
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Backup_LocationType string

const (
    // Destination is FTP server
     Backup_LocationType_FTP Backup_LocationType = "FTP"
    // Destination is HTTP server
     Backup_LocationType_HTTP Backup_LocationType = "HTTP"
    // Destination is FTPS server
     Backup_LocationType_FTPS Backup_LocationType = "FTPS"
    // Destination is HTTPS server
     Backup_LocationType_HTTPS Backup_LocationType = "HTTPS"
    // Destination is SSH server
     Backup_LocationType_SCP Backup_LocationType = "SCP"
    // Destination is SFTP server
     Backup_LocationType_SFTP Backup_LocationType = "SFTP"
    // Destination is NFS server
     Backup_LocationType_NFS Backup_LocationType = "NFS"
    // Destination is SMB server
     Backup_LocationType_SMB Backup_LocationType = "SMB"
)

func (l Backup_LocationType) Backup_LocationType() bool {
    switch l {
        case Backup_LocationType_FTP:
            return true
        case Backup_LocationType_HTTP:
            return true
        case Backup_LocationType_FTPS:
            return true
        case Backup_LocationType_HTTPS:
            return true
        case Backup_LocationType_SCP:
            return true
        case Backup_LocationType_SFTP:
            return true
        case Backup_LocationType_NFS:
            return true
        case Backup_LocationType_SMB:
            return true
        default:
            return false
    }
}





// ``LocalizableMessage`` class Structure representing message
 type BackupLocalizableMessage struct {
    // id in message bundle
    Id string
    // text in english
    DefaultMessage string
    // nested data
    Args []string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``ReturnResult`` class Structure representing precheck result
 type BackupReturnResult struct {
    // Check status
    Status Backup_ReturnStatus
    // List of messages
    Messages []BackupLocalizableMessage
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``BackupRequest`` class Structure representing requested backup piece
 type BackupBackupRequest struct {
    // a list of optional parts. Run backup parts APIs to get list of optional parts and description
    Parts []string
    // a password for a backup piece The backupPassword must adhere to the following password requirements: At least 8 characters, cannot be more than 20 characters in length. At least 1 uppercase letter. At least 1 lowercase letter. At least 1 numeric digit. At least 1 special character (i.e. any character not in [0-9,a-z,A-Z]). Only visible ASCII characters (for example, no space).
    BackupPassword *string
    // a type of location
    LocationType Backup_LocationType
    // path or url
    Location string
    // username for location
    LocationUser *string
    // password for location
    LocationPassword *string
    // Custom comment
    Comment *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func backupValidateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["piece"] = bindings.NewReferenceType(BackupBackupRequestBindingType)
    fieldNameMap["piece"] = "Piece"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func backupValidateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(BackupReturnResultBindingType)
}

func backupValidateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"FeatureInUse": 400,"Error": 500})
}



func BackupLocalizableMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["default_message"] = bindings.NewStringType()
    fieldNameMap["default_message"] = "DefaultMessage"
    fields["args"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["args"] = "Args"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.localizable_message",fields, reflect.TypeOf(BackupLocalizableMessage{}), fieldNameMap, validators)
}

func BackupReturnResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.return_status", reflect.TypeOf(Backup_ReturnStatus(Backup_ReturnStatus_FAIL)))
    fieldNameMap["status"] = "Status"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(BackupLocalizableMessageBindingType), reflect.TypeOf([]BackupLocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.return_result",fields, reflect.TypeOf(BackupReturnResult{}), fieldNameMap, validators)
}

func BackupBackupRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["parts"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["parts"] = "Parts"
    fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["backup_password"] = "BackupPassword"
    fields["location_type"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.location_type", reflect.TypeOf(Backup_LocationType(Backup_LocationType_FTP)))
    fieldNameMap["location_type"] = "LocationType"
    fields["location"] = bindings.NewStringType()
    fieldNameMap["location"] = "Location"
    fields["location_user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["location_user"] = "LocationUser"
    fields["location_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["location_password"] = "LocationPassword"
    fields["comment"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["comment"] = "Comment"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.backup_request",fields, reflect.TypeOf(BackupBackupRequest{}), fieldNameMap, validators)
}


