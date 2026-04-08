/*
 *
 * Copyright © 2023-2024 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"net/http"
	"testing"

	"github.com/dell/gopowerstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetMaxVolumeSizeTestSuite struct {
	suite.Suite
	C gopowerstore.Client
}

func TestGetMaxVolumeSizeSuite(t *testing.T) {
	suite.Run(t, new(GetMaxVolumeSizeTestSuite))
}

func (s *GetMaxVolumeSizeTestSuite) SetupTest() {
	s.C = GetNewClient()
}

func (s *GetMaxVolumeSizeTestSuite) TestGetMaxVolumeSize() {
	customHeaders := s.C.GetCustomHTTPHeaders()
	if customHeaders == nil {
		customHeaders = make(http.Header)
	}
	customHeaders.Add("DELL-VISIBILITY", "internal")
	s.C.SetCustomHTTPHeaders(customHeaders)

	limit, err := s.C.GetMaxVolumeSize(context.Background())

	// reset custom header
	customHeaders.Del("DELL-VISIBILITY")
	s.C.SetCustomHTTPHeaders(customHeaders)

	checkAPIErr(s.T(), err)
	assert.Positive(s.T(), limit)
}

func (s *GetMaxVolumeSizeTestSuite) TestGetMaxVolumeSizeEndpointNotFound() {
	limit, err := s.C.GetMaxVolumeSize(context.Background())

	assert.Equal(s.T(), "The REST endpoint [GET /api/rest/limit?select=id%2Climit] cannot be found.", err.Error())
	assert.Negative(s.T(), limit)
}
