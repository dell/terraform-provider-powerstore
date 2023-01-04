package acctest

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// Test to Create SnapShotRule
func TestAccSnapshotRule_CreateSnapShotRule(t *testing.T) {

	tests := []resource.TestStep{
		{
			Config: providerConfigForTesting + SnapshotRuleParamsWithTimeOfDay,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "random_name_xcfgdfg"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "time_of_day", "21:00"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "timezone", "UTC"),
				resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "delete_snaps", "true"),
			),
		},
		{
			Config: providerConfigForTesting + SnapshotRuleParamsWithInterval,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "random_name_xcfgdfg"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "interval", "Four_Hours"),
				resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "delete_snaps", "false"),
			),
		},
	}

	for i := range tests {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testProviderFactory,
			Steps:                    []resource.TestStep{tests[i]},
		})
	}
}

// Test to update existing snapshotRule params
func TestAccSnapshotRule_UpdateSnapShotRule(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: providerConfigForTesting + SnapshotRuleParamsWithTimeOfDay,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "random_name_xcfgdfg"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "time_of_day", "21:00"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "timezone", "UTC"),
					resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
				),
			},
			{
				Config: providerConfigForTesting + SnapshotRuleParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "random_name_xcfgdfg_updated"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "time_of_day", "22:00"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "timezone", "UTC"),
					resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
				),
			},
		},
	})
}

// Test to Create SnapShotRule with invalid values, will result in error
func TestAccSnapshotRule_CreateSnapShotRuleWithInvalidValues(t *testing.T) {

	tests := []resource.TestStep{
		{
			Config:      providerConfigForTesting + SnapshotRuleParamsWithInvalidInterval,
			ExpectError: regexp.MustCompile("Attribute interval value must be one of"),
		},
		{
			Config:      providerConfigForTesting + SnapshotRuleParamsWithInvalidTimezone,
			ExpectError: regexp.MustCompile("Attribute timezone value must be one of"),
		},
		{
			Config:      providerConfigForTesting + SnapshotRuleParamsWithInvalidDaysOfWeek,
			ExpectError: regexp.MustCompile("Attribute days_of_week[^ ]* value must be one of"),
		},
		{
			Config:      providerConfigForTesting + SnapshotRuleParamsWithInvalidNasAccessType,
			ExpectError: regexp.MustCompile("Attribute nas_access_type value must be one of"),
		},
	}

	for i := range tests {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testProviderFactory,
			Steps:                    []resource.TestStep{tests[i]},
		})
	}
}

// Test to Create SnapShotRule with mutually exclusive params available
func TestAccSnapshotRule_CreateSnapShotRuleWithTimeOfdayandInterval(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      providerConfigForTesting + SnapshotRuleParamsWithTimeOfDayAndInterval,
				ExpectError: regexp.MustCompile("Attribute \"time_of_day\" cannot be specified when \"interval\" is specified"),
			},
		},
	})
}

// Test to import resource but resulting in error
func TestAccSnapshotRule_ImportFailure(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:        providerConfigForTesting + SnapshotRuleParamsWithTimeOfDay,
				ResourceName:  "powerstore_snapshotrule.test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile("Could not import snapshot rule"),
				ImportStateId: "invalid-id",
			},
		},
	})
}

// Test to import successfully
func TestAccSnapshotRule_ImportSuccess(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: providerConfigForTesting + SnapshotRuleParamsWithInterval,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "random_name_xcfgdfg"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "interval", "Four_Hours"),
					resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
				),
			},
			{
				Config:            providerConfigForTesting + SnapshotRuleParamsWithInterval,
				ResourceName:      "powerstore_snapshotrule.test",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: true,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "random_name_xcfgdfg", s[0].Attributes["name"])
					assert.Equal(t, "Four_Hours", s[0].Attributes["interval"])
					assert.Equal(t, "56", s[0].Attributes["desired_retention"])
					return nil
				},
			},
		},
	})

}

var SnapshotRuleParamsWithTimeOfDay = `
resource "powerstore_snapshotrule" "test" {
	name = "random_name_xcfgdfg"	
	time_of_day = "21:00"
	timezone = "UTC"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"	
	delete_snaps = true
}
`

var SnapshotRuleParamsWithInterval = `
resource "powerstore_snapshotrule" "test" {
	name = "random_name_xcfgdfg"	
	interval = "Four_Hours"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"	
	delete_snaps = false
}
`

var SnapshotRuleParamsUpdate = `
resource "powerstore_snapshotrule" "test" {
	name = "random_name_xcfgdfg_updated"	
	time_of_day = "22:00"
	timezone = "UTC"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"	
}
`

var SnapshotRuleParamsWithInvalidInterval = `
resource "powerstore_snapshotrule" "test" {
	name = "random_name_xcfgdfg"	
	interval = "invalid"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"	
}
`

var SnapshotRuleParamsWithInvalidTimezone = `
resource "powerstore_snapshotrule" "test" {
	name = "random_name_xcfgdfg"
	time_of_day = "22:00"
	timezone = "invalid"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"	
}
`

var SnapshotRuleParamsWithInvalidNasAccessType = `
resource "powerstore_snapshotrule" "test" {
	name = "random_name_xcfgdfg"
	time_of_day = "22:00"
	timezone = "UTC"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "invalid"	
}
`

var SnapshotRuleParamsWithInvalidDaysOfWeek = `
resource "powerstore_snapshotrule" "test" {
	name = "random_name_xcfgdfg"
	time_of_day = "22:00"
	timezone = "UTC"
	days_of_week = ["invalid"]
	desired_retention = 56
	nas_access_type = "Snapshot"	
}
`

var SnapshotRuleParamsWithTimeOfDayAndInterval = `
resource "powerstore_snapshotrule" "test" {
	name = "random_name_xcfgdfg"
	interval = "Four_Hours"
	time_of_day = "21:00"
	timezone = "UTC"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"	
}
`

var SnapshotRuleParamsWithInvalidUpdate = `
resource "powerstore_snapshotrule" "test" {
	name = "random_name_xcfgdfg"
	interval = "Four_Hours"
	timezone = "Brazil__East"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"	
}
`
