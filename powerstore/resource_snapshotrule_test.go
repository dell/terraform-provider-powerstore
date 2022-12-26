package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// Test to Create SnapShotRule
func TestAccSnapshotRule_CreateSnapShotRule(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	tests := []resource.TestStep{
		{
			Config: SnapshotRuleParamsWithTimeOfDay,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "test_snapshotrule"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "time_of_day", "21:00"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "timezone", "UTC"),
				resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "is_read_only", "false"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "delete_snaps", "true"),
			),
		},
		{
			Config: SnapshotRuleParamsWithInterval,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "test_snapshotrule"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "interval", "Four_Hours"),
				resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "is_read_only", "false"),
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
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: SnapshotRuleParamsWithTimeOfDay,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "test_snapshotrule"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "time_of_day", "21:00"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "timezone", "UTC"),
					resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "is_read_only", "false"),
				),
			},
			{
				Config: SnapshotRuleParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "test_snapshotrule_updated"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "time_of_day", "22:00"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "timezone", "UTC"),
					resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "is_read_only", "false"),
				),
			},
		},
	})
}

// Test to Create SnapShotRule with invalid values, will result in error
func TestAccSnapshotRule_CreateSnapShotRuleWithInvalidValues(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	tests := []resource.TestStep{
		{
			Config:      SnapshotRuleParamsWithInvalidInterval,
			ExpectError: regexp.MustCompile("Attribute interval must be one of these"),
		},
		{
			Config:      SnapshotRuleParamsWithInvalidTimezone,
			ExpectError: regexp.MustCompile("Attribute timezone must be one of these"),
		},
		{
			Config:      SnapshotRuleParamsWithInvalidDaysOfWeek,
			ExpectError: regexp.MustCompile("Attribute days_of_week[^ ]* must be one of these"),
		},
		{
			Config:      SnapshotRuleParamsWithInvalidNasAccessType,
			ExpectError: regexp.MustCompile("Attribute nas_access_type must be one of these"),
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

// Test to Create SnapShotRule with empty string for some params
func TestAccSnapshotRule_CreateSnapShotRuleWithEmptyString(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	tests := []resource.TestStep{
		{
			Config: SnapshotRuleParamsWithEmptyStringTimeZone,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "test_snapshotrule"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "time_of_day", "21:00"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "timezone", ""),
				resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "is_read_only", "false"),
			),
		},
		{
			Config: SnapshotRuleParamsWithEmptyStringNasAccessType,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "test_snapshotrule"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "time_of_day", "21:00"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "timezone", "UTC"),
				resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", ""),
				resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "is_read_only", "false"),
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

// Test to Create SnapShotRule with mutually exclusive params available
func TestAccSnapshotRule_CreateSnapShotRuleWithTimeOfdayandInterval(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      SnapshotRuleParamsWithTimeOfDayAndInterval,
				ExpectError: regexp.MustCompile("Could not create snapshot rule"),
			},
		},
	})
}

// Test to Update SnapShotRule with invalid values
func TestAccSnapshotRule_UpdateSnapShotRuleWithInvalidValues(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: SnapshotRuleParamsWithInterval,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "test_snapshotrule"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "interval", "Four_Hours"),
					resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "is_read_only", "false"),
				),
			},
			{
				Config:      SnapshotRuleParamsWithInvalidUpdate,
				ExpectError: regexp.MustCompile("Error updating snapshotRule"),
			},
		},
	})
}

// Test to Update SnapShotRule with invalid values
func TestAccSnapshotRule_ReadSnapShotRuleUnavailable(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: SnapshotRuleParamsWithInterval,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "test_snapshotrule"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "interval", "Four_Hours"),
					resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "is_read_only", "false"),
				),
			},
			{
				Config:      SnapshotRuleParamsWithInvalidUpdate,
				ExpectError: regexp.MustCompile("Error updating snapshotRule"),
			},
		},
	})
}

// Test to import resource but resulting in error
func TestAccSnapshotRule_ImportFailure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:        SnapshotRuleParamsWithTimeOfDay,
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

	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: SnapshotRuleParamsWithInterval,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "name", "test_snapshotrule"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "interval", "Four_Hours"),
					resource.TestCheckTypeSetElemAttr("powerstore_snapshotrule.test", "days_of_week.*", "Monday"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "desired_retention", "56"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "nas_access_type", "Snapshot"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test", "is_read_only", "false"),
				),
			},
			{
				Config:            SnapshotRuleParamsWithInterval,
				ResourceName:      "powerstore_snapshotrule.test",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: true,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "test_snapshotrule", s[0].Attributes["name"])
					assert.Equal(t, "Four_Hours", s[0].Attributes["interval"])
					assert.Equal(t, "56", s[0].Attributes["desired_retention"])
					return nil
				},
			},
		},
	})

}

var SnapshotRuleParamsWithTimeOfDay = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_snapshotrule" "test" {
	name = "test_snapshotrule"	
	time_of_day = "21:00"
	timezone = "UTC"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"
	is_read_only = false
	delete_snaps = true
}
`

var SnapshotRuleParamsWithInterval = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_snapshotrule" "test" {
	name = "test_snapshotrule"	
	interval = "Four_Hours"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"
	is_read_only = false
	delete_snaps = false
}
`

var SnapshotRuleParamsUpdate = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_snapshotrule" "test" {
	name = "test_snapshotrule_updated"	
	time_of_day = "22:00"
	timezone = "UTC"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"
	is_read_only = false
}
`

var SnapshotRuleParamsWithInvalidInterval = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_snapshotrule" "test" {
	name = "test_snapshotrule"	
	interval = "invalid"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"
	is_read_only = false
}
`

var SnapshotRuleParamsWithInvalidTimezone = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_snapshotrule" "test" {
	name = "test_snapshotrule"
	time_of_day = "22:00"
	timezone = "invalid"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"
	is_read_only = false
}
`

var SnapshotRuleParamsWithInvalidNasAccessType = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_snapshotrule" "test" {
	name = "test_snapshotrule"
	time_of_day = "22:00"
	timezone = "UTC"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "invalid"
	is_read_only = false
}
`

var SnapshotRuleParamsWithInvalidDaysOfWeek = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_snapshotrule" "test" {
	name = "test_snapshotrule"
	time_of_day = "22:00"
	timezone = "UTC"
	days_of_week = ["invalid"]
	desired_retention = 56
	nas_access_type = "Snapshot"
	is_read_only = false
}
`

var SnapshotRuleParamsWithEmptyStringTimeZone = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_snapshotrule" "test" {
	name = "test_snapshotrule"
	time_of_day = "21:00"
	timezone = ""
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"
	is_read_only = false
}
`

var SnapshotRuleParamsWithEmptyStringNasAccessType = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_snapshotrule" "test" {
	name = "test_snapshotrule"
	time_of_day = "21:00"
	timezone = "UTC"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = ""
	is_read_only = false
}
`

var SnapshotRuleParamsWithTimeOfDayAndInterval = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_snapshotrule" "test" {
	name = "test_snapshotrule"
	interval = "Four_Hours"
	time_of_day = "21:00"
	timezone = "UTC"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"
	is_read_only = false
}
`

var SnapshotRuleParamsWithInvalidUpdate = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_snapshotrule" "test" {
	name = "test_snapshotrule"
	interval = "Four_Hours"
	timezone = "Brazil__East"
	days_of_week = ["Monday"]
	desired_retention = 56
	nas_access_type = "Snapshot"
	is_read_only = false
}
`
