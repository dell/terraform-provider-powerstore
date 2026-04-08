/*
 *
 * Copyright © 2020-2024 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientOptions_Insecure(t *testing.T) {
	co := NewClientOptions()
	co.SetInsecure(true)
	assert.Equal(t, true, co.Insecure())
}

func TestClientOptions_DefaultTimeout(t *testing.T) {
	co := NewClientOptions()
	value := int64(120)
	co.SetDefaultTimeout(value)
	assert.Equal(t, value, co.DefaultTimeout())
}

func TestClientOptions_RequestIDKey(t *testing.T) {
	co := NewClientOptions()
	co.SetRequestIDKey("foobar")
	assert.Equal(t, "foobar", string(co.RequestIDKey()))
}

func TestClientOptions_RateLimit(t *testing.T) {
	co := NewClientOptions()
	value := 10
	co.SetRateLimit(value)
	assert.Equal(t, value, co.RateLimit())
}
