package powerstore

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnapshotRule() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Create: resourceSnapshotRuleCreate,
		Read:   resourceSnapshotRuleRead,
		Update: resourceSnapshotRuleUpdate,
		Delete: resourceSnapshotRuleDelete,
		Schema: map[string]*schema.Schema{
			"resource_id": &schema.Schema{
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
			"desired_retention": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_replica": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"interval_l10n": &schema.Schema{
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
			"days_of_week_l10n": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceSnapshotRuleRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	//getting snapshot rule id from resource data
	snapshotRuleID := d.Id()

	snapshotRule, err := c.GetSnapshotRule(c, c.HostURL, snapshotRuleID)

	if err != nil {

		return err
	}

	//addming the unmarshalled json to resource data
	d.Set("resource_id", snapshotRule.ID)
	d.Set("name", snapshotRule.Name)
	d.Set("interval", snapshotRule.Interval)
	d.Set("time_of_day", snapshotRule.TimeOfDay)
	d.Set("days_of_week", snapshotRule.DaysOfWeek)
	d.Set("desired_retention", snapshotRule.DesiredRetention)
	d.Set("is_replica", snapshotRule.IsReplica)
	d.Set("interval_l10n", snapshotRule.IntervalL10N)
	d.Set("days_of_week_l10n", snapshotRule.DaysOfWeekL10N)

	d.SetId(snapshotRuleID)
	return nil

}

//placeholder for resource snapshot rule create
func resourceSnapshotRuleCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)
	snapshotRule := SnapshotRule{}

	snapshotRule.Name = d.Get("name").(string)
	snapshotRule.Interval = d.Get("interval").(string)
	snapshotRule.TimeOfDay = d.Get("time_of_day").(string)
	snapshotRule.DesiredRetention = d.Get("desired_retention").(int)
	daysOfWeeksInterface := d.Get("days_of_week").([]interface{})

	daysOfWeek := make([]string, len(daysOfWeeksInterface))
	for i, v := range daysOfWeeksInterface {
		daysOfWeek[i] = v.(string)
	}
	snapshotRule.DaysOfWeek = daysOfWeek
	id, err := c.CreateSnapshotRule(c, c.HostURL, snapshotRule)
	log.Println(" Snapshot Create: ID:", id)
	if err != nil {
		log.Println(" Create Snapshot rule error:", err.Error())
		return err
	}

	d.SetId(id)
	return resourceSnapshotRuleRead(d, meta)
}

//placeholder for resource snapshot rule update
func resourceSnapshotRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	ID := d.Id()
	snapshotRule := SnapshotRule{}

	snapshotRule.Name = d.Get("name").(string)
	snapshotRule.Interval = d.Get("interval").(string)
	snapshotRule.TimeOfDay = d.Get("time_of_day").(string)
	snapshotRule.DesiredRetention = d.Get("desired_retention").(int)
	daysOfWeeksInterface := d.Get("days_of_week").([]interface{})

	daysOfWeek := make([]string, len(daysOfWeeksInterface))
	for i, v := range daysOfWeeksInterface {
		daysOfWeek[i] = v.(string)
	}
	snapshotRule.DaysOfWeek = daysOfWeek

	err := c.UpdateSnapshotRule(c, c.HostURL, snapshotRule, ID)
	if err != nil {
		log.Println(" Create Snapshot rule error:", err.Error())
		return err
	}

	d.SetId(ID)
	return resourceSnapshotRuleRead(d, meta)
}

//placeholder for resource snapshot rule delete
func resourceSnapshotRuleDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	//reading the resource data
	id := d.Id()

	err := c.DeleteSnapshotRule(c, c.HostURL, id)
	if err != nil {
		log.Println("Delete snapshot rule error:", err.Error())
		return err
	}

	d.SetId("")

	log.Println("deleted")
	return nil
}
