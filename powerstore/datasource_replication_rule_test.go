/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccReplicationRuleDataSource(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	log.Println("Running TestAccReplicationRuleDataSource")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + ReplicationRuleDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.powerstore_replication_rule.test1", "replication_rules.#", "1"),
					resource.TestCheckResourceAttrPair("data.powerstore_replication_rule.test1", "replication_rules.0.name", "powerstore_replication_rule.test", "name"),
				),
			},
			{
				Config: ProviderConfigForTesting + ReplicationRuleDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.powerstore_replication_rule.test1", "replication_rules.#", "1"),
					resource.TestCheckResourceAttrPair("data.powerstore_replication_rule.test1", "replication_rules.0.id", "powerstore_replication_rule.test", "id"),
				),
			},
			{
				Config: ProviderConfigForTesting + ReplicationRuleDataSourceparamsAll,
			},
			{
				Config:      ProviderConfigForTesting + ReplicationRuleDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Replication Rules"),
			},
		},
	})
}

var ReplicationRuleDataSourceName = ReplicationRuleParamsCreate + `
data "powerstore_replication_rule" "test1" {
	name = powerstore_replication_rule.test.name
}
`

var ReplicationRuleDataSourceID = ReplicationRuleParamsCreate + `
data "powerstore_replication_rule" "test1" {
	id = powerstore_replication_rule.test.id
}
`

var ReplicationRuleDataSourceparamsAll = ReplicationRuleParamsCreate + `
data "powerstore_replication_rule" "test1" {
}
`

var ReplicationRuleDataSourceparamsNameNegative = `
data "powerstore_replication_rule" "test1" {
	name = "invalid"
}
`
