# v1.2.0

## Release Summary
The release supports PowerStore 4.1, addresses security vulnerabilities, and introduces the following resources and data sources.

## Features

### Data Sources

* `powerstore_filesystem` for reading file systems in PowerStore.
* `powerstore_filesystem_snapshot` for reading snapshots of file systems in PowerStore.
* `powerstore_nas_server` for reading NAS servers in PowerStore.
* `powerstore_nfs_export` for reading NFS exports of file systems in PowerStore.
* `powerstore_smb_share` for reading SMB Shares of file systems in PowerStore.
* `powerstore_remote_system` for reading remote systems in PowerStore.
* `powerstore_replication_rule` for reading Replication Rules in PowerStore.

### Resources

* `powerstore_filesystem` for managing file systems in PowerStore.
* `powerstore_filesystem_snapshot` for managing snapshots of file systems in PowerStore.
* `powerstore_nfs_export` for managing NFS exports of file systems in PowerStore.
* `powerstore_smb_share` for managing SMB Shares of file systems in PowerStore.
* `powerstore_replication_rule` for managing Replication Rules in PowerStore.

### Others
N/A

## Enhancements
N/A

## Bug Fixes

* `host` resource does not support unknown values inside its `initiators` field, leading to wrong validation error messages whenever variables and such are used. ([#130](https://github.com/dell/terraform-provider-powerstore/issues/130))

# v1.1.3
## Release Summary
The release supports PowerStore 4.0, upgrades to gopowerstore version 1.15.1, and addresses security vulnerablilites.
# v1.1.2
## Release Summary
The release upgrades go version to 1.22 and gopowerstore version to 1.15 and addresses security vulnerabilities.
# v1.1.1
## Release Summary
The release address security vulnerability reported in gRPC library.
# v1.1.0
## Release Summary
The release supports resources and data sources mentioned in the Features section for Dell PowerStore.
## Features

### Data Sources:
* `powerstore_volume_group` for reading volume group in PowerStore.
* `powerstore_host` for reading host in PowerStore.
* `powerstore_host_group` for reading host group in PowerStore.
* `powerstore_volume_snapshot` for reading volume snapshot in PowerStore.
* `powerstore_volume_group_snapshot` for reading volume group snapshot in PowerStore.
* `powerstore_snapshot_rule` for reading snapshot rule in PowerStore.
* `powerstore_protection_policy` for reading protection policy in PowerStore.

### Resources
* `powerstore_volume_group` for managing volume group in PowerStore.
* `powerstore_host` for managing host in PowerStore.
* `powerstore_host_group` for managing host group in PowerStore.
* `powerstore_volume_snapshot` for managing volume snapshot in PowerStore.
* `powerstore_volume_group_snapshot` for managing volume group snapshot in PowerStore.

### Others
N/A

## Enhancements
N/A

## Bug Fixes
N/A

# v1.0.0
## Release Summary
The release supports resources and data sources mentioned in the Features section for Dell PowerStore.
## Features

### Data Sources:
* `powerstore_volume` for reading volume in PowerStore.

### Resources
* `powerstore_volume` for managing Volume in PowerStore.
* `powerstore_snapshot_rule` for managing Snapshot Rule in PowerStore.
* `powerstore_protection_policy` for managing Protection Policy in PowerStore.
* `powerstore_storage_container` for managing Storage Container in PowerStore.

### Others
N/A

## Enhancements
N/A

## Bug Fixes
N/A
