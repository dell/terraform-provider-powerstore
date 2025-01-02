/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccReplicationRuleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + ReplicationRuleParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_replication_rule.test", "name", "tfacc_replication_rule"),
					resource.TestCheckResourceAttr("powerstore_replication_rule.test", "rpo", "One_Hour"),
					resource.TestCheckResourceAttr("powerstore_replication_rule.test", "remote_system_id", remoteSystemID),
					resource.TestCheckResourceAttr("powerstore_replication_rule.test", "alert_threshold", "1000"),
				),
			},
			{
				Config:       ProviderConfigForTesting + ReplicationRuleParamsCreate,
				ResourceName: "powerstore_replication_rule.test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tfacc_replication_rule", s[0].Attributes["name"])
					assert.Equal(t, "One_Hour", s[0].Attributes["rpo"])
					assert.Equal(t, remoteSystemID, s[0].Attributes["remote_system_id"])
					assert.Equal(t, "1000", s[0].Attributes["alert_threshold"])
					return nil
				},
			},
			{
				Config: ProviderConfigForTesting + ReplicationRuleParamsModify,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_replication_rule.test", "name", "tfacc_replication_rule_renamed"),
					resource.TestCheckResourceAttr("powerstore_replication_rule.test", "rpo", "Twelve_Hours"),
					resource.TestCheckResourceAttr("powerstore_replication_rule.test", "remote_system_id", remoteSystemID),
					resource.TestCheckResourceAttr("powerstore_replication_rule.test", "alert_threshold", "1200"),
				),
			},
			{
				Config:      ProviderConfigForTesting + ReplicationRuleParamsModifyFlag,
				ExpectError: regexp.MustCompile("The attribute is_read_only cannot be modified"),
			},
			{
				Config:      ProviderConfigForTesting + ReplicationRuleCreateError,
				ExpectError: regexp.MustCompile("Error creating replication rule"),
			},
			{
				Config: ProviderConfigForTesting + ReplicationRuleReadOnly,
			},
			{
				Config:      ProviderConfigForTesting + ReplicationRuleReadOnlyModify,
				ExpectError: regexp.MustCompile("Error updating replication rule"),
			},
		},
	})
}

var ReplicationRuleParamsCreate = `
resource "powerstore_replication_rule" "test" {
	name = "tfacc_replication_rule"
    rpo = "One_Hour"
    remote_system_id = "` + remoteSystemID + `"
    alert_threshold = 1000
}
`

var ReplicationRuleCreateError = ReplicationRuleParamsCreate + `
resource "powerstore_replication_rule" "test2" {
	depends_on = ["powerstore_replication_rule.test"]
	name = "tfacc_replication_rule"
    rpo = "One_Hour"
    remote_system_id = "` + remoteSystemID + `"
    alert_threshold = 100
}
`

var ReplicationRuleParamsModify = `
resource "powerstore_replication_rule" "test" {
	name = "tfacc_replication_rule_renamed"
    rpo = "Twelve_Hours"
    remote_system_id = "` + remoteSystemID + `"
    alert_threshold = 1200
}
`

var ReplicationRuleParamsModifyFlag = `
resource "powerstore_replication_rule" "test" {
	name = "tfacc_replication_rule_renamed"
    rpo = "Twelve_Hours"
    remote_system_id = "` + remoteSystemID + `"
    alert_threshold = 1200
	is_read_only = true
}
`

var ReplicationRuleReadOnly = `
resource "powerstore_replication_rule" "test1" {
	name = "tfacc_replication_rule1"
    rpo = "One_Hour"
    remote_system_id = "` + remoteSystemID + `"
    alert_threshold = 1000
	is_read_only = true
}
`

var ReplicationRuleReadOnlyModify = `
resource "powerstore_replication_rule" "test1" {
	name = "tfacc_replication_rule"
    rpo = "One_Hour"
    remote_system_id = "` + remoteSystemID + `"
    alert_threshold = 100
}
`
