/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the License);
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package powerstore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMapRecycleBinItemsToState tests the mapRecycleBinItemsToState function
func TestMapRecycleBinItemsToState(t *testing.T) {
	tests := []struct {
		name     string
		items    []map[string]interface{}
		expected int
	}{
		{
			name: "valid items",
			items: []map[string]interface{}{
				{
					"id":                   "item1",
					"name":                 "test_volume",
					"resource_type":        "volume",
					"logical_provisioned":  float64(1073741824),
					"logical_used":         float64(536870912),
					"appliance_id":         "A1",
					"deletion_timestamp":   "2024-01-01T00:00:00Z",
					"expiration_timestamp": "2024-01-08T00:00:00Z",
					"resource_type_l10n":   "Volume",
				},
				{
					"id":                   "item2",
					"name":                 "test_vg",
					"resource_type":        "volume_group",
					"logical_provisioned":  float64(2147483648),
					"logical_used":         float64(1073741824),
					"appliance_id":         "A1",
					"deletion_timestamp":   "2024-01-02T00:00:00Z",
					"expiration_timestamp": "2024-01-09T00:00:00Z",
					"resource_type_l10n":   "Volume Group",
				},
			},
			expected: 2,
		},
		{
			name: "items with nil fields",
			items: []map[string]interface{}{
				{
					"id": "item1",
					// name is nil
					"resource_type": "volume",
					// logical_provisioned is nil
					// logical_used is nil
					"appliance_id": "A1",
					// deletion_timestamp is nil
					// expiration_timestamp is nil
					// resource_type_l10n is nil
				},
			},
			expected: 1,
		},
		{
			name:     "empty items",
			items:    []map[string]interface{}{},
			expected: 0,
		},
		{
			name: "item with all nil except id",
			items: []map[string]interface{}{
				{
					"id": "item1",
				},
			},
			expected: 1,
		},
		{
			name: "item with string numeric fields",
			items: []map[string]interface{}{
				{
					"id":                   "item1",
					"name":                 "test",
					"resource_type":        "volume",
					"logical_provisioned":  "1073741824",
					"logical_used":         "536870912",
					"appliance_id":         "A1",
					"deletion_timestamp":   "2024-01-01T00:00:00Z",
					"expiration_timestamp": "2024-01-08T00:00:00Z",
					"resource_type_l10n":   "Volume",
				},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapRecycleBinItemsToState(tt.items)
			assert.Equal(t, tt.expected, len(result))
			if len(result) > 0 {
				assert.Equal(t, tt.items[0]["id"], result[0].ID.ValueString())
			}
		})
	}
}
