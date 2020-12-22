/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: AlbNetworkSecurityPolicies
 * Used by client-side stubs.
 */

package global_infra

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type AlbNetworkSecurityPoliciesClient interface {

    // Delete the ALBNetworkSecurityPolicy along with all the entities contained by this ALBNetworkSecurityPolicy.
    //
    // @param albNetworksecuritypolicyIdParam ALBNetworkSecurityPolicy ID (required)
    // @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(albNetworksecuritypolicyIdParam string, forceParam *bool) error

    // Read a ALBNetworkSecurityPolicy.
    //
    // @param albNetworksecuritypolicyIdParam ALBNetworkSecurityPolicy ID (required)
    // @return com.vmware.nsx_policy.model.ALBNetworkSecurityPolicy
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(albNetworksecuritypolicyIdParam string) (model.ALBNetworkSecurityPolicy, error)

    // Paginated list of all ALBNetworkSecurityPolicy for infra.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.ALBNetworkSecurityPolicyApiResponse
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ALBNetworkSecurityPolicyApiResponse, error)

    // If a ALBnetworksecuritypolicy with the alb-networksecuritypolicy-id is not already present, create a new ALBnetworksecuritypolicy. If it already exists, update the ALBnetworksecuritypolicy. This is a full replace.
    //
    // @param albNetworksecuritypolicyIdParam ALBnetworksecuritypolicy ID (required)
    // @param aLBNetworkSecurityPolicyParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(albNetworksecuritypolicyIdParam string, aLBNetworkSecurityPolicyParam model.ALBNetworkSecurityPolicy) error

    // If a ALBNetworkSecurityPolicy with the alb-NetworkSecurityPolicy-id is not already present, create a new ALBNetworkSecurityPolicy. If it already exists, update the ALBNetworkSecurityPolicy. This is a full replace.
    //
    // @param albNetworksecuritypolicyIdParam ALBNetworkSecurityPolicy ID (required)
    // @param aLBNetworkSecurityPolicyParam (required)
    // @return com.vmware.nsx_policy.model.ALBNetworkSecurityPolicy
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(albNetworksecuritypolicyIdParam string, aLBNetworkSecurityPolicyParam model.ALBNetworkSecurityPolicy) (model.ALBNetworkSecurityPolicy, error)
}
