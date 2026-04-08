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

package inttests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/dell/gopowerstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type customLogger struct{}

type APITestSuite struct {
	suite.Suite
	C gopowerstore.Client
}

func TestApiSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

func (s *APITestSuite) SetupTest() {
	s.C = GetNewClient()
}

func (s *APITestSuite) TestTimeout() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancelFunc()
	_, err := s.C.GetVolumes(ctx)
	assert.NotNil(s.T(), err)
}

func (s *APITestSuite) TestTraceID() {
	_, err := s.C.GetVolumes(s.C.SetTraceID(context.Background(), "reqid-1"))
	assert.Nil(s.T(), err)
}

func (s *APITestSuite) TestCustomLogger() {
	s.C.SetLogger(&customLogger{})
	_, err := s.C.GetVolumes(s.C.SetTraceID(context.Background(), "reqid-1"))
	assert.Nil(s.T(), err)
}

func (cl *customLogger) Info(_ context.Context, format string, args ...interface{}) {
	log.Printf("INFO:"+format, args...)
}

func (cl *customLogger) Debug(_ context.Context, format string, args ...interface{}) {
	log.Printf("DEBUG:"+format, args...)
}

func (cl *customLogger) Error(_ context.Context, format string, args ...interface{}) {
	log.Printf("ERROR:"+format, args...)
}
