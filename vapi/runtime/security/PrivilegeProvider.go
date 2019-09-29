package security

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/
import "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"

type PrivilegeProvider interface {
	GetPrivilegeInfo(fullyQualifiedOperName string, inputValue data.DataValue) (map[ResourceIdentifier][]string, error)
}
