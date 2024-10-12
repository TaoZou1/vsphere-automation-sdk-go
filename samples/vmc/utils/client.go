// Copyright © 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause
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

	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/security"
)

const ServerUrl = "https://stg.skyscraper.vmware.com"
const TokenUrl = "https://console-stg.cloud.vmware.com"

func NewVmcConnector(refreshToken, serverUrl, cspURL string) (client.Connector, error) {

	var jsondata map[string]interface{}
	cspRefreshURLSuffix := "/csp/gateway/am/api/auth/api-tokens/authorize"
	var accessToken string

	if serverUrl == "" {
		serverUrl = "https://stg.skyscraper.vmware.com"
	}

	if cspURL == "" {
		cspURL = "https://console-stg.cloud.vmware.com" + cspRefreshURLSuffix
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
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := http.Client{
		Transport: tr,
	}

	connector := client.NewRestConnector(serverUrl, httpClient)
	connector.SetSecurityContext(securityCtx)

	return connector, nil
}
