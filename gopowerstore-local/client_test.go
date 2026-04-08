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
	"os"
	"testing"

	"github.com/dell/gopowerstore/api"
	"github.com/stretchr/testify/assert"
)

var C Client

const APIMockURL = "https://mock-server/api/rest/"

func initClient() {
	clientOptions := NewClientOptions()
	C, _ = NewClientWithArgs(
		APIMockURL,
		"admin",
		"Password",
		clientOptions)
}

func init() {
	initClient()
}

func TestNewClient(t *testing.T) {
	os.Setenv(InsecureEnv, "true")
	os.Setenv(APIURLEnv, "api")
	os.Setenv(UsernameEnv, "admin")
	os.Setenv(PasswordEnv, "password")
	os.Setenv(HTTPTimeoutEnv, "120")
	_, err := NewClient()
	assert.Nil(t, err)
	os.Unsetenv(UsernameEnv)
	_, err = NewClient()
	assert.NotNil(t, err)
}

func TestClientIMPL_SetTraceID(t *testing.T) {
	ctx := context.Background()
	ctx = C.SetTraceID(ctx, "123")
	assert.Equal(t, "123", ctx.Value(api.ContextKey(clientOptionsDefaultRequestIDKey)))
}
