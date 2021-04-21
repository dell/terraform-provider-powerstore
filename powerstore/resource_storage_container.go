package powerstore

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceStorageContainer() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Create: resourceStorageContainerCreate,
		Read:   resourceStorageContainerRead,
		Update: resourceStorageContainerUpdate,
		Delete: resourceStorageContainerDelete,
		Schema: map[string]*schema.Schema{
			"resource_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceStorageContainerRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	//getting storage container rule id from resource data
	storageContainerID := d.Id()

	storageContainer, err := c.GetStorageContainer(c, c.HostURL, storageContainerID)

	if err != nil {

		return err
	}

	//addming the unmarshalled json to resource data
	d.Set("resource_id", storageContainer.ID)
	d.Set("name", storageContainer.Name)
	d.Set("quota", storageContainer.Quota)

	d.SetId(storageContainerID)
	return nil

}

func resourceStorageContainerCreate(d *schema.ResourceData, meta interface{}) error {

	c := meta.(*Client)

	storageContainer := StorageContainer{}

	//reading the resource data
	storageContainer.Name = d.Get("name").(string)
	storageContainer.Quota = d.Get("quota").(int)

	id, err := c.CreateStorageContainer(c, c.HostURL, storageContainer)	
	if err != nil {
		return err
	}

	d.SetId(id)

	return resourceStorageContainerRead(d, meta)
}

//placeholder for resource volume update
func resourceStorageContainerUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	storageContainer := StorageContainer{}
	storageContainerID := d.Id()

	//reading the resource data
	storageContainer.Name = d.Get("name").(string)
	storageContainer.Quota = d.Get("quota").(int)

	err := c.UpdateStorageContainer(c, c.HostURL, storageContainer, storageContainerID)
	if err != nil {
		return err
	}

	return resourceStorageContainerRead(d, meta)
}

//placeholder for resource volume delete
func resourceStorageContainerDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	//reading the resource data
	id := d.Id()

	err := c.DeleteStorageContainer(c, c.HostURL, id)
	if err != nil {
		//log.Println("Delete Storage Container error:",err.Error())
		return err
	}

	d.SetId("")

	//log.Println("deleted")
	return nil
}
