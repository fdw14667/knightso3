/*
 *  Copyright 2020-2021 Huawei Technologies Co., Ltd.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

// server address url config
package config

import (
	"errors"
	"mep-agent/src/util"
	"os"
	"strings"
)

type ServerUrl struct {
	MepServerRegisterUrl   string
	MepAuthUrl             string
	MepHeartBeatUrl        string
	MepServiceDiscoveryUrl string
}

const (
	MepAuthApigwUrl           string = "https://${MEP_IP}:${MEP_APIGW_PORT}/mep/token"
	MepSerRegisterApigwUrl    string = "https://${MEP_IP}:${MEP_APIGW_PORT}/mep/mec_service_mgmt/v1/applications/${appInstanceId}/services"
	MepSerQueryByNameApigwUrl string = "https://${MEP_IP}:${MEP_APIGW_PORT}/mep/mec_service_mgmt/v1/services?ser_name="
	MepHeartBeatApigwUrl      string = "https://${MEP_IP}:${MEP_APIGW_PORT}"
	MepIp                     string = "${MEP_IP}"
	MepApigwPort              string = "${MEP_APIGW_PORT}"
)

var ServerUrlConfig ServerUrl

// Returns server URL
func GetServerUrl() (ServerUrl, error) {

	var serverUrl ServerUrl
	// validate the env params
	mepIp := os.Getenv("MEP_IP")
	if util.ValidateDns(mepIp) != nil {
		return serverUrl, errors.New("validate MEP_IP failed")
	}
	mepApiGwPort := os.Getenv("MEP_APIGW_PORT")
	if len(mepApiGwPort) == 0 || util.ValidateByPattern(util.PORT_PATTERN, mepApiGwPort) != nil {
		return serverUrl, errors.New("validate MEP_APIGW_PORT failed")
	}

	serverUrl.MepServerRegisterUrl = strings.Replace(
		strings.Replace(MepSerRegisterApigwUrl, MepIp, mepIp, 1),
		MepApigwPort, mepApiGwPort, 1)

	serverUrl.MepAuthUrl = strings.Replace(
		strings.Replace(MepAuthApigwUrl, MepIp, mepIp, 1),
		MepApigwPort, mepApiGwPort, 1)

	serverUrl.MepHeartBeatUrl = strings.Replace(
		strings.Replace(MepHeartBeatApigwUrl, MepIp, mepIp, 1),
		MepApigwPort, mepApiGwPort, 1)

	serverUrl.MepServiceDiscoveryUrl = strings.Replace(
		strings.Replace(MepSerQueryByNameApigwUrl, MepIp, mepIp, 1),
		MepApigwPort, mepApiGwPort, 1)
	return serverUrl, nil
}
