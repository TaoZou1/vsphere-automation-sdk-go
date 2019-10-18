package utils

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/security"
)

// type JWTToken struct {
// 	Id_token      string
// 	Token_type    string
// 	Expires_in    string
// 	Scope         string
// 	Access_token  string
// 	Refresh_token string
// }

func NewVmcConnector(refreshToken, vmcURL, cspURL string) (client.Connector, error) {

	var jsondata map[string]interface{}

	if vmcURL == "" {
		vmcURL = "https://vmc.vmware.com"
	}

	cspRefreshURLSuffix := "/csp/gateway/am/api/auth/api-tokens/authorize"

	if cspURL == "" {
		cspURL = "https://console.cloud.vmware.com" + cspRefreshURLSuffix
	} else {
		cspURL = cspURL + cspRefreshURLSuffix
	}

	payload := strings.NewReader("refresh_token=" + refreshToken)

	req, _ := http.NewRequest("POST", cspURL, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("status code %d trying to get csp auth token. %s", res.StatusCode, string(b))
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&jsondata)

	var accessToken string
	if token, ok := jsondata["access_token"]; ok {
		if accessTokenStr, ok := token.(string); ok {
			accessToken = accessTokenStr
		} else {
			errMsg := fmt.Sprintf("Invalid type for access_token, expected string, actual %s", reflect.TypeOf(token).String())
			return nil, errors.New(errMsg)
		}
	} else {
		return nil, errors.New("CSP auth response doesn't contain access token")
	}

	// log.Info("CSP Access token is " + accessToken)
	securityCtx := security.NewOauthSecurityContext(accessToken)

	// TODO remove this for production
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := http.Client{
		Transport: tr,
	}

	connector := client.NewRestConnector(vmcURL, httpClient)
	connector.SetSecurityContext(securityCtx)

	return connector, nil
}
