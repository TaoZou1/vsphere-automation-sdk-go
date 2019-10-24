/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Providers
 * Used by client-side stubs.
 */
package orgs

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vmc/model"
)

type ProvidersClient interface {

    // @param orgParam Organization identifier. (required)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Organization doesn't exist
	List(orgParam string) ([]model.AwsCloudProvider, error)
}
