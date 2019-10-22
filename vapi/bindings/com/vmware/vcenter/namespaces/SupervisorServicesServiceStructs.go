/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: SupervisorServices.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package namespaces

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``SetSpec`` class contains the specification required to set the desired state of a SupervisorService. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type SupervisorServicesSetSpec struct {
    // Desired state for service enablement. If the state is changed to be enabled, the service will be deployed. Typically this involves creation of a namespace, a Kubernetes operator, RBAC rules, Kubernetes Service objects, as well as registration of new CustomResourceDefinitions. Thus, as new pods may be scheduled, additional CPU, memory and possibly storage resources may be consumed. The amount depends on the specific service. If the state if changed to be disabled, the service is marked for deletion. Note though that disabling may block if the service is still using resources it doesn't want to release. Hence it is recommended to explicitly delete any instances before disabling the service. This is done to prevent accidental deletion of critical customer data. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Enabled bool
    // Desired version. Changing the version only relates to the version of the service operator, i.e. its control plane, not to any running instances, if applicable. Updating may lead to temporary service disruption of the operator. Available versions are published by the service, and only published versions are allowed to be set. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Version *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func supervisorServicesSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["service_ID"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.SupervisorService"}, "")
    fields["spec"] = bindings.NewReferenceType(SupervisorServicesSetSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["service_ID"] = "ServiceID"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func supervisorServicesSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func supervisorServicesSetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(SupervisorServicesSetSpecBindingType)
    paramsTypeMap["service_ID"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.SupervisorService"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["serviceID"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.SupervisorService"}, "")
    pathParams["cluster"] = "cluster"
    pathParams["service_ID"] = "serviceID"
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
      "/vcenter/namespace-management/clusters/{cluster}/supervisorservices/{serviceID}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func SupervisorServicesSetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["enabled"] = bindings.NewBooleanType()
    fieldNameMap["enabled"] = "Enabled"
    fields["version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.supervisor_services.set_spec",fields, reflect.TypeOf(SupervisorServicesSetSpec{}), fieldNameMap, validators)
}


