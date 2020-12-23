/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: AlbVsDataScriptSets
 * Used by client-side stubs.
 */

package infra

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type AlbVsDataScriptSetsClient interface {

    // Delete the ALBVSDataScriptSet along with all the entities contained by this ALBVSDataScriptSet.
    //
    // @param albVsdatascriptsetIdParam ALBVSDataScriptSet ID (required)
    // @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(albVsdatascriptsetIdParam string, forceParam *bool) error

    // Read a ALBVSDataScriptSet.
    //
    // @param albVsdatascriptsetIdParam ALBVSDataScriptSet ID (required)
    // @return com.vmware.nsx_policy.model.ALBVSDataScriptSet
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(albVsdatascriptsetIdParam string) (model.ALBVSDataScriptSet, error)

    // Paginated list of all ALBVSDataScriptSet for infra.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.ALBVSDataScriptSetApiResponse
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ALBVSDataScriptSetApiResponse, error)

    // If a ALBvsdatascriptset with the alb-vsdatascriptset-id is not already present, create a new ALBvsdatascriptset. If it already exists, update the ALBvsdatascriptset. This is a full replace.
    //
    // @param albVsdatascriptsetIdParam ALBvsdatascriptset ID (required)
    // @param aLBVSDataScriptSetParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(albVsdatascriptsetIdParam string, aLBVSDataScriptSetParam model.ALBVSDataScriptSet) error

    // If a ALBVSDataScriptSet with the alb-VSDataScriptSet-id is not already present, create a new ALBVSDataScriptSet. If it already exists, update the ALBVSDataScriptSet. This is a full replace.
    //
    // @param albVsdatascriptsetIdParam ALBVSDataScriptSet ID (required)
    // @param aLBVSDataScriptSetParam (required)
    // @return com.vmware.nsx_policy.model.ALBVSDataScriptSet
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(albVsdatascriptsetIdParam string, aLBVSDataScriptSetParam model.ALBVSDataScriptSet) (model.ALBVSDataScriptSet, error)
}
