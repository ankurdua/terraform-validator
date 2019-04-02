// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceAppEngineFirewallRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppEngineFirewallRuleCreate,
		Read:   resourceAppEngineFirewallRuleRead,
		Update: resourceAppEngineFirewallRuleUpdate,
		Delete: resourceAppEngineFirewallRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAppEngineFirewallRuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"action": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"UNSPECIFIED_ACTION", "ALLOW", "DENY"}, false),
			},
			"source_range": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAppEngineFirewallRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandAppEngineFirewallRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceRangeProp, err := expandAppEngineFirewallRuleSourceRange(d.Get("source_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_range"); !isEmptyValue(reflect.ValueOf(sourceRangeProp)) && (ok || !reflect.DeepEqual(v, sourceRangeProp)) {
		obj["sourceRange"] = sourceRangeProp
	}
	actionProp, err := expandAppEngineFirewallRuleAction(d.Get("action"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action"); !isEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	priorityProp, err := expandAppEngineFirewallRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !isEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}

	url, err := replaceVars(d, config, "https://appengine.googleapis.com/v1/apps/{{project}}/firewall/ingressRules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FirewallRule: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating FirewallRule: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}/{{priority}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FirewallRule %q: %#v", d.Id(), res)

	return resourceAppEngineFirewallRuleRead(d, meta)
}

func resourceAppEngineFirewallRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://appengine.googleapis.com/v1/apps/{{project}}/firewall/ingressRules/{{priority}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AppEngineFirewallRule %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading FirewallRule: %s", err)
	}

	if err := d.Set("description", flattenAppEngineFirewallRuleDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading FirewallRule: %s", err)
	}
	if err := d.Set("source_range", flattenAppEngineFirewallRuleSourceRange(res["sourceRange"], d)); err != nil {
		return fmt.Errorf("Error reading FirewallRule: %s", err)
	}
	if err := d.Set("action", flattenAppEngineFirewallRuleAction(res["action"], d)); err != nil {
		return fmt.Errorf("Error reading FirewallRule: %s", err)
	}
	if err := d.Set("priority", flattenAppEngineFirewallRulePriority(res["priority"], d)); err != nil {
		return fmt.Errorf("Error reading FirewallRule: %s", err)
	}

	return nil
}

func resourceAppEngineFirewallRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandAppEngineFirewallRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceRangeProp, err := expandAppEngineFirewallRuleSourceRange(d.Get("source_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_range"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sourceRangeProp)) {
		obj["sourceRange"] = sourceRangeProp
	}
	actionProp, err := expandAppEngineFirewallRuleAction(d.Get("action"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	priorityProp, err := expandAppEngineFirewallRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}

	url, err := replaceVars(d, config, "https://appengine.googleapis.com/v1/apps/{{project}}/firewall/ingressRules/{{priority}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FirewallRule %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("source_range") {
		updateMask = append(updateMask, "sourceRange")
	}

	if d.HasChange("action") {
		updateMask = append(updateMask, "action")
	}

	if d.HasChange("priority") {
		updateMask = append(updateMask, "priority")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating FirewallRule %q: %s", d.Id(), err)
	}

	return resourceAppEngineFirewallRuleRead(d, meta)
}

func resourceAppEngineFirewallRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://appengine.googleapis.com/v1/apps/{{project}}/firewall/ingressRules/{{priority}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting FirewallRule %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "FirewallRule")
	}

	log.Printf("[DEBUG] Finished deleting FirewallRule %q: %#v", d.Id(), res)
	return nil
}

func resourceAppEngineFirewallRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"(?P<project>[^/]+)/(?P<priority>[^/]+)", "(?P<priority>[^/]+)"}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}/{{priority}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAppEngineFirewallRuleDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineFirewallRuleSourceRange(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineFirewallRuleAction(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineFirewallRulePriority(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func expandAppEngineFirewallRuleDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFirewallRuleSourceRange(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFirewallRuleAction(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFirewallRulePriority(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}
