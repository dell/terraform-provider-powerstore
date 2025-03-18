---
# Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://mozilla.org/MPL/2.0/


# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
page_title: "Datasource Filter Expression",
title: "Datasource Filter Expression",
linkTitle: "Datasource Filter Expression",
---

## Introduction

Many datasources for `powerstore` provider support a string field `filter_expression`.
Eventually, support for this field will be rolled out to all datasources.

This field accepts a string query in the PowerStore GET REST API format that will be used to filter the results of the datasource.
Powerful queries can be performed via this field.

## PowerStore Filter Expression Syntax

PowerStore REST API has excellent support for filtering collection query results.
It accepts queries specifying criteria that may involve all fields of a "PowerStore resource", and more.

For example, a GET request to https://<powerstore_endpoint>/api/rest/volume?and=(size.gt.1073741824,performance_policy_id.eq.default_high,or(id.eq.f64ef46f-2837-4d7d-947c-61a40fdcb074,state.eq.Ready))&select=* will fetch all volumes that meet the following criteria:
size > 1073741824 AND performance_policy_id = default_high AND (id = f64ef46f-2837-4d7d-947c-61a40fdcb074 OR state = Ready)

The complete PowerStore filter expression reference may be found in [PowerStore REST API - Filtering Response Data](https://www.dell.com/support/manuals/en-us/powerstore-1200t/pwrstr-apidevg/filtering-response-data?guid=guid-131de794-b8af-4214-bed8-64ad89c6d1cd&lang=en-us)

A more beginner-friendly guide to such expressions can be found in [Fine Tuning Your Queries](https://infohub.delltechnologies.com/en-us/p/powerstore-rest-api-using-filtering-to-fine-tune-your-queries/)

## Examples

Filtering a collection of NFS exports matching a name regex pattern

```terraform
data powerstore_nfs_export nfs_export_by_name_regex {
  filter_expression = "name=ilike.*production*"
}
```

Filtering a collection of NFS exports matching regex patterns on both name and path

```terraform
data powerstore_nfs_export nfs_export_by_name_regex_and_path {
  filter_expression = "path=ilike./us-east-revenue/sports_cars/*&name=ilike.*production*"
}
```

Quite complex filter expressions involving AND and OR operators are supported.
For example, to get all NFS exports for subdirectories of /us-east-revenue/sports_cars with name containing the string "production" or "preprod", we can write the following code:

```terraform
data powerstore_nfs_export nfs_export_by_name_regex_and_path {
  filter_expression = "and=(path=ilike./us-east-revenue/sports_cars/*,or=(name=ilike.*production*,name=ilike.*preprod*))"
}
```
