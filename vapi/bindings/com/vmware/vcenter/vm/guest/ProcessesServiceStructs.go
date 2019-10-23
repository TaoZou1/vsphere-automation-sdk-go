/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Processes.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package guest

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "time"
)




// The ``CreateSpec`` {\\\\@term structure) describes the arguments to Processes#create. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ProcessesCreateSpec struct {
    // The absolute path to the program to start. 
    //
    //  For Linux guest operating systems, /bin/bash is used to start the program. 
    //
    //  For Solaris guest operating systems, if /bin/bash exists, its used to start the program, otherwise /bin/sh is used. If /bin/sh is used, then the process ID returned by Processes#create will be that of the shell used to start the program, rather than the program itself, due to the differences in how /bin/sh and /bin/bash work. This PID will still be usable for watching the process with Processes#list to find its exit code and elapsed time. 
    //
    //  For Windows, no shell is used. Using a simple batch file instead by prepending ``c:\windows\system32\cmd.exe /c`` will allow stdio redirection to work if passed in the ``arguments`` parameter.. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Path string
    // The arguments to the program. 
    //
    //  Characters which must be escaped to the shell should also be escaped in ``arguments``. 
    //
    //  In Linux and Solaris guest operating systems, stdio redirection arguments may be used. 
    //
    //  For Windows, stdio redirection can be added to the argments if ``path`` is prefixed with ``c:\windows\system32\cmd.exe /c``.. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Arguments *string
    // The absolute path of the working directory for the program to be run. VMware recommends explicitly setting the working directory for the program to be run. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    WorkingDirectory *string
    // A map of environment variables, specified using the guest OS rules (for example ``PATH, c:\bin;c:\windows\system32`` or ``LD_LIBRARY_PATH,/usr/lib:/lib``), to be set for the program being run. Note that these are not additions to the default environment variables; they define the complete set available to the program. If this value is empty, then the environment variables used are guest dependent defaults. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EnvironmentVariables map[string]string
    // Makes any program window start minimized in Windows operating systems. Returns an error if map with bool value for non-Windows guests. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StartMinimized *bool
}



func (ProcessesCreateSpec ProcessesCreateSpec) Error() string {
    return "com.vmware.vcenter.vm.guest.create_spec"
}



// The ``Info`` class describes the state of a guest process. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ProcessesInfo struct {
    // The process name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // The process ID. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Pid int64
    // The process owner. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Owner string
    // The full command line of the process. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Command string
    // The start time of the process. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Started time.Time
    // If the process was started using Processes#create then the process completion time will be available if queried within 5 minutes after it completes. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Finished *time.Time
    // If the process was started using Processes#create then the process exit code will be available if queried within 5 minutes after it completes. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ExitCode *int64
}



func (ProcessesInfo ProcessesInfo) Error() string {
    return "com.vmware.vcenter.vm.guest.info"
}






func processesCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["credentials"] = bindings.NewReferenceType(CredentialsBindingType)
    fields["spec"] = bindings.NewReferenceType(ProcessesCreateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["credentials"] = "Credentials"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func processesCreateOutputType() bindings.BindingType {
    return bindings.NewIntegerType()
}

func processesCreateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unsupported": 400,"NotAllowedInCurrentState": 400,"UnableToAllocateResource": 500,"ResourceBusy": 500,"Unauthenticated": 401,"Unauthorized": 403,"ServiceUnavailable": 503})
}


func processesGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["credentials"] = bindings.NewReferenceType(CredentialsBindingType)
    fields["pid"] = bindings.NewIntegerType()
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["credentials"] = "Credentials"
    fieldNameMap["pid"] = "Pid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func processesGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ProcessesInfoBindingType)
}

func processesGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unsupported": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"Unauthenticated": 401,"ServiceUnavailable": 503})
}


func processesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["credentials"] = bindings.NewReferenceType(CredentialsBindingType)
    fields["pids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{})))
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["credentials"] = "Credentials"
    fieldNameMap["pids"] = "Pids"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func processesListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ProcessesInfoBindingType), reflect.TypeOf([]ProcessesInfo{}))
}

func processesListRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unsupported": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"Unauthenticated": 401,"ServiceUnavailable": 503})
}


func processesDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["credentials"] = bindings.NewReferenceType(CredentialsBindingType)
    fields["pid"] = bindings.NewIntegerType()
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["credentials"] = "Credentials"
    fieldNameMap["pid"] = "Pid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func processesDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func processesDeleteRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unsupported": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"Unauthenticated": 401,"Unauthorized": 403,"ServiceUnavailable": 503})
}



func ProcessesCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["path"] = bindings.NewStringType()
    fieldNameMap["path"] = "Path"
    fields["arguments"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["arguments"] = "Arguments"
    fields["working_directory"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["working_directory"] = "WorkingDirectory"
    fields["environment_variables"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{}))
    fieldNameMap["environment_variables"] = "EnvironmentVariables"
    fields["start_minimized"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["start_minimized"] = "StartMinimized"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.processes.create_spec",fields, reflect.TypeOf(ProcessesCreateSpec{}), fieldNameMap, validators)
}

func ProcessesInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["pid"] = bindings.NewIntegerType()
    fieldNameMap["pid"] = "Pid"
    fields["owner"] = bindings.NewStringType()
    fieldNameMap["owner"] = "Owner"
    fields["command"] = bindings.NewStringType()
    fieldNameMap["command"] = "Command"
    fields["started"] = bindings.NewDateTimeType()
    fieldNameMap["started"] = "Started"
    fields["finished"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["finished"] = "Finished"
    fields["exit_code"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["exit_code"] = "ExitCode"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.processes.info",fields, reflect.TypeOf(ProcessesInfo{}), fieldNameMap, validators)
}


