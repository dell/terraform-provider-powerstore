/*
 *
 * Copyright © 2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	softwareInstalledMockURL = APIMockURL + apiSoftwareInstalledURL
)

var (
	softwareInstalledID1 = "043294b6-b9b9-4adf-9e94-c227d24e8e4e"
	softwareInstalledID2 = "1add59ae-8302-4e2d-842f-337bd4c60e6c"
)

func TestClientIMPL_GetSoftwareInstalled(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, softwareInstalledID1, softwareInstalledID2)
	httpmock.RegisterResponder("GET", softwareInstalledMockURL,
		httpmock.NewStringResponder(200, respData))
	softwareInstalled, err := C.GetSoftwareInstalled(context.Background())
	assert.Nil(t, err)
	assert.Len(t, softwareInstalled, 2)
	assert.Equal(t, softwareInstalledID1, softwareInstalled[0].ID)
}
