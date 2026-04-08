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

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryParams_Select(t *testing.T) {
	qp := QueryParams{}
	qp.Select("foo", "bar")
	qp.Select("spam")
	assert.NotNil(t, qp.Select)
	assert.Contains(t, *(qp.selectParam), "foo")
	assert.Contains(t, *(qp.selectParam), "bar")
	assert.Contains(t, *(qp.selectParam), "spam")
}

func TestQueryParams_Order(t *testing.T) {
	qp := QueryParams{}
	qp.Order("foo", "bar")
	qp.Order("spam")
	assert.NotNil(t, qp.Order)
	assert.Contains(t, *(qp.orderParam), "foo")
	assert.Contains(t, *(qp.orderParam), "bar")
	assert.Contains(t, *(qp.orderParam), "spam")
}

func TestQueryParams_Encode_Empty(t *testing.T) {
	qp := QueryParams{}
	assert.Len(t, qp.Encode(), 0)
}

func TestQueryParams_Encode(t *testing.T) {
	qp := QueryParams{}
	qp.Async(false)
	qp.Limit(10)
	qp.Offset(20)
	qp.Select("foo", "bar")
	qp.Order("foo")
	qp.RawArg("spam", "ham,foobar")
	paramString := qp.Encode()
	assert.Contains(t, paramString, "limit=10")
	assert.Contains(t, paramString, "offset=20")
	assert.Contains(t, paramString, "async=false")
	assert.Contains(t, paramString, "order=foo")
	assert.Contains(t, paramString, "select=foo%2Cbar")
	assert.Contains(t, paramString, "spam=ham%2Cfoobar")
}

func TestQueryParamsOverride(t *testing.T) {
	qp := QueryParams{}
	qp.Async(false)
	qp.RawArg("async", "true")
	assert.Contains(t, qp.Encode(), "async=true")
}

func TestQueryParamsChaining(t *testing.T) {
	qp := QueryParams{}
	qp.Async(true).Select("foo", "bar")
	assert.Contains(t, qp.Encode(), "select=foo%2Cbar")
}
