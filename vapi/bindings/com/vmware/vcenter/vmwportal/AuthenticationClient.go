/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Authentication
 * Used by client-side stubs.
 */

package vmwportal

import (
)

// The ``Authentication`` interface allows the admin to login to my.vmware.com using the my.vmware.com account. This API is prerequisite for APIs that download binaries from my.vmware.com such as the NSX stage download API. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AuthenticationClient interface {


    // Create API authenticates the user on the my.vmware.com portal using the username password passed in the spec. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam Specification providing the my.vmware.com username and password.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
    Create(specParam AuthenticationCreateSpec) error 

}
