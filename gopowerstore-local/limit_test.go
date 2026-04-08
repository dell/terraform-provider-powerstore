/*
 *
 * Copyright © 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"errors"
	"fmt"
	"testing"

	"github.com/dell/gopowerstore/api"
	"github.com/jarcoal/httpmock"
)

const limitMockURL = APIMockURL + limitURL

func TestClientIMPL_GetMaxVolumeSize(t *testing.T) {
	options := NewClientOptions()
	client, _ := api.New(APIMockURL, "admin", "Password", options.Insecure(), options.DefaultTimeout(), options.RateLimit(), options.RequestIDKey())

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		API          api.Client
		args         args
		mockResponse string
		want         int64
		wantErr      bool
	}{
		{
			name:         "limit found for id Max_Volume_Size",
			API:          client,
			mockResponse: fmt.Sprintf(`[{"id": "%s", "limit": %d}, {"id": "%s", "limit": %d}]`, MaxFolderSize, 128, MaxVolumeSize, 256),
			want:         256,
			wantErr:      false,
		},
		{
			name:         "limit not found for id Max_Volume_Size",
			API:          client,
			mockResponse: fmt.Sprintf(`[{"id": "%s", "limit": %d}, {"id": "%s", "limit": %d}]`, MaxFolderSize, 128, MaxVirtualVolumeSize, 512),
			want:         -1,
			wantErr:      false,
		},
		{
			name:    "limit endpoint not found",
			API:     client,
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			if tt.wantErr {
				//nolint:revive
				httpmock.RegisterResponder("GET", limitMockURL, httpmock.NewErrorResponder(errors.New("The REST endpoint [GET /api/rest/limit?select=id%2Climit] cannot be found.")))
			} else {
				httpmock.RegisterResponder("GET", limitMockURL, httpmock.NewStringResponder(200, tt.mockResponse))
			}
			c := &ClientIMPL{
				API: tt.API,
			}
			got, err := c.GetMaxVolumeSize(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientIMPL.GetMaxVolumeSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClientIMPL.GetMaxVolumeSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
