/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ServiceReferences
 * Used by client-side stubs.
 */

package infra

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type ServiceReferencesClient interface {

    // This API can be used to delete a service reference with the given service-reference-id.
    //
    // @param serviceReferenceIdParam Id of Service Reference (required)
    // @param cascadeParam Flag to cascade delete all children associated with service reference (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(serviceReferenceIdParam string, cascadeParam *bool) error

    // This API can be used to read service reference with the given service-reference-id.
    //
    // @param serviceReferenceIdParam Id of Service Reference (required)
    // @return com.vmware.nsx_policy.model.ServiceReference
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(serviceReferenceIdParam string) (model.ServiceReference, error)

    // List all the partner service references available for service insertion
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.ServiceReferenceListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ServiceReferenceListResult, error)

    // Create Service Reference representing the intent to consume a given 3rd party service.
    //
    // @param serviceReferenceIdParam Service reference id (required)
    // @param serviceReferenceParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(serviceReferenceIdParam string, serviceReferenceParam model.ServiceReference) error

    // Create Service Reference representing the intent to consume a given 3rd party service.
    //
    // @param serviceReferenceIdParam Service reference id (required)
    // @param serviceReferenceParam (required)
    // @return com.vmware.nsx_policy.model.ServiceReference
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(serviceReferenceIdParam string, serviceReferenceParam model.ServiceReference) (model.ServiceReference, error)
}
