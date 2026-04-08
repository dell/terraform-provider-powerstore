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

package main

import (
	"context"
	"fmt"
	"os"

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

// MyClient is extended client implementation
type MyClient struct {
	gopowerstore.Client
}

// GetMyFavoriteVolume returns info about volume with well known id
func (myc *MyClient) GetMyFavoriteVolume(ctx context.Context) (resp gopowerstore.Volume, err error) {
	favVolID := "96506fe8-cb82-41fb-91b5-2d06de887f19"
	apiClient := myc.APIClient()
	qp := apiClient.QueryParams()
	// select which fields to query
	qp.Select(resp.Fields()...)
	_, err = apiClient.Query(
		ctx,
		gopowerstore.RequestConfig{
			Method:      "GET",
			Endpoint:    "volume",
			ID:          favVolID,
			QueryParams: qp,
		},
		&resp)
	return
}

// GetVolumesByNamePrefix returns list of volumes witch names start from prefix
func (myc *MyClient) GetVolumesByNamePrefix(ctx context.Context,
	prefix string,
) (resp []gopowerstore.Volume, err error) {
	apiClient := myc.APIClient()
	qp := apiClient.QueryParams()
	qp.Select("id", "name")
	qp.RawArg("name", fmt.Sprintf("like.%s*", prefix))
	_, err = apiClient.Query(
		ctx,
		gopowerstore.RequestConfig{
			Method:      "GET",
			Endpoint:    "volume",
			QueryParams: qp,
		},
		&resp)
	return
}

func main() {
	defaultClient := initClient()
	myClient := MyClient{defaultClient}
	var err error
	_, err = myClient.GetMyFavoriteVolume(context.Background())
	_, err = myClient.GetVolumesByNamePrefix(context.Background(), "test_vol")
	if err != nil {
		panic("unexpected error")
	}
}
