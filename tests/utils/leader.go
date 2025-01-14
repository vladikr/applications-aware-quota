/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2023 Red Hat, Inc.
 *
 */

package utils

import (
	"context"
	"encoding/json"
	"kubevirt.io/applications-aware-quota/pkg/aaq-operator/resources/utils"

	"kubevirt.io/client-go/kubecli"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

func GetLeader(virtClient kubecli.KubevirtClient, aaqNS string) string {
	controllerEndpoint, err := virtClient.CoreV1().Endpoints(aaqNS).Get(context.Background(), utils.ControllerPodName, v1.GetOptions{})
	if err != nil {
		return ""
	}
	var record resourcelock.LeaderElectionRecord
	if recordBytes, found := controllerEndpoint.Annotations[resourcelock.LeaderElectionRecordAnnotationKey]; found {
		err := json.Unmarshal([]byte(recordBytes), &record)
		if err != nil {
			return ""
		}
	} else {

		return ""
	}
	return record.HolderIdentity
}
