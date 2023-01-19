resource "powerstore_storagecontainer" "test" {
  name = "scterraform"
  quota = 10737418240
  storage_protocol = "SCSI"
  high_water_mark = 70
}
