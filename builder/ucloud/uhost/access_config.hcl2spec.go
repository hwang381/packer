// Code generated by "hcl2-schema -type AccessConfig"; DO NOT EDIT.\n

package uhost

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*AccessConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"PublicKey":  &hcldec.AttrSpec{Name: "public_key", Type: cty.String, Required: false},
		"PrivateKey": &hcldec.AttrSpec{Name: "private_key", Type: cty.String, Required: false},
		"Region":     &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"ProjectId":  &hcldec.AttrSpec{Name: "project_id", Type: cty.String, Required: false},
		"BaseUrl":    &hcldec.AttrSpec{Name: "base_url", Type: cty.String, Required: false},
	}
	return s
}
