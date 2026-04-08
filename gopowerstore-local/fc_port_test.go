/*
 *
 * Copyright © 2020 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package gopowerstore

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	fcPortMockURL = APIMockURL + apiFCPortURL
)

var (
	fcPortID  = "9197cba085b04a86b63a17267b11d1e6"
	fcPortID2 = "5c8af4bd7e4542418b9a5ec857912a28"
)

func TestClientIMPL_GetFCPorts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, fcPortID, fcPortID2)
	httpmock.RegisterResponder("GET", fcPortMockURL,
		httpmock.NewStringResponder(200, respData))
	ports, err := C.GetFCPorts(context.Background())
	assert.Nil(t, err)
	assert.Len(t, ports, 2)
	assert.Equal(t, fcPortID, ports[0].ID)
}

func TestClientIMPL_GetFCPort(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, fcPortID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", fcPortMockURL, fcPortID),
		httpmock.NewStringResponder(200, respData))
	port, err := C.GetFCPort(context.Background(), fcPortID)
	assert.Nil(t, err)
	assert.Equal(t, fcPortID, port.ID)
}
