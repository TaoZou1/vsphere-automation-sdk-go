package utils

import (
	"crypto/tls"
	"github.com/sirupsen/logrus"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/cis"
	"net/http"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/security"
)

func NewVsphereConnector(server, username, password string) (client.Connector, error) {

	server = server + "/api"

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

	sessionClient := cis.NewSessionClientImpl(connector)
	sessionID, sessionError := sessionClient.Create()
	logrus.Infof("Session %s created successfully", sessionID)
	if sessionError != nil {
		logrus.Errorf("Could not create session %s", sessionError)
	}

	connector.SetSecurityContext(security.NewSessionSecurityContext(sessionID))
	//TODO
	//figure out when to terminate the session
	//defer func() {
	//	err := sessionClient.Delete()
	//	if err!= nil {
	//		logrus.Error("Error deleting session")
	//	} else{
	//		logrus.Infof("Session %s deleted sucessfully", sessionID)
	//	}
	//}()

	return connector, nil
}
