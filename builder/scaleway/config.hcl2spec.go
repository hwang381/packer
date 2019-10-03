// Code generated by "hcl2-schema -type Config"; DO NOT EDIT.\n

package scaleway

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Token":          &hcldec.AttrSpec{Name: "api_token", Type: cty.String, Required: false},
		"Organization":   &hcldec.AttrSpec{Name: "organization_id", Type: cty.String, Required: false},
		"Region":         &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"Image":          &hcldec.AttrSpec{Name: "image", Type: cty.String, Required: false},
		"CommercialType": &hcldec.AttrSpec{Name: "commercial_type", Type: cty.String, Required: false},
		"SnapshotName":   &hcldec.AttrSpec{Name: "snapshot_name", Type: cty.String, Required: false},
		"ImageName":      &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"ServerName":     &hcldec.AttrSpec{Name: "server_name", Type: cty.String, Required: false},
		"Bootscript":     &hcldec.AttrSpec{Name: "bootscript", Type: cty.String, Required: false},
		"BootType":       &hcldec.AttrSpec{Name: "boottype", Type: cty.String, Required: false},
		"RemoveVolume":   &hcldec.AttrSpec{Name: "remove_volume", Type: cty.Bool, Required: false},
		"UserAgent":      &hcldec.AttrSpec{Name: "useragent", Type: cty.String, Required: false},
	}
	return s
}
