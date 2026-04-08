/*
 *
 * Copyright © 2021-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package api

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSemaphore(t *testing.T) {
	f := func(sec int, ctx context.Context, ts TimeoutSemaphoreInterface) error {
		if err := ts.Acquire(ctx); err != nil {
			return err
		}
		time.Sleep(time.Duration(sec) * time.Second)
		ts.Release(ctx)

		return nil
	}

	// long running function
	ts := NewTimeoutSemaphore(1, 1, &defaultLogger{})
	go f(3, context.Background(), ts)
	// wait for run long function
	time.Sleep(1 * time.Second)
	err := f(1, context.Background(), ts)
	assert.NotNil(t, err)

	// fast running function
	ts = NewTimeoutSemaphore(3, 1, &defaultLogger{})
	go f(1, context.Background(), ts)
	err = f(2, context.Background(), ts)
	assert.Nil(t, err)
}
