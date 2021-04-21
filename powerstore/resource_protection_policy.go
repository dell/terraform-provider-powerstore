package powerstore

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProtectionPolicy() *schema.Resource {

	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Create: resourceProtectionPolicyCreate,
		Read:   resourceProtectionPolicyRead,
		Update: resourceProtectionPolicyUpdate,
		Delete: resourceProtectionPolicyDelete,
		Schema: map[string]*schema.Schema{
			"resource_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_replica": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"file_systems": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"replication_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"snapshot_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"time_of_day": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"days_of_week": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"policy_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type_l10n": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_machines": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"volume_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"volume": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}

}

//function to create protection policy
func resourceProtectionPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	policy := ProtectionPolicyRequest{}

	//reading the resource data
	policy.Name = d.Get("name").(string)
	policy.Description = d.Get("description").(string)
	snapshotRulesInterface := d.Get("snapshot_rules").([]interface{})

	snapshotRules := make([]string, len(snapshotRulesInterface))
	for i, v := range snapshotRulesInterface {
		snapshotRules[i] = v.(map[string]interface{})["id"].(string)
	}
	policy.SnapshotRulesIDs = snapshotRules

	replicationRulesInterface := d.Get("replication_rules").([]interface{})

	replicationRules := make([]string, len(replicationRulesInterface))
	for i, v := range replicationRulesInterface {
		replicationRules[i] = v.(map[string]interface{})["id"].(string)
	}
	policy.ReplicationRuleIDs = replicationRules

	id, err := c.CreateProtectionPolicy(c, c.HostURL, policy)
	if err != nil {
		return err
	}

	d.SetId(id)

	return resourceProtectionPolicyRead(d, meta)
}

//function to read protection policy
func resourceProtectionPolicyRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	//getting protection policy id from resource data
	protectionPolicyID := d.Id()

	protectionPolicy, err := c.GetProtectionPolicy(c, c.HostURL, protectionPolicyID)

	if err != nil {

		return err
	}

	//adding the unmarshalled json to resource data
	d.Set("resource_id", protectionPolicy.ID)
	d.Set("name", protectionPolicy.Name)
	d.Set("description", protectionPolicy.Description)
	d.Set("is_replica", protectionPolicy.IsReplica)
	fileSystems := flattenFileSystemsData(&protectionPolicy.FileSystems)
	if err := d.Set("file_systems", fileSystems); err != nil {
		return err
	}
	volumeGroups := flattenVolumeGroupsData(&protectionPolicy.VolumeGroups)
	if err := d.Set("volume_group", volumeGroups); err != nil {
		return err
	}
	volumes := flattenVolumesData(&protectionPolicy.Volumes)
	if err := d.Set("volume", volumes); err != nil {
		return err
	}
	virtualMachines := flattenVirtualMachinesData(&protectionPolicy.VirtualMachines)
	if err := d.Set("virtual_machines", virtualMachines); err != nil {
		return err
	}
	replicationRules := flattenReplicationRulesData(&protectionPolicy.ReplicationRules)
	if err := d.Set("replication_rules", replicationRules); err != nil {
		return err
	}
	snapshotRules := flattenSnapshotRulesData(&protectionPolicy.SnapshotRules)
	if err := d.Set("snapshot_rules", snapshotRules); err != nil {
		return err
	}
	d.SetId(protectionPolicyID)
	return nil
}
func resourceProtectionPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	ID := d.Id()
	protectionPolicyRequest := ProtectionPolicyRequest{}

	protectionPolicyRequest.Name = d.Get("name").(string)
	protectionPolicyRequest.Description = d.Get("description").(string)

	protectionPolicyRequest.ReplicationRuleIDs = getStringArrayFromInterface(d.Get("replication_rules").([]interface{}))

	protectionPolicyRequest.SnapshotRulesIDs = getStringArrayFromInterface(d.Get("snapshot_rules").([]interface{}))

	err := c.UpdateProtectionPolicy(c, c.HostURL, protectionPolicyRequest, ID)
	if err != nil {
		log.Println(" Update Protection Policy error:", err.Error())
		return err
	}

	d.SetId(ID)

	return resourceProtectionPolicyRead(d, meta)

}
func resourceProtectionPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	//reading the resource data
	id := d.Id()

	err := c.DeleteProtectionPolicy(c, c.HostURL, id)
	if err != nil {
		log.Println("Delete protection policy error:", err.Error())
		return err
	}

	d.SetId("")

	log.Println("deleted")
	return nil
}

func flattenFileSystemsData(fileSystems *[]FileSystems) []interface{} {
	if fileSystems != nil {
		fss := make([]interface{}, len(*fileSystems))

		for i, fileSystem := range *fileSystems {
			fs := make(map[string]interface{})

			fs["id"] = fileSystem.ID
			fs["name"] = fileSystem.Name

			fss[i] = fs
		}

		return fss
	}

	return make([]interface{}, 0)
}

func flattenVolumeGroupsData(volumeGroups *[]VolumeGroup) []interface{} {
	if volumeGroups != nil {
		vgs := make([]interface{}, len(*volumeGroups))

		for i, volumeGroup := range *volumeGroups {
			vg := make(map[string]interface{})

			vg["id"] = volumeGroup.ID
			vg["name"] = volumeGroup.Name

			vgs[i] = vg
		}

		return vgs
	}

	return make([]interface{}, 0)
}

func flattenVolumesData(volumes *[]Volume) []interface{} {
	if volumes != nil {
		vols := make([]interface{}, len(*volumes))

		for i, volume := range *volumes {
			vol := make(map[string]interface{})

			vol["id"] = volume.ID
			vol["name"] = volume.Name

			vols[i] = vol
		}

		return vols
	}

	return make([]interface{}, 0)
}

func flattenVirtualMachinesData(virtualMachines *[]VirtualMachine) []interface{} {
	if virtualMachines != nil {
		vms := make([]interface{}, len(*virtualMachines))

		for i, virtualMachine := range *virtualMachines {
			vm := make(map[string]interface{})

			vm["id"] = virtualMachine.ID
			vm["name"] = virtualMachine.Name

			vms[i] = vm
		}

		return vms
	}

	return make([]interface{}, 0)
}

func flattenReplicationRulesData(replicationRules *[]ReplicationRule) []interface{} {
	if replicationRules != nil {
		rrs := make([]interface{}, len(*replicationRules))

		for i, replicationRule := range *replicationRules {
			rr := make(map[string]interface{})

			rr["id"] = replicationRule.ID
			rr["name"] = replicationRule.Name

			rrs[i] = rr
		}

		return rrs
	}

	return make([]interface{}, 0)
}

func flattenSnapshotRulesData(snapshotRules *[]SnapshotRule) []interface{} {
	if snapshotRules != nil {
		srs := make([]interface{}, len(*snapshotRules))

		for i, snapshotRule := range *snapshotRules {
			sr := make(map[string]interface{})

			sr["id"] = snapshotRule.ID
			sr["name"] = snapshotRule.Name
			sr["interval"] = snapshotRule.Interval
			sr["time_of_day"] = snapshotRule.TimeOfDay
			sr["days_of_week"] = snapshotRule.DaysOfWeek

			srs[i] = sr
		}

		return srs
	}

	return make([]interface{}, 0)
}

func getStringArrayFromInterface(input []interface{}) []string {
	stringOutput := make([]string, len(input))
	for i, v := range input {
		stringOutput[i] = v.(map[string]interface{})["id"].(string)
	}
	return stringOutput
}
