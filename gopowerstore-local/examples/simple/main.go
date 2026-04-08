/*
 *
 * Copyright © 2020-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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

package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/dell/gopowerstore"
)

func initClient() gopowerstore.Client {
	err := os.Setenv("GOPOWERSTORE_DEBUG", "true")
	if err != nil {
		panic(err)
	}
	clientOptions := gopowerstore.NewClientOptions()
	clientOptions.SetInsecure(true)
	c, err := gopowerstore.NewClientWithArgs(
		"https://127.0.0.1/api/rest",
		"admin",
		"Password",
		clientOptions)
	if err != nil {
		panic(err)
	}
	return c
}

func main() {
	c := initClient()
	// By default PowerStore API will return only volume ID
	v, err := c.GetVolume(context.Background(), "52209728-d23c-44c3-9ae0-2ada07ce81d0")
	if err != nil {
		panic(err)
	}
	fmt.Println(v.ID)

	// use context to cancel request or to set per-request timeout
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		v, err = c.GetVolume(ctx, "52209728-d23c-44c3-9ae0-2ada07ce81d0")
		wg.Done()
	}()
	cancelFunc()
	wg.Wait()

	// simple write request
	name := "test_vol1"
	size := int64(1048576)
	r, err := c.CreateVolume(context.Background(), &gopowerstore.VolumeCreate{Name: &name, Size: &size})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.ID)
	name = "test_vol2"
	size = int64(1)
	// api error handling
	_, err = c.CreateVolume(context.Background(), &gopowerstore.VolumeCreate{Name: &name, Size: &size})
	if err != nil {
		apiErr, ok := err.(gopowerstore.APIError)
		if !ok {
			// handle general errors
			panic(err)
		}
		// handle api errors(logical, validation, etc)
		if apiErr.StatusCode == 400 || apiErr.StatusCode == 422 {
			// validation failure
			fmt.Println(apiErr.Message)
		}
	}
}
