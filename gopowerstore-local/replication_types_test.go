/*
 *
 * Copyright © 2024 Dell Inc. or its subsidiaries. All Rights Reserved.
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

package gopowerstore

import "testing"

func TestRPOEnum_IsValid(t *testing.T) {
	const (
		RpoInvalid RPOEnum = "RPO_Invalid"
	)

	tests := []struct {
		name    string
		rpo     RPOEnum
		wantErr bool
	}{
		{
			name:    "Rpo FiveMinutes",
			rpo:     RpoFiveMinutes,
			wantErr: false,
		},
		{
			name:    "Rpo FifteenMinutes",
			rpo:     RpoFifteenMinutes,
			wantErr: false,
		},
		{
			name:    "Rpo ThirtyMinutes",
			rpo:     RpoThirtyMinutes,
			wantErr: false,
		},
		{
			name:    "Rpo OneHour",
			rpo:     RpoOneHour,
			wantErr: false,
		},
		{
			name:    "Rpo SixHours",
			rpo:     RpoSixHours,
			wantErr: false,
		},
		{
			name:    "Rpo TwelveHours",
			rpo:     RpoTwelveHours,
			wantErr: false,
		},
		{
			name:    "Rpo OneDay",
			rpo:     RpoOneDay,
			wantErr: false,
		},
		{
			name:    "Rpo Zero",
			rpo:     RpoZero,
			wantErr: false,
		},
		{
			name:    "Rpo Invalid",
			rpo:     RpoInvalid,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rpo.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("RPOEnum.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
