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

package inttests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFCPorts(t *testing.T) {
	ports, err := C.GetFCPorts(context.Background())
	checkAPIErr(t, err)
	assert.NotEmpty(t, ports)
	if len(ports) > 0 {
		port := ports[0]
		port, err = C.GetFCPort(context.Background(), port.ID)
		checkAPIErr(t, err)
		assert.NotEmpty(t, port.Wwn)
	}
}
