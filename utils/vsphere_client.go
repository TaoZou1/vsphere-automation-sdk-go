package utils

import (
	"crypto/tls"
	"net/http"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/security"
)

func NewVsphereConnector(server, username, password string) (client.Connector, error) {

	securityCtx := security.NewUserPasswordSecurityContext(username, password)

	// TODO remove this for production
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := http.Client{
		Transport: tr,
	}

	connector := client.NewJsonRpcConnector(server, httpClient)
	connector.SetSecurityContext(securityCtx)

	return connector, nil
}
