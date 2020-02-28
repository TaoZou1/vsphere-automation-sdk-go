/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CpuMemThresholdsProfiles
 * Used by client-side stubs.
 */

package firewall

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type CpuMemThresholdsProfilesClient interface {

    // Delete CPU and memory thresholds profile.
    //
    // @param profileIdParam (required)
    // @param overrideParam Locally override the global object (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(profileIdParam string, overrideParam *bool) error

    // Read the CPU and memory thresholds profile.
    //
    // @param profileIdParam (required)
    // @return com.vmware.nsx_policy.model.PolicyFirewallCpuMemThresholdsProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(profileIdParam string) (model.PolicyFirewallCpuMemThresholdsProfile, error)

    // List all CPU and memory thresholds profiles.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.PolicyFirewallCpuMemThresholdsProfileListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyFirewallCpuMemThresholdsProfileListResult, error)

    // Create or update CPU and memory thresholds profile.
    //
    // @param profileIdParam (required)
    // @param policyFirewallCpuMemThresholdsProfileParam (required)
    // @param overrideParam Locally override the global object (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(profileIdParam string, policyFirewallCpuMemThresholdsProfileParam model.PolicyFirewallCpuMemThresholdsProfile, overrideParam *bool) error

    // Create or update CPU and memory thresholds profile.
    //
    // @param profileIdParam (required)
    // @param policyFirewallCpuMemThresholdsProfileParam (required)
    // @param overrideParam Locally override the global object (optional, default to false)
    // @return com.vmware.nsx_policy.model.PolicyFirewallCpuMemThresholdsProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(profileIdParam string, policyFirewallCpuMemThresholdsProfileParam model.PolicyFirewallCpuMemThresholdsProfile, overrideParam *bool) (model.PolicyFirewallCpuMemThresholdsProfile, error)
}