// Code generated by "hcl2-schema -type TencentCloudAccessConfig"; DO NOT EDIT.\n

package cvm

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*TencentCloudAccessConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"SecretId":       &hcldec.AttrSpec{Name: "secret_id", Type: cty.String, Required: false},
		"SecretKey":      &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"Region":         &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"Zone":           &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
		"SkipValidation": &hcldec.AttrSpec{Name: "skip_region_validation", Type: cty.Bool, Required: false},
	}
	return s
}
