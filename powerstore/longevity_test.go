package powerstore

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// func BenchmarkTf(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         primeNumbers(num)
//     }
// }

func TestAccLongevity(t *testing.T) {
	waitTimeInput := os.Getenv("LONGEVITY_WAITTIME")
	waitTime, err := time.ParseDuration(waitTimeInput)
	assert.Nil(t, err, fmt.Sprintf("Error parsing time input %s", waitTimeInput))

	iterationInput := os.Getenv("LONGEVITY_ITERATION_COUNT")
	iterationCount, err2 := strconv.Atoi(iterationInput)
	assert.Nil(t, err2, fmt.Sprintf("Error parsing iteration input %s", iterationInput))

	f, err3 := os.Create("/tmp/log.txt")
	assert.Nil(t, err3, "Error creating log.txt")
	defer f.Close()

	logger := logrus.New()
	logger.SetOutput(f)
	logger.Printf("Starting Longevity test for %d iterations", iterationCount)

	for i := 1; i <= iterationCount; i++ {
		startTime := time.Now()
		volName := fmt.Sprintf("TF_Long_Volume_%d", i)
		ppName := fmt.Sprintf("TF_Long_Protection_%d", i)
		srName := fmt.Sprintf("TF_Long_Snapshot_%d", i)
		scName := fmt.Sprintf("TF_Long_Storage_%d", i)
		longTest := fmt.Sprintf(`
			provider "powerstore" {
				username = "`+username+`"
				password = "`+password+`"
				endpoint = "`+endpoint+`"
				insecure = true
			}
			resource "powerstore_volume" "test1" {
				name = "%s"
				size = 3
				capacity_unit= "GB"
				description = "Creating volume"
				host_id=""
				host_group_id=""
				appliance_id="A1"
				volume_group_id=""
				min_size=1048576
				sector_size=512
				protection_policy_id=""
				performance_policy_id="default_medium"
				app_type="Relational_Databases_Other"
				app_type_other=""
			}
			resource "powerstore_protectionpolicy" "test1" {
				description = "Creating Protection Policy"
				name = "%s"
				snapshot_rule_names = ["test_snapshotrule_1"]
			}
			resource "powerstore_snapshotrule" "test1" {
				name = "%s"
				time_of_day = "21:00"
				timezone = "UTC"
				days_of_week = ["Monday"]
				desired_retention = 56
				nas_access_type = "Snapshot"
				is_read_only = false
				delete_snaps = true
			}
			resource "powerstore_storagecontainer" "test1" {
				name = "%s"
				quota = 10737418240
				storage_protocol = "SCSI"
				high_water_mark = 70
			}
		`, volName, ppName, srName, scName)
		logger.Printf("Iteration %d of %d starting", i, iterationCount)
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testProviderFactory,
			Steps: []resource.TestStep{
				{
					Config: longTest,
					Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.test1", "name", volName),
						//volume resource validation
						resource.TestCheckResourceAttr("powerstore_volume.test1", "size", "3"),
						resource.TestCheckResourceAttr("powerstore_volume.test1", "capacity_unit", "GB"),
						resource.TestCheckResourceAttr("powerstore_volume.test1", "description", "Creating volume"),
						resource.TestCheckResourceAttr("powerstore_volume.test1", "appliance_id", "A1"),
						resource.TestCheckResourceAttr("powerstore_volume.test1", "min_size", "1048576"),
						resource.TestCheckResourceAttr("powerstore_volume.test1", "sector_size", "512"),
						resource.TestCheckResourceAttr("powerstore_volume.test1", "performance_policy_id", "default_medium"),
						resource.TestCheckResourceAttr("powerstore_volume.test1", "app_type", "Relational_Databases_Other"),

						//protection policy resource validation
						resource.TestCheckResourceAttr("powerstore_protectionpolicy.test1", "name", ppName),
						resource.TestCheckResourceAttr("powerstore_protectionpolicy.test1", "description", "Creating Protection Policy"),
						resource.TestCheckResourceAttr("powerstore_protectionpolicy.test1", "snapshot_rule_names.0", "test_snapshotrule_1"),

						//snapshot rule resource validation
						resource.TestCheckResourceAttr("powerstore_snapshotrule.test1", "name", srName),
						resource.TestCheckResourceAttr("powerstore_snapshotrule.test1", "time_of_day", "21:00"),
						resource.TestCheckResourceAttr("powerstore_snapshotrule.test1", "timezone", "UTC"),
						resource.TestCheckResourceAttr("powerstore_snapshotrule.test1", "days_of_week.0", "Monday"),
						resource.TestCheckResourceAttr("powerstore_snapshotrule.test1", "desired_retention", "56"),
						resource.TestCheckResourceAttr("powerstore_snapshotrule.test1", "nas_access_type", "Snapshot"),
						resource.TestCheckResourceAttr("powerstore_snapshotrule.test1", "is_read_only", "false"),
						resource.TestCheckResourceAttr("powerstore_snapshotrule.test1", "delete_snaps", "true"),

						//storage container resource validation
						resource.TestCheckResourceAttr("powerstore_storagecontainer.test1", "name", scName),
						resource.TestCheckResourceAttr("powerstore_storagecontainer.test1", "quota", "10737418240"),
						resource.TestCheckResourceAttr("powerstore_storagecontainer.test1", "storage_protocol", "SCSI"),
						resource.TestCheckResourceAttr("powerstore_storagecontainer.test1", "high_water_mark", "70"),
					),
				},
			},
		})
		logger.Printf("Iteration %d complete after %.3fsec. Waiting for %s...", i, time.Since(startTime).Seconds(), waitTimeInput)
		time.Sleep(waitTime)
	}
	logger.Println("Longevity test complete")
}
