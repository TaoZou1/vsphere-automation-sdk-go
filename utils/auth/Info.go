/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package auth

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/utils/auth/model"
)

type info struct {
	authInterface model.AuthInfo
	authScheme    Scheme
}

func (ai *info) AuthScheme() Scheme {
	return ai.authScheme
}

func (ai *info) AuthInterface() interface{} {
	switch ai.authScheme {
	case BasicAuth, OAuthRefreshToken, SAMLBearer:
		return ai.authInterface.AuthInterface()
	default:
		return nil
	}
}

func (ai *info) SecurityContext() (core.SecurityContext, error) {
	switch ai.authScheme {
	case BasicAuth, OAuthRefreshToken, SAMLBearer:
		return ai.authInterface.SecurityContext()
	default:
		return nil, nil
	}
}
