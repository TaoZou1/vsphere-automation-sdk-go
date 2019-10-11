
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: MetadataIdentifier
 * Functions that implement the generated MetadataIdentifierClient interface
 */


package metamodel

import (
    "reflect"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
// //
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
// 
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)


type MetadataIdentifierClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewMetadataIdentifierClientImpl(connector client.Connector) *MetadataIdentifierClientImpl {
      interfaceName := "com.vmware.vapi.metadata.metamodel.metadata_identifier"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      mIface := MetadataIdentifierClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      mIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      return &mIface
}


func (mIface *MetadataIdentifierClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := mIface.connector.GetApiProvider().Invoke(mIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}


// Identifier representing the CanonicalName metadata.
const MetadataIdentifier_CANONICAL_NAME = "CanonicalName"
// Identifier representing the Component metadata.
const MetadataIdentifier_COMPONENT = "Component"
// Identifier representing the Create metadata.
const MetadataIdentifier_CREATE = "Create"
// Identifier representing the Crud metadata.
const MetadataIdentifier_CRUD = "Crud"
// Identifier representing the HasFieldsOf metadata.
const MetadataIdentifier_HAS_FIELDS_OF = "HasFieldsOf"
// Identifier representing the Includable metadata.
const MetadataIdentifier_INCLUDABLE = "Includable"
// Identifier representing the Include metadata.
const MetadataIdentifier_INCLUDE = "Include"
// Identifier representing the IsOneOf metadata.
const MetadataIdentifier_IS_ONE_OF = "IsOneOf"
// Identifier representing the Model metadata.
const MetadataIdentifier_MODEL = "Model"
// Identifier representing the Read metadata.
const MetadataIdentifier_READ = "Read"
// Identifier representing the Resource metadata.
const MetadataIdentifier_RESOURCE = "Resource"
// Identifier representing the UnionCase metadata.
const MetadataIdentifier_UNION_CASE = "UnionCase"
// Identifier representing the UnionTag metadata.
const MetadataIdentifier_UNION_TAG = "UnionTag"
// Identifier representing the Update metadata.
const MetadataIdentifier_UPDATE = "Update"








