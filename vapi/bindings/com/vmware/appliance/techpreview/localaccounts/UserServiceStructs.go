/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: User.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package localaccounts

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``UserAccountStatus`` enumeration class Defines status of user accounts
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type User_UserAccountStatus string

const (
    // The user account is disabled.
     User_UserAccountStatus_disabled User_UserAccountStatus = "disabled"
    // The user account is enabled.
     User_UserAccountStatus_enabled User_UserAccountStatus = "enabled"
)

func (u User_UserAccountStatus) User_UserAccountStatus() bool {
    switch u {
        case User_UserAccountStatus_disabled:
            return true
        case User_UserAccountStatus_enabled:
            return true
        default:
            return false
    }
}




// ``UserPasswordStatus`` enumeration class Defines state of user password
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type User_UserPasswordStatus string

const (
    // No password has been set
     User_UserPasswordStatus_notset User_UserPasswordStatus = "notset"
    // The password has expired.
     User_UserPasswordStatus_expired User_UserPasswordStatus = "expired"
    // The password is still valid.
     User_UserPasswordStatus_valid User_UserPasswordStatus = "valid"
)

func (u User_UserPasswordStatus) User_UserPasswordStatus() bool {
    switch u {
        case User_UserPasswordStatus_notset:
            return true
        case User_UserPasswordStatus_expired:
            return true
        case User_UserPasswordStatus_valid:
            return true
        default:
            return false
    }
}




// ``UserRole`` enumeration class Defines user roles for appliance
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type User_UserRole string

const (
    // Able to configure the appliance.
     User_UserRole_admin User_UserRole = "admin"
    // Able to read the appliance configuration.
     User_UserRole_operator User_UserRole = "operator"
    // Able to configure the appliance, manage local accounts and use the BASH shell
     User_UserRole_superAdmin User_UserRole = "superAdmin"
)

func (u User_UserRole) User_UserRole() bool {
    switch u {
        case User_UserRole_admin:
            return true
        case User_UserRole_operator:
            return true
        case User_UserRole_superAdmin:
            return true
        default:
            return false
    }
}





// ``UserConfigGet`` class Structure defines a user configuration for user.get API.
 type UserUserConfigGet struct {
    // User login name
    Username string
    // User roles
    Role User_UserRole
    // User full name
    Fullname string
    // Shows whether the user account is enabled or disabled.
    Status User_UserAccountStatus
    // Shows whether the user account is still valid or expired.
    Passwordstatus User_UserPasswordStatus
    // Email address of the local account.
    Email string
}



//


// ``UserConfig`` class Structure that defines a new user configuration for CLI.
 type UserUserConfig struct {
    // User login name
    Username string
    // User roles
    Role User_UserRole
    // User full name
    Fullname string
    // Enabled status of the local account
    Status User_UserAccountStatus
    // email of the local account
    Email string
}



//


// ``NewUserConfig`` class Structure that defines a new user configuration.
 type UserNewUserConfig struct {
    // User login name
    Username string
    // User roles. The default role is operator.
    Role *User_UserRole
    // User login password In Interactive mode, provide --password as part of the command, and enter the value on the prompt. When accessed remotely, provide --password value as part the command.
    Password string
    // User full name
    Fullname *string
    // Email address of the local account.
    Email *string
}



//





func userDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func userDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func userDeleteRestMetadata() protocol.OperationRestMetadata {
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


func userAddInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(UserNewUserConfigBindingType)
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func userAddOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func userAddRestMetadata() protocol.OperationRestMetadata {
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


func userSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(UserUserConfigBindingType)
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func userSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func userSetRestMetadata() protocol.OperationRestMetadata {
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


func userListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func userListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(UserUserConfigGetBindingType), reflect.TypeOf([]UserUserConfigGet{}))
}

func userListRestMetadata() protocol.OperationRestMetadata {
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


func userGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func userGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(UserUserConfigGetBindingType)
}

func userGetRestMetadata() protocol.OperationRestMetadata {
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



func UserUserConfigGetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["role"] = bindings.NewEnumType("com.vmware.appliance.techpreview.localaccounts.user.user_role", reflect.TypeOf(User_UserRole(User_UserRole_admin)))
    fieldNameMap["role"] = "Role"
    fields["fullname"] = bindings.NewStringType()
    fieldNameMap["fullname"] = "Fullname"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.techpreview.localaccounts.user.user_account_status", reflect.TypeOf(User_UserAccountStatus(User_UserAccountStatus_disabled)))
    fieldNameMap["status"] = "Status"
    fields["passwordstatus"] = bindings.NewEnumType("com.vmware.appliance.techpreview.localaccounts.user.user_password_status", reflect.TypeOf(User_UserPasswordStatus(User_UserPasswordStatus_notset)))
    fieldNameMap["passwordstatus"] = "Passwordstatus"
    fields["email"] = bindings.NewStringType()
    fieldNameMap["email"] = "Email"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.localaccounts.user.user_config_get",fields, reflect.TypeOf(UserUserConfigGet{}), fieldNameMap, validators)
}

func UserUserConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["role"] = bindings.NewEnumType("com.vmware.appliance.techpreview.localaccounts.user.user_role", reflect.TypeOf(User_UserRole(User_UserRole_admin)))
    fieldNameMap["role"] = "Role"
    fields["fullname"] = bindings.NewStringType()
    fieldNameMap["fullname"] = "Fullname"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.techpreview.localaccounts.user.user_account_status", reflect.TypeOf(User_UserAccountStatus(User_UserAccountStatus_disabled)))
    fieldNameMap["status"] = "Status"
    fields["email"] = bindings.NewStringType()
    fieldNameMap["email"] = "Email"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.localaccounts.user.user_config",fields, reflect.TypeOf(UserUserConfig{}), fieldNameMap, validators)
}

func UserNewUserConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["role"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.appliance.techpreview.localaccounts.user.user_role", reflect.TypeOf(User_UserRole(User_UserRole_admin))))
    fieldNameMap["role"] = "Role"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    fields["fullname"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["fullname"] = "Fullname"
    fields["email"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["email"] = "Email"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.localaccounts.user.new_user_config",fields, reflect.TypeOf(UserNewUserConfig{}), fieldNameMap, validators)
}


