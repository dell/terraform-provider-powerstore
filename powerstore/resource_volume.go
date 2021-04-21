package powerstore

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVolume() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Create: resourceVolumeCreate,
		Read:   resourceVolumeRead,
		Update: resourceVolumeUpdate,
		Delete: resourceVolumeDelete,
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
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wwn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"appliance_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"node_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"creation_timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protection_policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"performance_policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_replication_destination": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"migration_session_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"location_history": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type_l10n": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"state_l10n": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"node_affinity_l10n": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"family_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"creator_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"copy_signature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"creator_type_l10n": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_app_consistent": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"created_by_rule_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_by_rule_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"expiration_timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVolumeRead(d *schema.ResourceData, meta interface{}) error {

	c := meta.(*Client)

	//getting volume id from resource data
	volumeID := d.Id()

	vol, err := c.GetVolume(c, c.HostURL, volumeID)
	if err != nil {
		return err
	}

	//addming the unmarshalled json to resource data
	d.Set("resource_id", vol.ID)
	d.Set("name", vol.Name)
	d.Set("description", vol.Description)
	d.Set("type", vol.Type)
	d.Set("wwn", vol.WWN)
	d.Set("nsid", vol.NSID)
	d.Set("appliance_id", vol.ApplianceID)
	d.Set("state", vol.State)
	d.Set("size", vol.Size)
	d.Set("node_affinity", vol.NodeAffinity)
	d.Set("creation_timestamp", vol.CreationTimeStamp)
	d.Set("protection_policy_id", vol.ProtectionPolicyID)
	d.Set("is_replication_destination", vol.IsReplicationDestination)
	d.Set("migration_session_id", vol.MigrationSessionID)
	d.Set("location_history", vol.LocationHistory)
	d.Set("type_l10n", vol.TypeL10N)
	d.Set("state_l10n", vol.StateL10N)
	d.Set("node_affinity_l10n", vol.NodeAffinityL10N)

	d.Set("family_id", vol.ProtectionData.FamilyID)
	d.Set("parent_id", vol.ProtectionData.ParentID)
	d.Set("source_id", vol.ProtectionData.SoruceID)
	d.Set("creator_type", vol.ProtectionData.CreatorType)
	d.Set("copy_signature", vol.ProtectionData.CopySignature)
	d.Set("source_timestamp", vol.ProtectionData.SourceTimeStamp)
	d.Set("creator_type_l10n", vol.ProtectionData.CreatorTypeL10N)
	d.Set("is_app_consistent", vol.ProtectionData.IsAppConsistent)
	d.Set("created_by_rule_id", vol.ProtectionData.CreatedByRuleID)
	d.Set("created_by_rule_name", vol.ProtectionData.CreatedByRuleName)
	d.Set("expiration_timestamp", vol.ProtectionData.ExpirationTimestamp)

	d.SetId(volumeID)

	return nil

}

//placeholder for resource volume create
func resourceVolumeCreate(d *schema.ResourceData, meta interface{}) error {

	c := meta.(*Client)

	vol := VolRequest{}

	//reading the resource data
	vol.Name = d.Get("name").(string)
	vol.Description = d.Get("description").(string)
	vol.ApplianceID = d.Get("appliance_id").(string)
	vol.Size = d.Get("size").(int)
	vol.IsReplicationDestination = d.Get("is_replication_destination").(bool)

	id, err := c.CreateVolume(c, c.HostURL, vol)
	if err != nil {
		return err
	}

	d.SetId(id)

	return resourceVolumeRead(d, meta)
}

//placeholder for resource volume update
func resourceVolumeUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	vol := VolRequest{}
	volumeID := d.Id()

	//reading the resource data
	vol.Name = d.Get("name").(string)
	vol.Description = d.Get("description").(string)
	vol.Size = d.Get("size").(int)
	vol.IsReplicationDestination = d.Get("is_replication_destination").(bool)

	err := c.UpdateVolume(c, c.HostURL, vol, volumeID)
	if err != nil {
		return err
	}

	return resourceVolumeRead(d, meta)
}

//placeholder for resource volume delete
func resourceVolumeDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	//reading the resource data
	id := d.Id()

	err := c.DeleteVolume(c, c.HostURL, id)
	if err != nil {
		log.Println("Delete volume error:", err.Error())
		return err
	}

	d.SetId("")

	log.Println("deleted")
	return nil
}
